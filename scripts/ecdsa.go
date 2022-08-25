package scripts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
)

// GenKey returns uncompressed public key without 04 prefix
func GenKey() ([]byte, []byte) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkeyBytes := elliptic.Marshal(secp256k1.S256(), key.X, key.Y)
	prvkeyBytes := key.D.Bytes()
	return pubkeyBytes[1:], prvkeyBytes
}

func getPubkeyFromString(adminSeed string) [64]byte {
	decoded, err := base58.Decode(adminSeed)
	if err != nil {
		panic(err)
	}

	var key [64]byte

	copy(key[:], decoded[:64])
	return key
}
