package scripts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
)

func GetBridgeAdmin(seed [32]byte, programId solana.PublicKey) (solana.PublicKey, error) {
	return solana.CreateProgramAddress([][]byte{seed[:]}, programId)
}

func GetCommissionAdmin(bridgeAdmin, programId solana.PublicKey) (solana.PublicKey, error) {
	return solana.CreateProgramAddress([][]byte{[]byte("commission_admin"), bridgeAdmin.Bytes()}, programId)
}

func GetUpgradeAdmin(contract, programId solana.PublicKey) (solana.PublicKey, error) {
	return solana.CreateProgramAddress([][]byte{[]byte("upgrade_admin"), contract.Bytes()}, programId)
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
	//sig, err := Client.SendRawTransactionWithOpts(context.TODO(), tx, rpc.TransactionOpts{
	//	SkipPreflight: true,
	//})

	sig, err := Client.SendRawTransaction(context.TODO(), tx)
	if err != nil {
		panic(err)
	}

	fmt.Println(sig.String())
}

func AmountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return []byte{}
	}

	return am.Bytes()
}

func To32Bytes(arr []byte) []byte {
	if len(arr) >= 32 {
		return arr
	}

	res := make([]byte, 32-len(arr))
	return append(res, arr...)
}
