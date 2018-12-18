//
package qtumd_test

import (
	"fmt"
	"testing"

	"github.com/bebest2010/qtumsuite/qtumd/rpcclient"
)

type Network int

const (
	_ Network = iota
	// MainNetwork bitcoin mainnet
	MainNetwork
	// TestNetwork bitcoin testnet
	TestNetwork
)

// Btc class
type Qtum struct {
	Client     *rpcclient.Client
	Network    Network
	Compressed bool
}

type HttpConConfig struct {
	HTTPPostMode bool
	DisableTLS   bool
	Host         string
	User         string
	Pass         string
}

// NewQtum return a Qtum instance
func NewQtum(RPCUrl, RPCUser, RPCPwd string, net Network, compressed bool) (*Qtum, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         RPCUrl,
		User:         RPCUser,
		Pass:         RPCPwd,
	}, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &Qtum{
		Client:     client,
		Network:    net,
		Compressed: compressed,
	}, nil
}

// Close close the client connection
func (qtum *Qtum) Close() {
	qtum.Client.Shutdown()
}

func TestRpcBlockHash(t *testing.T) {
	qtum, err := NewQtum("127.0.0.1:3889", "test", "test1234", TestNetwork, false)
	if err != nil {
		fmt.Println("NewQtum err : ", err)
		return
	}
	defer qtum.Close()
	blockCount, err := qtum.Client.GetBlockCount()
	if err != nil {
		fmt.Println("GetBlockCount err : ", err)
		return
	}
	fmt.Println("blockCount is ", blockCount)
	blockHash, err := qtum.Client.GetBlockHash(blockCount)
	if err != nil {
		fmt.Println("GetBlockHash err : ", err)
		return
	}
	fmt.Println("blockHash is ", blockHash)
}
