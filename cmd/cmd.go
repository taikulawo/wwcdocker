package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/iamwwc/wwcdocker/container"

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
		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}
		
		enableTTY := ctx.Bool("ti")
		detachContainer := ctx.Bool("d")
		if enableTTY && detachContainer {
			return fmt.Errorf("ti and d args cannot both provided")
		}

		envs := ctx.StringSlice("env")
		info := &container.ContainerInfo {
			EnableTTY: enableTTY,
			Detach: detachContainer,
			Env: append(os.Environ(),envs...),
		}
		process, writePipe := container.GetContainerProcess(info)
		
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
			Name:  "name",
			Usage: "container name",
		},
		cli.StringSliceFlag{
			Name: "env",
			Usage: "environment variables",
		},
	},
}

var InitCommand = cli.Command{
	Name:  "__DON'T__CALLED__wwcdocker__init__",
	Usage: "Used In Container, User are forbidden to call this command",
	Action: func(ctx *cli.Context) {
		initCmd := ctx.Args()
		syscall.Exec(initCmd)
	},
}
