package main

import (
	"context"
	"time"

	"github.com/masnax/microclustertest/v1/microcluster"
	"github.com/spf13/cobra"
)

type cmdWaitready struct {
	common *CmdControl

	flagTimeout int
}

func (c *cmdWaitready) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "waitready",
		Short: "Wait for the daemon to be ready to process requests",
		RunE:  c.Run,
	}

	cmd.Flags().IntVarP(&c.flagTimeout, "timeout", "t", 0, "Number of seconds to wait before giving up"+"``")

	return cmd
}

func (c *cmdWaitready) Run(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return cmd.Help()
	}

	m, err := microcluster.App(microcluster.Args{StateDir: c.common.FlagStateDir, Verbose: c.common.FlagLogVerbose, Debug: c.common.FlagLogDebug})
	if err != nil {
		return err
	}

	ctx, cancel := cmd.Context(), func() {}
	if c.flagTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(c.flagTimeout)*time.Second)
	}
	defer cancel()

	return m.Ready(ctx)
}
