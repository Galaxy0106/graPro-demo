package cgns

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

const cgroupCpuHierarchyMount = "/sys/fs/cgroup/cpu,cpuacct"
const defaultCgroup = "mygroup_1"

func RemoveCgroups() error {
	//clear tasks
	os.Truncate(path.Join(cgroupCpuHierarchyMount, defaultCgroup, "tasks"), 0)

	//remove cgroup
	err := os.RemoveAll(path.Join(cgroupCpuHierarchyMount, "mygroup_1"))
	return err
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("Delete cgroup dir successfully!")
	// }
}

func AddProcessToCgroups(pid int, period int, quota int) error {
	err := os.Mkdir(path.Join(cgroupCpuHierarchyMount, defaultCgroup), 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("Create cgroup dir successfully!")
	// }
	ioutil.WriteFile(path.Join(cgroupCpuHierarchyMount, defaultCgroup, "tasks"), []byte(strconv.Itoa(pid)), 0777)
	ioutil.WriteFile(path.Join(cgroupCpuHierarchyMount, defaultCgroup, "cpu.cfs_period_us"), []byte(strconv.Itoa(period)), 0777)
	ioutil.WriteFile(path.Join(cgroupCpuHierarchyMount, defaultCgroup, "cpu.cfs_quota_us"), []byte(strconv.Itoa(quota)), 0777)
	// fmt.Printf("%v\n", pid)
	return err
}
