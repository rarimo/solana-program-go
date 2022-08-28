package scripts

import (
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

const RPC = "https://api.devnet.solana.com"
const FeePayer = "4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm"

var FeePayerKey, _ = solana.PrivateKeyFromBase58(FeePayer)
var Client = rpc.New(RPC)

var (
	content2 = HashContent{
		TxHash:         "test content 2",
		CurrentAddress: "2",
		CurrentId:      "2",
		TargetAddress:  "2",
		TargetId:       "2",
		Receiver:       "2",
		CurrentNetwork: "2",
		TargetNetwork:  "2",
		Amount:         "2",
		Type:           contract.ERC721,
	}

	content3 = HashContent{
		TxHash:         "test content 3",
		CurrentAddress: "3",
		CurrentId:      "3",
		TargetAddress:  "3",
		TargetId:       "3",
		Receiver:       "3",
		CurrentNetwork: "3",
		TargetNetwork:  "3",
		Amount:         "3",
		Type:           contract.ERC721,
	}
)
