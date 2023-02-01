package scripts

import (
	"github.com/olegfomenko/solana-go/rpc"
)

const RPC = "https://api.devnet.solana.com"

var Client = rpc.New(RPC)
