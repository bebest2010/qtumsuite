package qtumutil_test

import (
	"strings"
	"testing"

	"github.com/bebest2010/qtumsuite/qtumd/chaincfg"
	"github.com/bebest2010/qtumsuite/qtumutil"
)

// TestTx tests the API for Tx.
func TestAddrMainNet(t *testing.T) {
	var chainParams = &chaincfg.MainNetAddrParams
	pubKeyHash := make([]byte, 20)
	addr, err := qtumutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		t.Errorf("create address err %v", err)
	}
	addrOut := addr.String()
	expectAddr := "QLbz7JHiBTspS962RLKV8GndWFwiJNvEPz"
	if strings.Compare(addrOut, expectAddr) != 0 {
		t.Errorf("address err expect %v, actual %v", expectAddr, addrOut)
	}
}

func TestAddrTestNet(t *testing.T) {
	var chainParams = &chaincfg.TestNet3AddrParams
	pubKeyHash := make([]byte, 20)
	addr, err := qtumutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		t.Errorf("create address err %v", err)
	}
	addrOut := addr.String()
	expectAddr := "qHZPA2maCec991jPwLyFC3fQXXvCQKGhkU"
	if strings.Compare(addrOut, expectAddr) != 0 {
		t.Errorf("address err expect %v, actual %v", expectAddr, addrOut)
	}
}
