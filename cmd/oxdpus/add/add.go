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

package add

import (
	"github.com/sematext/oxdpus/pkg/blacklist"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
)

func NewCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Appends a new IP address to the blacklist",
		Run: func(cmd *cobra.Command, args []string) {
			ip, _ := cmd.Flags().GetString("ip")
			m, err := blacklist.NewMap()
			if err != nil {
				logger.Fatal(err)
			}
			if m.Add(net.ParseIP(ip)); err != nil {
				logger.Error(err)
				return
			}
			logger.Infof("%s address added to the blacklist", ip)
		},
	}
	return cmd
}
