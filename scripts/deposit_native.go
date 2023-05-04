package scripts

import (
	"context"
	"crypto/rand"

	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contract/bridge"
)

func DepositNative(adminSeed, program, receiver, network string, amount uint64, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)
	args := bridge.DepositNativeArgs{
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

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := bridge.DepositNativeInstruction(programId, bridgeAdmin, owner.PublicKey(), args)
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

func getRandomNonce() [32]byte {
	var nonce [32]byte
	_, err := rand.Read(nonce[:])
	if err != nil {
		panic(err)
	}

	return nonce
}
