package conclient

import (
	"github.com/reiver/go-erorr"
)

const (
	errMissingContractAddress = erorr.Error("conclient: missing contract-address")
	errMissingRPCURL          = erorr.Error("conclient: missing rpc-url")
	errNilReceiver            = erorr.Error("conclient: nil receiver")
	errNilRPCClient           = erorr.Error("conclient: nil rpc-client")
)
