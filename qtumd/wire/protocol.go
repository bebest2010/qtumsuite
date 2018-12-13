// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
)

// QtumcoinNet represents which bitcoin network a message belongs to.
type QtumcoinNet uint32

// Constants used to indicate the message bitcoin network.  They can also be
// used to seek to the next message when a stream's state is unknown, but
// this package does not provide that functionality since it's generally a
// better idea to simply disconnect clients that are misbehaving over TCP.
const (
	// pchMessageStart[0] = 0xf1;
	// pchMessageStart[1] = 0xcf;
	// pchMessageStart[2] = 0xa6;
	// pchMessageStart[3] = 0xd3;
	// MainNet represents the main qtum network.
	MainNet QtumcoinNet = 0xd3a6cff1

	// pchMessageStart[0] = 0x0d;
	// pchMessageStart[1] = 0x22;
	// pchMessageStart[2] = 0x15;
	// pchMessageStart[3] = 0x06;
	// TestNet3 represents the test network (version 3).
	TestNet3 QtumcoinNet = 0x0615220d
)

// bnStrings is a map of bitcoin networks back to their constant names for
// pretty printing.
var bnStrings = map[QtumcoinNet]string{
	MainNet:  "MainNet",
	TestNet3: "TestNet3",
}

// String returns the BitcoinNet in human-readable form.
func (n QtumcoinNet) String() string {
	if s, ok := bnStrings[n]; ok {
		return s
	}

	return fmt.Sprintf("Unknown QtumcoinNet (%d)", uint32(n))
}
