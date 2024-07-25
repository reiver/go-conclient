package conclient

import (
	"github.com/reiver/go-chain10"
	"github.com/reiver/go-chaincode"
	"github.com/reiver/go-chainname"
	"github.com/reiver/go-ethaddr"
	"github.com/reiver/go-opt"
)

// Client is a client for a smart-contract on some blockchain-network.
//
// To create a new conclient.Client, use conclient.Make.
// For example:
//
//	client := conclient.MakeClient(contract, rpcurl)
type Client struct {
	contract Contract
	rpcurl   opt.Optional[string]
}

// MakeClient makes a conclient.Client for contract 'Contract' using RPC-URL 'rpcurl'.
//
// Example:
//
//	client := conclient.MakeClient(contract, rpcurl)
func MakeClient(contract Contract, rpcurl string) Client {
	return Client{
		contract:contract,
		rpcurl:opt.Something(rpcurl),
	}
}

// Chain10Exponent returns the exponent of the blockchain-network this client is attached to.
//
// The 'exponent' comes from the following equation that shows up the default-currency-code relates to the smallest-currency-code:
//
// 1 [DEFAULT-CURRENCY-CODE] = 10ⁿ [SMALLEST-CURRENCY-CODE]
//
// (Where "n" is the value that Chain10Exponent returns.)
//
// For example, on Ethereum:
//
// 1 ETH = 1,000,000,000,000,000,000 wei = 10¹⁸ wei
//
// Therefore Chain10Exponent would return: 18
//
// And also, for example, on Avalanche:
//
// 1 AVAX = 1,000,000,000 nanoAVAX = 10⁹ nanoAVAX
//
// Therefore Chain10Exponent would return: 9
//
// The reason this number is useful is this —
//
// Often you will be working with numbers that aer in the smallest-currency of the blockchain-network.
// (Ex: "wei" rather than "ETH".)
// But you need to display things in the default-currency.
// (Ex: "ETH" rather than "wei".)
// So you need to do a conversion.
// (Ex: convert "wei" to "ETH".)
// The number Chain10Exponent returns will enable you to do that conversion.
//
// For example:
//
//	var chainid uint64 = // ...
//	
//	// ...
//	
//	var wei *big.Int = // ...
//	
//	// ...
//	
//	var exponent *big.Int = new(big.Int).SetUint64( chain10.Exponent(chainid) )
//	
//	var conversionDenominator *big.Int = new(big.Int).Exp(10, exponent)
//	
//	var ethrat *big.Rat = new(big.Rat).SetFrac(wei, conversionDenominator)
//	var eth *big.Float = new(big.Float).SetRat(ethrat)
func (receiver Client) Chain10Exponent() (uint64, bool) {
	return chain10.Exponent(receiver.contract.ChainID)
}

// ChainCode returns the (curency) chain-code of the blockchain-network this client is attached to.
//
// See https://chainlist.org/ for the list of chain-ids mapped to (currency) chain-codes.
func (receiver Client) ChainCode() string {
	return chaincode.ChainCode(receiver.contract.ChainID)
}

// ChainName returns the (human-legible) chain-name of the blockchain-network this client is attached to.
//
// See https://chainlist.org/ for the list of chain-ids mapped to chain-names.
func (receiver Client) ChainName() string {
	return chainname.ChainName(receiver.contract.ChainID)
}

// ChainID returns the chain-id of the blockchain-network this client is attached to.
//
// See https://chainlist.org/ for the list of chain-ids.
func (receiver Client) ChainID() uint64 {
	return receiver.contract.ChainID
}

func (receiver Client) ContractAddress() ethaddr.Address {
	return receiver.contract.Address
}
