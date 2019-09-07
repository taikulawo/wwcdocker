package container

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/iamwwc/wwcdocker/common"
	log "github.com/sirupsen/logrus"
)

// Run start container run
func Run(info *ContainerInfo) error {
	process, writePipe := GetContainerProcess(info)
	ExecProcess(process, info)
	_, err := writePipe.WriteString(strings.Join(info.InitCmd, " "))
	if err != nil {
		log.Error(err)
	}
	writePipe.Close()
	err = recordContainerInfo(info)
	if info.EnableTTY {
		process.Wait()
	}
	return err
}

// GetContainerProcess returns cmd and writePipe.
func GetContainerProcess(info *ContainerInfo) (*exec.Cmd, *os.File) {
	readPipe, writePipe, err := NewFilePipe()
	if err != nil {
		log.Errorln(err)
		return nil, nil
	}

	initCmd, err := os.Readlink("/proc/self/exe")

	cmd := exec.Command(initCmd, "__DON'T__CALL__wwcdocker__init__")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
	}
	cmd.Env = info.Env
	cmd.ExtraFiles = []*os.File{readPipe}
	if info.EnableTTY {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		logDir := path.Join(common.DefaultContainerLogLocation, info.ID)
		// x 1
		// w 2
		// r 4
		if err := os.MkdirAll(logDir, 0622); err != nil {
			log.Errorf("Failed to create container %s Log folder, cause: [%s]", info.ID, err)
			return nil, nil
		}
		logFilePath := path.Join(logDir, info.ID)
		logFile, err := os.Create(logFilePath)
		if err != nil {
			log.Error(err)
			return nil, nil
		}
		cmd.Stdout = logFile
		info.FilePath["logFolder"] = logDir
		info.FilePath["logFile"] = logFilePath
	}
	cwd := NewWorkspace(common.ContainerMountRoot, info.ID, info.ImageName, info.VolumePoints)
	cmd.Dir = cwd
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

func PivotRoot(rootfs string) error {
	// 可算被我找到了
	// https://github.com/torvalds/linux/blob/d41a3effbb53b1bcea41e328d16a4d046a508381/fs/namespace.c#L3582
	if err := syscall.Mount(rootfs, rootfs, "bind", syscall.MS_BIND|syscall.MS_REC | syscall.MS_PRIVATE, ""); err != nil {
		log.Errorf("Mount %s to itself error, %v", rootfs, err)
		return err
	}

	putOld := path.Join(rootfs, ".pivotroot")
	if err := os.Mkdir(putOld, 0777); err != nil {
		log.Errorf("Failed to create putOld folder %s, error: %v", putOld, err)
		return err
	}
	if err := syscall.PivotRoot(rootfs, putOld); err != nil {
		log.Errorf("Failed to Pivot Rootfs %s, error: %v", rootfs, err)
		return err
	}
	log.Debug("PivotRoot done")

	if err := os.Chdir("/"); err != nil {
		return fmt.Errorf("chdir error %v", err)
	}

	old := path.Join("/", ".pivotroot")
	if err := syscall.Unmount(old, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("Unmount failed %v", err)
	}
	return os.Remove(old)
}
