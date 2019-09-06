package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/iamwwc/wwcdocker/common"
	"github.com/iamwwc/wwcdocker/container"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// wwcdocker -ti ubuntu bash
var RunCommand = cli.Command{
	Name:  "run",
	Usage: "create a new container from given image",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}

		var cmdArray []string
		// cmdArray[0] is image name
		// ubuntu bash
		// 0				1
		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}
		imageName := cmdArray[0]
		enableTTY := ctx.Bool("ti")
		detachContainer := ctx.Bool("d")
		if enableTTY && detachContainer {
			return fmt.Errorf("ti and d args cannot both provided")
		}

		name := ctx.String("name")
		envs := ctx.StringSlice("env")
		id := common.GetRandomNumber()
		volumepoints := make(map[string]string)
		for _, point := range ctx.StringSlice("v") {
			p := strings.Split(point, ":")
			volumepoints[p[0]] = p[1]
		}
		info := &container.ContainerInfo{
			Name:          name,
			ID:            id,
			Rm:            ctx.Bool("rm"),
			EnableTTY:     enableTTY,
			Detach:        detachContainer,
			Env:           append(os.Environ(), envs...),
			VolumePoints:  volumepoints,
			InitCmd:       cmdArray[1:],
			ImageName:     imageName,
			ResourceLimit: parseResourceLimitFromcli(ctx),
		}
		return container.Run(info)
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "mem",
			Usage: "memery limit (mb)",
		},
		cli.StringFlag{
			Name:  "cpushares",
			Usage: "cpu shares",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		cli.StringSliceFlag{
			Name:  "env",
			Usage: "environment variables",
		},
		// cli.BoolFlag{
		// 	Name: "rm",
		// 	Usage: "Remove container after container stopped",
		// }
		cli.StringSliceFlag{
			Name:  "v",
			Usage: "mount volume",
		},
	},
}

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
		if err != nil {
			log.Error(err)
			return err
		}
		defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NODEV | syscall.MS_NOSUID
		if err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), ""); err != nil {
			return fmt.Errorf("Fail to mount /proc fs in container process. Error: %v", err)
		}

		cmdArrays := strings.Split(b, " ")
		absolutePath, err := exec.LookPath(cmdArrays[0])
		if err != nil {
			return fmt.Errorf("Fail to Lookup path %s. Error: %v", cmdArrays[0], err)
		}
		// env 在 容器里已经注入过了，这里 Environ 包含着 user 注入进来的 env
		if err := syscall.Exec(absolutePath, cmdArrays[1:], os.Environ()); err != nil {
			return fmt.Errorf("Fail to Exec process in container. Error: %v", err)
		}
		return nil
	},
	HideHelp: true,
}
