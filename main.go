package main

import (
	"fmt"
	"os"

	"github.com/iamwwc/wwcdocker/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wwcdocker"
	app.Commands = []cli.Command{cmd.RunCommand}
	app.Before = func(ctx *cli.Context) error {
		log.SetReportCaller(false)
		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}

