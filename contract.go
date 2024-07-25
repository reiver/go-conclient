package conclient

import (
	"github.com/reiver/go-ethaddr"
)

// Contract represents a smart-contract (specified by its 'Address')
// on a particular blockchain-network (specified by its `ChainID`)
// which was created at a particular block-number (specified by 'FromBlockNumber').
type Contract struct {
	ChainID         uint64
	Address         ethaddr.Address
	FromBlockNumber uint64
}
