package cmd

import (
	"github.com/iamwwc/wwcdocker/cgroups/subsystems"
	"github.com/urfave/cli"
)

func parseResourceLimitFromcli(ctx *cli.Context) (config *subsystems.ResourceConfig) {
	memLimits := ctx.String("mem")
	cpushares := ctx.String("cpushares")
	return &subsystems.ResourceConfig{
		CPUShares: cpushares,
		MemLimit: memLimits,
	}
}
