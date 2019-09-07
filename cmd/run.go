package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/iamwwc/wwcdocker/common"
	"github.com/iamwwc/wwcdocker/container"
	"github.com/urfave/cli"
)

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
