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

import (
	"testing"
)

func TestFromCIDR(t *testing.T) {
	addrs, err := FromCIDR("192.169.1.0/24")
	if err != nil {
		t.Fatal(err)
	}
	if len(addrs) != 254 {
		t.Fatalf("addrs should contain 254 but has %d addresses", len(addrs))
	}
}
