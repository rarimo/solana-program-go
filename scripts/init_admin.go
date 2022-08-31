package scripts

import (
	"context"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func InitBridgeAdmin(adminSeed, program, key string, payerPrivateKey string) {
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	seed := getSeedFromString(adminSeed)
	pubkey := getPubkeyFromString(key)

	payer, err := solana.PrivateKeyFromBase58(payerPrivateKey)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.InitializeAdminInstruction(programId, bridgeAdmin, payer.PublicKey(), contract.InitializeAdminArgs{
		Instruction: 0,
		PublicKey:   pubkey,
		Seeds:       seed,
	})
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
		solana.TransactionPayer(payer.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.AddSignature(payer)
	if err != nil {
		panic(err)
	}

	binTx, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	Submit(binTx)
}
