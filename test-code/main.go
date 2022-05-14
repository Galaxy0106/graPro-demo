package main

import (
	"fmt"
	"graPro-demo/cgns"
	"os/exec"
)

func main() {

	cmd := exec.Command("../fib", "../wasm-file/fib.wasm")
	// cmd.Stdout = os.Stdout
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		panic(err)
	}
	// cmd := exec.Command("date")
	cgns.SetNetworkNamespace(cmd)
	err = cmd.Start()
	if err != nil {
		panic(err)
	} else {
		pid := cmd.Process.Pid
		fmt.Println(pid)
		cgns.AddProcessToCgroups(pid, 50000, 50000)
	}
	// 从管道中实时获取输出并打印到终端
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
	// fmt.Println(string(out))

}
