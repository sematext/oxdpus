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

package list

import (
	"fmt"
	"github.com/sematext/oxdpus/pkg/blacklist"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Shows all IP addresses registered in the blacklist",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := blacklist.NewMap()
			if err != nil {
				logger.Fatal(err)
			}
			for _, ip := range m.List() {
				fmt.Println(fmt.Sprintf("* %s", ip))
			}
		},
	}
	return cmd
}
