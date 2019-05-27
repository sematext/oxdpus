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

package blacklist

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	libbpf "github.com/rabbitstack/gobpf/elf"
	"github.com/sematext/oxdpus/pkg/xdp"
	"net"
	"unsafe"
)

const (
	blacklistMap = "blacklist"
)

// Map is responsible for controlling the IP addresses that are part of blacklist map.
type Map struct {
	mod *libbpf.Module
	m   *libbpf.Map
}

// NewMap constructs a new instance of map for manipulating/consulting the blacklist entries.
func NewMap() (*Map, error) {
	mod := libbpf.NewModuleFromReader(bytes.NewReader(xdp.LoadXDPBytecode()))
	if mod == nil {
		return nil, errors.New("ELF module is not initialized")
	}
	if err := mod.Load(nil); err != nil {
		return nil, err
	}
	m := mod.Map(blacklistMap)
	if m == nil {
		return nil, fmt.Errorf("unable to find %q map in ELF sections", blacklistMap)
	}
	return &Map{mod: mod, m: m}, nil
}

// Add appends a new IP address to the blacklist.
func (m *Map) Add(ip net.IP) error {
	addr := convertIPToNumber(ip)
	if err := m.mod.UpdateElement(m.m, unsafe.Pointer(&addr), unsafe.Pointer(&addr), 0); err != nil {
		return fmt.Errorf("couldn't add %s address to blacklist map: %v", ip, err)
	}
	return nil
}

// Remove deletes an IP address from the blacklist.
func (m *Map) Remove(ip net.IP) error {
	addr := convertIPToNumber(ip)
	if err := m.mod.DeleteElement(m.m, unsafe.Pointer(&addr)); err != nil {
		return fmt.Errorf("couldn't remove %s address from blacklist map: %v", ip, err)
	}
	return nil
}

// List lists all IP addresses in the blacklist map.
func (m *Map) List() []net.IP {
	var key uint32
	var nextKey uint32
	var value uint32
	addrs := make([]net.IP, 0)
	for {
		hasNext, _ := m.mod.LookupNextElement(m.m, unsafe.Pointer(&key), unsafe.Pointer(&nextKey), unsafe.Pointer(&value))
		if !hasNext {
			break
		}
		key = nextKey
		buffer := bytes.NewBuffer([]byte{})
		err := binary.Write(buffer, binary.LittleEndian, key)
		if err != nil {
			continue
		}
		addrs = append(addrs, buffer.Bytes()[:4])
	}
	return addrs
}

// Close the map and disposes all allocated resources.
func (m *Map) Close() {
	m.mod.Close()
}

// convertIPToNumber converts the native IP address to numeric representation.
func convertIPToNumber(ip net.IP) uint32 {
	var num uint32
	binary.Read(bytes.NewBuffer(ip.To4()), binary.LittleEndian, &num)
	return num
}
