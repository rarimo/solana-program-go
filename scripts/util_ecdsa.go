package scripts

import (
	"github.com/mr-tron/base58"
)

func getPubkeyFromString(adminSeed string) [64]byte {
	decoded, err := base58.Decode(adminSeed)
	if err != nil {
		panic(err)
	}

	var key [64]byte

	copy(key[:], decoded[:64])
	return key
}
