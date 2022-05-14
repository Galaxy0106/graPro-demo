package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"graPro-demo/cgns"
	"log"
	"os/exec"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

type TaskPayload struct {
	Process_name  string
	Wasm_name     string
	Cgroup_period int
	Cgroup_quota  int
	Ns            int // 0 or 1
}

func runProcessCmd(cmd *exec.Cmd, p TaskPayload) {
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		panic(err)
	}
	if p.Ns == 1 {
		cgns.SetNetworkNamespace(cmd)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	} else {
		pid := cmd.Process.Pid
		cgns.AddProcessToCgroups(pid, p.Cgroup_period, p.Cgroup_quota)
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		panic(err)
	}

	cgns.RemoveCgroups()
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "wasm-process":
		var p TaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}
		log.Printf("process_name=%s, wasm_name=%s, cgroup_period=%d, cgroup_quota=%d, ns=%d", p.Process_name, p.Wasm_name, p.Cgroup_period, p.Cgroup_quota, p.Ns)
		cmd := exec.Command("./"+p.Process_name, "wasm-file/"+p.Wasm_name)
		runProcessCmd(cmd, p)
	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}

func StartServer() *asynq.Server {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
		},
	)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := server.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}

	return server
}
