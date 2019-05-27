/*
 * Copyright (c) Sematext Group, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package detach

import (
	"github.com/sematext/oxdpus/pkg/xdp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCommand builds a new detach command.
func NewCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "detach",
		Short: "Removes the XDP program from the specified device",
		Run: func(cmd *cobra.Command, args []string) {
			dev, _ := cmd.Flags().GetString("dev")
			hook, err := xdp.NewHook()
			if err != nil {
				logger.Fatal(err)
			}
			defer hook.Close()
			if err := hook.Remove(dev); err != nil {
				logger.Error(err)
			}
			logger.Infof("XDP program successfully unloaded from %s device", dev)
		},
	}
	return cmd
}
