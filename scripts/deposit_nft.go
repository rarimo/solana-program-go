package scripts

import (
	"context"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contract/bridge"
)

func DepositNFT(adminSeed, program, token, receiver, network string, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)

	args := bridge.DepositNFTArgs{
		NetworkTo:       network,
		ReceiverAddress: receiver,
		Seeds:           seed,
	}

	owner, err := solana.PrivateKeyFromBase58(ownerPrivateKey)
	if err != nil {
		panic(err)
	}

	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	mint, err := solana.PublicKeyFromBase58(token)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := bridge.DepositNFTInstruction(programId, bridgeAdmin, mint, owner.PublicKey(), args)
	if err != nil {
		panic(err)
	}

	blockhash, err := Client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			instruction,
		},
		blockhash.Value.Blockhash,
		solana.TransactionPayer(owner.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.AddSignature(owner)
	if err != nil {
		panic(err)
	}

	binTx, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	Submit(binTx)
}
