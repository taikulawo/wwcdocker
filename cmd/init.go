package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/iamwwc/wwcdocker/common"
	"github.com/iamwwc/wwcdocker/container"
	"github.com/urfave/cli"
)

var InitCommand = cli.Command{
	Name:  "__DON'T__CALL__wwcdocker__init__",
	Usage: "Used in Container, User are forbidden to call this command",
	Action: func(ctx *cli.Context) error {
		log.Info("Init process started")
		// 注入的 fd 是3，不是4...
		// 就因为这个问题，相继引出了
		// https://stackoverflow.com/questions/57806908/the-task-is-removed-from-cgroup-after-the-exit
		// https://github.com/iamwwc/wwcdocker/issues/1
		// 可算搞明白僵尸进程了
		// 可算坑死我了 :(
		b, err := common.ReadFromFd(3)
		log.Infof("Read from parent process %s", b)
		if err != nil {
			log.Error(err)
			return err
		}

		setUpMount()

		cmdArrays := strings.Split(b, " ")
		absolutePath, err := exec.LookPath(cmdArrays[0])
		args := cmdArrays[1:]
		log.Debugf("Found exec binary %s with cmd args %s", absolutePath, args)
		if err != nil {
			return fmt.Errorf("Fail to Lookup path %s. Error: %v", cmdArrays[0], err)
		}
		// env 在 容器里已经注入过了，这里 Environ 包含着 user 注入进来的 env
		if err := syscall.Exec(absolutePath, args, os.Environ()); err != nil {
			log.Error(err)
			return fmt.Errorf("Fail to Exec process in container. Error: %v", err)
		}
		return nil
	},
	Hidden:   true,
	HideHelp: true,
}

func setUpMount() error {
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("Get current working directory error. %s", err)
		return err
	}
	// base := path.Dir(pwd)

	// syscall.Mount(base, base, "bind", syscall.MS_BIND | syscall.MS_REC, "")
	// if err := syscall.Mount("", base, "", syscall.MS_PRIVATE, ""); err != nil {
	// 	log.Error(err)
	// 	return err
	// }
	
	// common.Exec("mount","--make-rprivate","/")
	
	syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	if err := container.PivotRoot(pwd); err != nil {
		log.Errorf("Error when call pivotRoot %v", err)
		return err
	}


	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NODEV | syscall.MS_NOSUID
	if err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), ""); err != nil {
		return fmt.Errorf("Fail to mount /proc fs in container process. Error: %v", err)
	}
	return syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}
