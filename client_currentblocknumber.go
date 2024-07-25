package conclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/reiver/go-erorr"
)

// CurrentBlockNumber returns the current (newest) block-number for the blockchain-network which this client is conencted to.
//
// Each blockchain-network will (likely) have a different current (newest) blockchain-number.
//
// And each blockchain-network's current (newest) block-number will change over time.
//
// Example:
//
//	currentBlockNumber, err := client.CurrentBlockNumber()
func (receiver Client) CurrentBlockNumber() (*big.Int, error) {

	var chainname string = receiver.ChainName()

	var rpcurl string
	{
		var something bool

		rpcurl, something = receiver.rpcurl.Get()
		if !something {
			return nil, errMissingRPCURL
		}
	}

	var client *ethclient.Client
	{
		rpcclient, err := ethrpc.Dial(rpcurl)
		if err != nil {
			return nil, erorr.Errorf("conclient: problem creating RPC client to %q network: %s", chainname, err)
		}
		defer rpcclient.Close()

		client = ethclient.NewClient(rpcclient)
		if nil == client {
			return nil, errNilRPCClient
		}
		defer client.Close()
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, erorr.Errorf("conclient: problem getting the current (latest) block-number for %q network: %s", chainname, err)
	}

	return header.Number, nil
}
