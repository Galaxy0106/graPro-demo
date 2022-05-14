package cgns

import (
	"os/exec"
	"syscall"
)

func SetNetworkNamespace(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWNET
}
