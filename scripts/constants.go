package scripts

import (
	"crypto/rand"
	"fmt"

	"github.com/olegfomenko/solana-go/rpc"
	xcrypto "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operations"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
)

const RPC = "https://api.devnet.solana.com"

var Client = rpc.New(RPC)

var (
	content1 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash1", "networkFrom1", "eventId1").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content2 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash2", "networkFrom2", "eventId2").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content3 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash3", "networkFrom3", "eventId3").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content4 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash3", "networkFrom3", "eventId3").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content5 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash5", "networkFrom5", "eventId5").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content6 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash6", "networkFrom6", "eventId6").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content7 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash7", "networkFrom7", "eventId7").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content8 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash8", "networkFrom8", "eventId8").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}

	content9 = xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin("txHash9", "networkFrom9", "eventId9").GetOrigin(),
		Receiver:       getRandSlice(),
		TargetNetwork:  "Solana",
		TargetContract: getRandSlice(),
		Data: operations.NewTransferOperation(
			"",
			"",
			fmt.Sprint("1"), "USDT", "USDT", "", 9).GetContent(),
	}
)

func getRandSlice() []byte {
	var hash [32]byte
	rand.Read(hash[:])
	return hash[:]
}
