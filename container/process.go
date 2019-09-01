package container

import (
	"os"
	"os/exec"
	"syscall"

	c "github.com/iamwwc/wwcdocker/cmd"
	log "github.com/sirupsen/logrus"
)

// GetContainerProcess returns cmd and writePipe.
func GetContainerProcess(info *ContainerInfo) (*exec.Cmd, *os.File) {
	readPipe, writePipe, err := NewFilePipe()
	if err != nil {
		log.Errorln(err)
		return nil, nil
	}

	initCmd, err := os.Readlink("/proc/self/exe")

	cmd := exec.Command(initCmd, c.InitCommand.Name)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | 
				syscall.CLONE_NEWIPC | 
				syscall.CLONE_NEWNET |
				syscall.CLONE_NEWPID |
				syscall.CLONE_NEWNS,
	}
	cmd.Env = info.Env
	cmd.ExtraFiles = []*os.File{readPipe}
	return cmd, writePipe
}

// NewFilePipe Create a File Pipe,
func NewFilePipe() (*os.File, *os.File, error) {
	reader, writer, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return reader, writer, nil
}
