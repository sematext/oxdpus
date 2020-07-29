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
	"github.com/sematext/oxdpus/pkg/iprange"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
	"strings"
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
			// IP range is specified in CIDR notation
			if strings.Contains(ip, "/") {
				addrs, err := iprange.FromCIDR(ip)
				if err != nil {
					logger.Fatal(err)
				}
				for _, addr := range addrs {
					if m.Add(net.ParseIP(addr)); err != nil {
						logger.Warnf("fail to add %s IP address to blacklist", addr)
						continue
					}
				}
				logger.Infof("%d addresses added to the blacklist", len(addrs))
				return
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
