package scripts

import (
	"crypto/rand"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
)

func GenSeed(key string) (solana.PublicKey, [32]byte) {
	programId, err := solana.PublicKeyFromBase58(key)
	if err != nil {
		panic(err)
	}

	for {
		var seed [32]byte
		_, err := rand.Read(seed[:])
		if err != nil {
			panic(err)
		}

		key, err := getBridgeAdmin(seed, programId)
		fmt.Println(key)
		if err == nil {
			return key, seed
		}
	}
}

func getBridgeAdmin(seed [32]byte, programId solana.PublicKey) (solana.PublicKey, error) {
	return solana.CreateProgramAddress([][]byte{seed[:]}, programId)
}

func getSeedFromString(adminSeed string) [32]byte {
	decoded, err := base58.Decode(adminSeed)
	if err != nil {
		panic(err)
	}

	var seed [32]byte

	copy(seed[:], decoded[:])
	return seed
}
