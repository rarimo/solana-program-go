package scripts

import (
	"context"
	"crypto/rand"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func DepositNative(adminSeed, program, receiver, network string, amount uint64) {
	seed := getSeedFromString(adminSeed)
	nonce := getRandomNonce()
	args := contract.DepositNativeArgs{
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

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	deposit, _, err := solana.FindProgramAddress([][]byte{nonce[:]}, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.DepositNativeInstruction(programId, bridgeAdmin, deposit, FeePayerKey.PublicKey(), args)
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

func getRandomNonce() [32]byte {
	var nonce [32]byte
	_, err := rand.Read(nonce[:])
	if err != nil {
		panic(err)
	}

	return nonce
}
