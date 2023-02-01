package scripts

import (
	"crypto/rand"

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
		if err == nil {
			return key, seed
		}
	}
}

func GenTokenSeed(key string) (solana.PublicKey, [32]byte) {
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

		key, _, err := solana.FindProgramAddress([][]byte{seed[:]}, programId)
		if err == nil {
			return key, seed
		}
	}
}