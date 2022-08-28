package scripts

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func WithdrawNative(adminSeed, program, txHash, networkFrom string, amount uint64, key string) {
	/*cKey, _ := base58.Decode("3XaERXVvVgf2MT4JtJZeEdLcpffZRHSBe8h1avTRctJeYW2itcqgjdt7Nrv1oCbaHxiwS7yv5tKs8dUccmodToJk")
	gKey, _ := base58.Decode("4kg4bHSiPM87s2zkaxZk2otrqjt5bm4AWQh4SrjwpwVcWnCkFnoUwT85z3KRDNNXkBebA9MZEDZdPNoQd4BFt1sK")

	for _, i := range cKey {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	for _, i := range gKey {
		fmt.Printf("%d ", i)
	}
	return*/

	seed := getSeedFromString(adminSeed)
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	bridgeAdmin, err := getBridgeAdmin(seed, programId)
	if err != nil {
		panic(err)
	}

	targetContent := HashContent{
		TxHash:         txHash,
		Receiver:       FeePayerKey.PublicKey().String(),
		CurrentNetwork: networkFrom,
		TargetNetwork:  "Solana",
		Amount:         fmt.Sprint(amount),
		Type:           contract.Native,
	}

	t, err := merkletree.NewTree([]merkletree.Content{targetContent, content2, content3})
	if err != nil {
		panic(err)
	}

	path, _, err := t.GetMerklePath(targetContent)
	if err != nil {
		panic(err)
	}

	prvKey, err := base58.Decode(key)
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
	r, s, err := ecdsa.Sign(rand.Reader, pk, t.MerkleRoot())
	if err != nil {
		panic(err)
	}

	var rb [32]byte
	var sb [32]byte

	copy(rb[:], r.Bytes())
	copy(sb[:], s.Bytes())

	signature := append(rb[:], sb[:]...)

	fmt.Printf("Signature %s\n", base58.Encode(signature[:64]))

	recoveredKey, err := secp256k1.RecoverPubkey(t.MerkleRoot(), append(signature, 1))
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey[1:]))

	args := contract.WithdrawArgs{
		Content: contract.SignedContent{
			TxHash:      targetContent.TxHash,
			AddressFrom: targetContent.CurrentAddress,
			TokenIdFrom: targetContent.CurrentId,
			NetworkFrom: targetContent.CurrentNetwork,
			Amount:      amount,
			TokenType:   targetContent.Type,
		},
		Path:       make([][32]byte, len(path)+1),
		RecoveryId: 1,
		Seeds:      seed,
	}

	copy(args.Signature[:], signature[:64])
	copy(args.Root[:], t.MerkleRoot())

	hash, _ := targetContent.CalculateHash()
	copy(args.Path[0][:], hash)

	for i, hash := range path {
		copy(args.Path[i+1][:], hash)
	}

	nonce := sha256.Sum256([]byte(txHash))
	withdraw, _, err := solana.FindProgramAddress([][]byte{nonce[:]}, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.WithdrawNativeInstruction(programId, bridgeAdmin, FeePayerKey.PublicKey(), withdraw, args)
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
