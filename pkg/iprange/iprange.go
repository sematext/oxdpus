/*
 * Copyright (c) Sematext Group, Inc.
 * All Rights Reserved
 *
 * THIS IS UNPUBLISHED PROPRIETARY SOURCE CODE OF Sematext Group, Inc.
 * The copyright notice above does not evidence any
 * actual or intended publication of such source code.
 *
 */

package iprange

import "net"

// FromCIDR returns all the IP address that pertain to the specified IP range.
func FromCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	addrs := make([]string, 0)
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); nextIP(ip) {
		addrs = append(addrs, ip.String())
	}
	// remove network/broadcast addresses
	return addrs[1 : len(addrs)-1], nil
}

func nextIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
