package scripts

import (
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
)

const RPC = "https://api.devnet.solana.com"
const FeePayer = "4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm"

var FeePayerKey, _ = solana.PrivateKeyFromBase58(FeePayer)
var Client = rpc.New(RPC)
