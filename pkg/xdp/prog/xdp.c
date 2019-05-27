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

#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <linux/types.h>
#include "maps.h"

struct vlan_hdr {
	__be16 h_vlan_TCI;
	__be16 h_vlan_encapsulated_proto;
};

/* helper functions called from eBPF programs */
static int (*bpf_trace_printk)(const char *fmt, int fmt_size, ...) =
	        (void *) BPF_FUNC_trace_printk;

/* macro for printing debug info to the tracing pipe, useful just for
 debugging purposes and not recommended to use in production systems.

 use `sudo cat /sys/kernel/debug/tracing/trace_pipe` to read debug info.
 */
#define printt(fmt, ...)                                                   \
            ({                                                             \
                char ____fmt[] = fmt;                                      \
                bpf_trace_printk(____fmt, sizeof(____fmt), ##__VA_ARGS__); \
            })

SEC("xdp/xdp_ip_filter")
int xdp_ip_filter(struct xdp_md *ctx) {
    void *end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;
    u32 ip_src;
    u64 offset;
    u16 eth_type;

    struct ethhdr *eth = data;
    offset = sizeof(*eth);

    if (data + offset > end) {
        return XDP_ABORTED;
    }
    eth_type = eth->h_proto;

    /* handle VLAN tagged packet */
    if (eth_type == htons(ETH_P_8021Q) || eth_type == htons(ETH_P_8021AD)) {
	struct vlan_hdr *vlan_hdr;

	vlan_hdr = (void *)eth + offset;
	offset += sizeof(*vlan_hdr);
	if ((void *)eth + offset > end)
		return false;
	eth_type = vlan_hdr->h_vlan_encapsulated_proto; 
   }

    /* let's only handle IPv4 addresses */
    if (eth_type == ntohs(ETH_P_IPV6)) {
        return XDP_PASS;
    }

    struct iphdr *iph = data + offset;
    offset += sizeof(struct iphdr);
    /* make sure the bytes you want to read are within the packet's range before reading them */
    if (iph + 1 > end) {
        return XDP_ABORTED;
    }
    ip_src = iph->saddr;

    if (bpf_map_lookup_elem(&blacklist, &ip_src)) {
        return XDP_DROP;
    }

    return XDP_PASS;
}

char __license[] SEC("license") = "GPL";
