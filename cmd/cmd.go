package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "Run",
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
		if(enableTTY && detachContainer){
			return fmt.Errorf("ti and d args cannot both provided")
		}
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
	},
}
