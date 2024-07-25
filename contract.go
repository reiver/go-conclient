package conclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/reiver/go-ethaddr"
)

// Contract represents a smart-contract (specified by its 'Address')
// on a particular blockchain-network (specified by its `ChainID`)
// which was created at a particular block-number (specified by 'FromBlockNumber').
type Contract struct {
	ABI             abi.ABI
	Address         ethaddr.Address
	ChainID         uint64
	FromBlockNumber uint64
}
