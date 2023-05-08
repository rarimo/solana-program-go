package scripts

import (
	"context"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contracts/bridge"
	"gitlab.com/rarimo/solana-program-go/contracts/commission"
)

func InitBridgeAdmin(adminSeed, program, key string, payerPrivateKey, commission string) {
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	seed := getSeedFromString(adminSeed)
	pubkey := getPubkeyFromString(key)
	comissionProgram := getSeedFromString(commission)

	payer, err := solana.PrivateKeyFromBase58(payerPrivateKey)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := bridge.InitializeAdminInstruction(programId, bridgeAdmin, payer.PublicKey(), bridge.InitializeAdminArgs{
		Instruction:       0,
		PublicKey:         pubkey,
		Seeds:             seed,
		CommissionProgram: comissionProgram,
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

func InitCommissionAdmin(program, admin, payerPrivateKey string) {
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := solana.PublicKeyFromBase58(admin)
	if err != nil {
		panic(err)
	}

	payer, err := solana.PrivateKeyFromBase58(payerPrivateKey)
	if err != nil {
		panic(err)
	}

	commissionAdmin, err := getCommissionAdmin(bridgeAdmin, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := commission.InitializeAdminInstruction(programId, commissionAdmin, bridgeAdmin, payer.PublicKey(), commission.InitializeAdminArgs{
		AcceptableTokens: []commission.CommissionTokenArg{
			commission.CommissionTokenArg{
				Token: commission.CommissionToken{
					Type:      commission.CommissionTokenTypeNative,
					PublicKey: nil,
				},
				Amount: 1234,
			},
		},
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
