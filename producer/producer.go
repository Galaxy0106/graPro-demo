package producer

import (
	"encoding/json"
	"log"

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

func StartClient() *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	return client
}

func CloseClient(client *asynq.Client) {
	client.Close()
}

func NewTask(pname string, wname string, period int, quota int, ns int) (*asynq.Task, error) {
	payload, err := json.Marshal(TaskPayload{Process_name: pname, Wasm_name: wname, Cgroup_period: period, Cgroup_quota: quota, Ns: ns})
	if err != nil {
		return nil, err
	}
	task, err := asynq.NewTask("wasm-process", payload), nil
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	return task, err
}

func EnQueue(client *asynq.Client, task *asynq.Task) {
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
