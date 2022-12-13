package scripts

import (
	"context"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contract"
)

func DepositFT(adminSeed, program, token, receiver, network string, amount uint64, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)

	args := contract.DepositFTArgs{
		Amount:          amount,
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

	instruction, err := contract.DepositFTInstruction(programId, bridgeAdmin, mint, owner.PublicKey(), args)
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

func DepositFTBurned(adminSeed, program, tokenSeed, receiver, network string, amount uint64, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)
	owner, err := solana.PrivateKeyFromBase58(ownerPrivateKey)
	if err != nil {
		panic(err)
	}

	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	token := getSeedFromString(tokenSeed)

	mint, _, err := solana.FindProgramAddress([][]byte{token[:]}, programId)
	if err != nil {
		panic(err)
	}

	args := contract.DepositFTArgs{
		Amount:          amount,
		NetworkTo:       network,
		ReceiverAddress: receiver,
		Seeds:           seed,
		TokenSeed:       &token,
	}

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.DepositFTInstruction(programId, bridgeAdmin, mint, owner.PublicKey(), args)
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
