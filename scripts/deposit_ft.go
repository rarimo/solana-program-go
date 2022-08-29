package scripts

import (
	"context"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func DepositFT(adminSeed, program, token, receiver, network string, amount uint64) {
	seed := getSeedFromString(adminSeed)
	nonce := getRandomNonce()

	args := contract.DepositFTArgs{
		Amount:          amount,
		NetworkTo:       network,
		ReceiverAddress: receiver,
		Seeds:           seed,
		Nonce:           nonce,
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

	deposit, _, err := solana.FindProgramAddress([][]byte{nonce[:]}, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.DepositFTInstruction(programId, bridgeAdmin, mint, deposit, FeePayerKey.PublicKey(), args)
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
