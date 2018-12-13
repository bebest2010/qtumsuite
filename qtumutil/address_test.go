// single sign address test
package qtumutil_test

import (
	"strings"
	"testing"

	"github.com/bebest2010/qtumsuite/qtumd/chaincfg"
	"github.com/bebest2010/qtumsuite/qtumutil"
)

func TestAddrMainNet(t *testing.T) {
	var chainParams = &chaincfg.MainNetAddrParams
	pubKeyHash := make([]byte, 20)
	addr, err := qtumutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
	if err != nil {
		t.Errorf("create address err %v", err)
		return
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
		return
	}
	addrOut := addr.String()
	expectAddr := "qHZPA2maCec991jPwLyFC3fQXXvCQKGhkU"
	if strings.Compare(addrOut, expectAddr) != 0 {
		t.Errorf("address err expect %v, actual %v", expectAddr, addrOut)
	}
}

func TestWifMainNet(t *testing.T) {
	var privKey1 = "cRfeMAyvuTtMVCWAb1EVBEDju7AivtmYL4trk4ckT4uorVEgYuoq"
	var expectAddr = "QNiYnZfDXtT79u2YRHHAjS2RFLnM59Jbku"

	//qKfwqJ95Z5BRrmfuwHvvoCuCGckqEhTNa9

	chainParams := &chaincfg.MainNetAddrParams

	priWif, err := qtumutil.DecodeWIF(privKey1)
	if err != nil {
		t.Errorf("DecodeWIF(%v) err %v", privKey1, err)
		return
	}
	// Create and print new payment address, specific to the active network.
	pubKeyBytes := priWif.PrivKey.PubKey().SerializeCompressed()
	pubKey, err := qtumutil.NewAddressPubKey(pubKeyBytes, chainParams)
	if err != nil {
		t.Errorf("Create pubkey err %v", err)
		return
	}
	addr := pubKey.EncodeAddress()
	if strings.Compare(addr, expectAddr) != 0 {
		t.Errorf("address err, expect %v,actual %v", expectAddr, addr)
	}
}

func TestWifTestNet(t *testing.T) {
	var privKey1 = "cRfeMAyvuTtMVCWAb1EVBEDju7AivtmYL4trk4ckT4uorVEgYuoq"
	var expectAddr = "qKfwqJ95Z5BRrmfuwHvvoCuCGckqEhTNa9"

	chainParams := &chaincfg.TestNet3AddrParams

	priWif, err := qtumutil.DecodeWIF(privKey1)
	if err != nil {
		t.Errorf("DecodeWIF(%v) err %v", privKey1, err)
		return
	}
	// Create and print new payment address, specific to the active network.
	pubKeyBytes := priWif.PrivKey.PubKey().SerializeCompressed()
	pubKey, err := qtumutil.NewAddressPubKey(pubKeyBytes, chainParams)
	if err != nil {
		t.Errorf("Create pubkey err %v", err)
		return
	}
	addr := pubKey.EncodeAddress()
	if strings.Compare(addr, expectAddr) != 0 {
		t.Errorf("address err, expect %v,actual %v", expectAddr, addr)
	}
}

func TestSegwitAddrMainNet(t *testing.T) {
	var privKey1 = "cRfeMAyvuTtMVCWAb1EVBEDju7AivtmYL4trk4ckT4uorVEgYuoq"
	expectAddr := "qc1qzuhqv5m9gm33gyq93k0d7zew5avmz5e8vlv49r"
	chainParams := &chaincfg.MainNetAddrParams
	wif, err := qtumutil.DecodeWIF(privKey1)
	if err != nil {
		t.Errorf("DecodeWIF(%v) err %v", privKey1, err)
		return
	}

	pubkey := wif.PrivKey.PubKey().SerializeCompressed()
	witnessProg := qtumutil.Hash160(pubkey)
	segAddr, err := qtumutil.NewAddressWitnessPubKeyHash(witnessProg, chainParams)
	if err != nil {
		t.Errorf("Create segAddr err %v", err)
		return
	}

	addr := segAddr.String()
	if strings.Compare(addr, expectAddr) != 0 {
		t.Errorf("address err, expect %v,actual %v", expectAddr, addr)
	}

}

func TestSegwitAddrTestNet(t *testing.T) {
	var privKey1 = "cRfeMAyvuTtMVCWAb1EVBEDju7AivtmYL4trk4ckT4uorVEgYuoq"
	expectAddr := "tq1qzuhqv5m9gm33gyq93k0d7zew5avmz5e8wfdt59"
	chainParams := &chaincfg.TestNet3AddrParams
	wif, err := qtumutil.DecodeWIF(privKey1)
	if err != nil {
		t.Errorf("DecodeWIF(%v) err %v", privKey1, err)
		return
	}

	pubkey := wif.PrivKey.PubKey().SerializeCompressed()
	witnessProg := qtumutil.Hash160(pubkey)
	segAddr, err := qtumutil.NewAddressWitnessPubKeyHash(witnessProg, chainParams)
	if err != nil {
		t.Errorf("Create segAddr err %v", err)
		return
	}

	addr := segAddr.String()
	if strings.Compare(addr, expectAddr) != 0 {
		t.Errorf("address err, expect %v,actual %v", expectAddr, addr)
	}

}
