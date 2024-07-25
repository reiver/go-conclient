package conclient

import (
	"context"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/reiver/go-erorr"
)

// Call calls the 'methodname' method on the contract, and returns the results.
func (receiver Client) URI(methodname string, parameters ...interface{}) ([]interface{}, error) {
	var contractAddress ethcommon.Address
	{
		var something bool

		contractAddress, something = receiver.contract.Address.Get()
		if !something {
			return nil, errMissingContractAddress
		}
	}

	var rpcurl string
	{
		var something bool

		rpcurl, something = receiver.rpcurl.Get()
		if !something {
			return nil, errMissingRPCURL
		}
	}

	var callData []byte
	{
		var err error

		callData, err = receiver.contract.ABI.Pack(methodname, parameters...)
		if nil != err {
			return nil, erorr.Errorf("conclient: problem packing method-name %q for call-data: %s", methodname, err)
		}
		if nil == callData {
			return nil, errNilCallData
		}
	}

	var callMsg = ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	var client *ethclient.Client
	{
		rpcclient, err := ethrpc.Dial(rpcurl)
		if err != nil {
			return nil, erorr.Errorf("conclient: problem creating RPC client: %s", err)
		}
		defer rpcclient.Close()

		client = ethclient.NewClient(rpcclient)
		if nil == client {
			return nil, errNilRPCClient
		}
		defer client.Close()
	}

	var result []byte
	{
		var err error

		result, err = client.CallContract(context.Background(), callMsg, nil)
		if nil != err {
			return nil, erorr.Errorf("conclient: problem calling contract (0x%x) method %q: (%T) %s", contractAddress[:], methodname, err, err)
		}
	}

	{
		return receiver.contract.ABI.Unpack(methodname, result)
	}
}
