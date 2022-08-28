package scripts

import (
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
)

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
