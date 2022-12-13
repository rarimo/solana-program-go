package scripts

import (
	"context"
	"crypto/elliptic"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	merkle "gitlab.com/rarimo/go-merkle"
	xcrypto "gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	"gitlab.com/rarimo/solana-program-go/contract"
)

func WithdrawNative(adminSeed, program, txHash, eventId, networkFrom string, amount uint64, privateKey string, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	owner, err := solana.PrivateKeyFromBase58(ownerPrivateKey)
	if err != nil {
		panic(err)
	}

	targetContent := xcrypto.HashContent{
		Origin:         origin.NewDefaultOrigin(txHash, networkFrom, eventId).GetOrigin(),
		Receiver:       owner.PublicKey().Bytes(),
		TargetNetwork:  "Solana",
		TargetContract: programId.Bytes(),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint(amount), "").GetContent(),
	}

	t := merkle.NewTree(crypto.Keccak256, content1, targetContent, content2)
	path, _ := t.Path(targetContent)

	prvKey, err := base58.Decode(privateKey)
	if err != nil {
		panic(err)
	}

	pk, err := crypto.ToECDSA(prvKey)
	if err != nil {
		panic(err)
	}

	puk := elliptic.Marshal(secp256k1.S256(), pk.X, pk.Y)
	fmt.Println("PUB KEY: " + base58.Encode(puk[1:]))

	//signature, err := crypto.Sign(t.MerkleRoot(), pk)
	signature, err := crypto.Sign(t.Root(), pk)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature %s\n", base58.Encode(signature[:64]))

	recoveredKey, err := secp256k1.RecoverPubkey(t.Root(), signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey[1:]))

	args := contract.WithdrawArgs{
		Amount:     amount,
		Path:       make([][32]byte, len(path)),
		RecoveryId: 1,
		Seeds:      seed,
		Origin:     targetContent.Origin,
	}

	copy(args.Signature[:], signature[:64])

	fmt.Println("Content hash: " + base58.Encode(targetContent.CalculateHash()))
	for i, hash := range path {
		copy(args.Path[i][:], hash)
	}

	withdraw, _, err := solana.FindProgramAddress([][]byte{targetContent.Origin[:]}, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.WithdrawNativeInstruction(programId, bridgeAdmin, owner.PublicKey(), withdraw, args)
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

/*privteKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
if err != nil {
	panic(err)
}

pvk := privteKey.D.Bytes()
puk := elliptic.Marshal(secp256k1.S256(), privteKey.X, privteKey.Y)

fmt.Println("PRV KEY: " + base58.Encode(pvk))
fmt.Println("PUB KEY: " + base58.Encode(puk))

//PRV KEY: DnEFMLJfXXmBETgsYcaeYpnR34och4sg2LZMEjEDA9H7
//PUB KEY: QixEa5nuXdbVL4cv21LYNeXA38vmAvx7439tmABDBgYXADgEqdSXqbH54L4S64soca8qc6aLnoDMNzHdG1AnRFf9

pvk, _ := base58.Decode("DnEFMLJfXXmBETgsYcaeYpnR34och4sg2LZMEjEDA9H7")
//puk, _ := base58.Decode("QixEa5nuXdbVL4cv21LYNeXA38vmAvx7439tmABDBgYXADgEqdSXqbH54L4S64soca8qc6aLnoDMNzHdG1AnRFf9")

rand.Seed(time.Now().UnixNano())

for i := 0; i < 10; i++ {


	fmt.Printf("Signature %d %s \n ", i, base58.Encode(signature))


}

return*/

//pk, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
/*signature, err := pk.Sign(rand.Reader, t.MerkleRoot(), crypto.SHA256)
if err != nil {
	panic(err)
}
*/

/*r, s, err := ecdsa.Sign(rand.Reader, &pk, t.MerkleRoot())
if err != nil {
	panic(err)
}

signature = append(r.Bytes(), s.Bytes()...)

fmt.Println(len(signature))
fmt.Println("Signature " + base58.Encode(signature))

for reid := range []byte{1, 2, 3, 4} {
	recoveredKey, _ := secp256k1.RecoverPubkey(t.MerkleRoot(), append(signature, byte(reid)))
	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey))
}
*/

//fmt.Printf("Revovery ID: %d\n", signature[64])
//fmt.Println(ecdsa.VerifyASN1(&pk.PublicKey, t.MerkleRoot(), signature))

/*pubkeyBytes := elliptic.Marshal(secp256k1.S256(), pk.X, pk.Y)
fmt.Printf("Original pub key: %s\n", base58.Encode(pubkeyBytes))

fmt.Println(len(pubkeyBytes))*/
