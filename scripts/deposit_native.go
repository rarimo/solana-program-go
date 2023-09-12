package scripts

import (
	"context"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contracts/bridge"
	"gitlab.com/rarimo/solana-program-go/contracts/commission"
)

func DepositNative(adminSeed, bridgeProgram, commissionProgram, receiver, network string, amount uint64, ownerPrivateKey string) {
	owner, err := solana.PrivateKeyFromBase58(ownerPrivateKey)
	if err != nil {
		panic(err)
	}

	bridgeProgramId, err := solana.PublicKeyFromBase58(bridgeProgram)
	if err != nil {
		panic(err)
	}

	commissionProgramId, err := solana.PublicKeyFromBase58(commissionProgram)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := GetBridgeAdmin(Get32ByteFromString(adminSeed), bridgeProgramId)
	if err != nil {
		panic(err)
	}

	commissionAdmin, err := GetCommissionAdmin(bridgeAdmin, commissionProgramId)
	if err != nil {
		panic(err)
	}

	deposit := getDepositNativeInstruction(receiver, network, amount, Get32ByteFromString(adminSeed), bridgeProgramId, bridgeAdmin, owner.PublicKey())
	charge := getGetChargeCommissionInstruction(bridge.Native, amount, commissionProgramId, commissionAdmin, bridgeAdmin, owner.PublicKey())

	blockhash, err := Client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			charge,
			deposit,
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

func getDepositNativeInstruction(receiver, network string, amount uint64, seed [32]byte, programId, bridgeAdmin, owner solana.PublicKey) solana.Instruction {
	args := bridge.DepositNativeArgs{
		Amount:          amount,
		NetworkTo:       network,
		ReceiverAddress: receiver,
		Seeds:           seed,
	}

	instruction, err := bridge.DepositNativeInstruction(programId, bridgeAdmin, owner, args)
	if err != nil {
		panic(err)
	}

	return instruction
}

func getGetChargeCommissionInstruction(typ bridge.TokenType, amount uint64, programId, commissionAdmin, bridgeAdmin, owner solana.PublicKey) solana.Instruction {
	args := commission.ChargeCommissionArgs{
		Token: commission.CommissionToken{
			Type:      commission.CommissionTokenTypeNative,
			PublicKey: nil,
		},
		Deposit:       typ,
		DepositAmount: amount,
	}

	instruction, err := commission.ChargeCommissionNativeInstruction(programId, commissionAdmin, bridgeAdmin, owner, args)
	if err != nil {
		panic(err)
	}

	return instruction
}
