package scripts

import (
	"crypto/rand"

	"github.com/olegfomenko/solana-go/rpc"
)

const RPC = "https://api.devnet.solana.com"

var Client = rpc.New(RPC)

var (
	content1 = HashContent{
		TxHash:         "test content 2",
		EventId:        "54321",
		TargetAddress:  getRandSlice(),
		TargetId:       getRandSlice(),
		Receiver:       getRandSlice(),
		CurrentNetwork: "Ethereum",
		TargetNetwork:  "Polygon",
		Amount:         getRandSlice(),
		ProgramId:      getRandSlice(),
	}

	content2 = HashContent{
		TxHash:         "test content 3",
		EventId:        "098765",
		TargetAddress:  getRandSlice(),
		TargetId:       getRandSlice(),
		Receiver:       getRandSlice(),
		CurrentNetwork: "Ethereum",
		TargetNetwork:  "Near",
		Amount:         getRandSlice(),
		ProgramId:      getRandSlice(),
	}
)

func getRandSlice() []byte {
	var hash [32]byte
	rand.Read(hash[:])
	return hash[:]
}
