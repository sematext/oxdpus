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

package root

import (
	"github.com/sematext/oxdpus/cmd/oxdpus/add"
	"github.com/sematext/oxdpus/cmd/oxdpus/attach"
	"github.com/sematext/oxdpus/cmd/oxdpus/detach"
	"github.com/sematext/oxdpus/cmd/oxdpus/list"
	"github.com/sematext/oxdpus/cmd/oxdpus/remove"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oxdpus",
	Short: "A toy tool that leverages the super powers of XDP to bring in-kernel IP filtering",
	Long:  `A toy tool that leverages the super powers of XDP to bring in-kernel IP filtering`,
}

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	attachCmd := attach.NewCommand(logger)
	attachCmd.Flags().StringP("dev", "d", "eth0", "network device to attach the XDP program")
	detachCmd := detach.NewCommand(logger)
	detachCmd.Flags().StringP("dev", "d", "eth0", "network device to attach the XDP program")
	addCmd := add.NewCommand(logger)
	addCmd.Flags().StringP("ip", "i", "172.17.0.2", "IP address to add to the blacklist")
	rmCmd := remove.NewCommand(logger)
	rmCmd.Flags().StringP("ip", "i", "172.17.0.2", "IP address to remove from the blacklist")
	listCmd := list.NewCommand(logger)
	rootCmd.AddCommand(attachCmd)
	rootCmd.AddCommand(detachCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(listCmd)
}

func Get() *cobra.Command {
	return rootCmd
}
