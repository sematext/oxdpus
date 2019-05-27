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

#define SEC(NAME) __attribute__((section(NAME), used))

static void *(*bpf_map_lookup_elem)(void *map, void *key) = (void *) BPF_FUNC_map_lookup_elem;

#define BUF_SIZE_MAP_NS 256
#define PIN_GLOBAL_NS 2

struct bpf_map_def {
    unsigned int type;
    unsigned int key_size;
    unsigned int value_size;
    unsigned int max_entries;
    unsigned int map_flags;
    unsigned int pinning;
    char namespace[BUF_SIZE_MAP_NS];
};

struct bpf_map_def SEC("maps/blacklist") blacklist = {
    .type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(u32),
    .value_size = sizeof(u32),
    .max_entries = 65536,
    .pinning = PIN_GLOBAL_NS,
    .namespace = "globals",
};