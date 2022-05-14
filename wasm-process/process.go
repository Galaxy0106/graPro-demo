package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

const cgroupCpuHierarchyMount = "/sys/fs/cgroup/cpu,cpuacct"

func main() {
	pid := os.Getpid()
	fmt.Printf("Current process pid: %v\n", pid)
	// cmd := exec.Command("stress", "-c", "1")
	// cmd := exec.Command("../wasm-src/loop")
	// cmd := exec.Command("ls")

	// 重定向 标准输入输出
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("ERROR", err)
	// 	os.Exit(1)
	// } else {
	// 	pid = cmd.Process.Pid
	// 	err = os.Mkdir(path.Join(cgroupCpuHierarchyMount, "mygroup_1"), 0777)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		fmt.Println("Create cgroup dir successfully!")
	// 	}

	// 	ioutil.WriteFile(path.Join(cgroupCpuHierarchyMount, "mygroup_1", "tasks"), []byte(strconv.Itoa(pid)), 0777)
	// 	ioutil.WriteFile(path.Join(cgroupCpuHierarchyMount, "mygroup_1", "cpu.cfs_period_us"), []byte("50000"), 0777)
	// 	ioutil.WriteFile(path.Join(cgroupCpuHierarchyMount, "mygroup_1", "cpu.cfs_quota_us"), []byte("10000"), 0777)
	// 	fmt.Printf("%v\n", pid)
	// }

	//clear tasks
	os.Truncate(path.Join(cgroupCpuHierarchyMount, "mygroup_1", "tasks"), 0)

	//remove cgroup
	err := os.RemoveAll(path.Join(cgroupCpuHierarchyMount, "mygroup_1"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Delete cgroup dir successfully!")
	}
}
