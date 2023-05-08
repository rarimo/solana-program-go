package scripts

import (
	"context"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
)

func GetBridgeAdmin(seed [32]byte, programId solana.PublicKey) (solana.PublicKey, error) {
	return solana.CreateProgramAddress([][]byte{seed[:]}, programId)
}

func GetCommissionAdmin(bridgeAdmin, programId solana.PublicKey) (solana.PublicKey, error) {
	return solana.CreateProgramAddress([][]byte{[]byte("commission_admin"), bridgeAdmin.Bytes()}, programId)
}

func Get32ByteFromString(str string) [32]byte {
	decoded, err := base58.Decode(str)
	if err != nil {
		panic(err)
	}

	var seed [32]byte

	copy(seed[:], decoded[:])
	return seed
}

func Get64ByteFromString(adminSeed string) [64]byte {
	decoded, err := base58.Decode(adminSeed)
	if err != nil {
		panic(err)
	}

	var key [64]byte

	copy(key[:], decoded[:64])
	return key
}

func Submit(tx []byte) {
	sig, err := Client.SendRawTransaction(context.TODO(), tx)
	if err != nil {
		panic(err)
	}

	fmt.Println(sig.String())
}
