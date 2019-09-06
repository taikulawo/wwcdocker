package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/iamwwc/wwcdocker/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wwcdocker"
	app.Commands = []cli.Command{
		cmd.RunCommand,
		cmd.InitCommand,
	}
	app.Before = func(ctx *cli.Context) error {
		log.SetOutput(os.Stdout)
		log.SetReportCaller(true)
		formatter := &log.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				filepath := strings.TrimPrefix(f.Function, "github.com/iamwwc/wwcdocker/")
				return fmt.Sprintf("%s()", filepath), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		}
		log.SetFormatter(formatter)
		// log.SetLevel(log.DebugLevel)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}
