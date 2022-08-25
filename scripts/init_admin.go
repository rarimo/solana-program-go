package scripts

import (
	"context"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func InitBridgeAdmin(adminSeed, program, key string) {
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	seed := getSeedFromString(adminSeed)
	pubkey := getPubkeyFromString(key)

	fmt.Println(base58.Encode(seed[:]))

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	fmt.Println(bridgeAdmin.String())
	fmt.Println(adminSeed)

	instruction, err := contract.InitializeAdminInstruction(programId, bridgeAdmin, FeePayerKey.PublicKey(), contract.InitializeAdminArgs{
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
		solana.TransactionPayer(FeePayerKey.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.AddSignature(FeePayerKey)
	if err != nil {
		panic(err)
	}

	binTx, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	Submit(binTx)
}
