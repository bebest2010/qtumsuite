chaincfg
========

## Sample Use

```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bebest2010/qtumsuite/qtumutil"
	"github.com/bebest2010/qtumsuite/qtumd/chaincfg"
)

var testnet = flag.Bool("testnet", false, "operate on the testnet Bitcoin network")

// By default (without -testnet), use mainnet.
var chainParams = &chaincfg.MainNetAddrParams

func main() {
	flag.Parse()

	// Modify active network parameters if operating on testnet.
	if *testnet {
		chainParams = &chaincfg.TestNet3AddrParams
	}

	// later...

	// Create and print new payment address, specific to the active network.
	pubKeyHash := make([]byte, 20)
	addr, err := qtumutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
```

## Installation and Updating

```bash
$ go get -u github.com/bebest2010/qtumsuite/qtumd/chaincfg
```
