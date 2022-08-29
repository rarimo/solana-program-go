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

func WithdrawNFT(adminSeed, program, txHash, token, collection, networkFrom string, key string) {
	seed := getSeedFromString(adminSeed)
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	mint, err := solana.PublicKeyFromBase58(token)
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
		TargetAddress:  collection,
		TargetId:       token,
		CurrentNetwork: networkFrom,
		TargetNetwork:  "Solana",
		Amount:         "1",
		Type:           contract.MetaplexNFT,
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
			Amount:      1,
			TokenType:   targetContent.Type,
		},
		Path:       make([][32]byte, len(path)+1),
		RecoveryId: 1,
		Seeds:      seed,
	}

	copy(args.Signature[:], signature[:64])
	copy(args.Root[:], t.MerkleRoot())

	hash, _ := targetContent.CalculateHash()
	fmt.Println("Content hash: " + base58.Encode(hash))
	copy(args.Path[0][:], hash)

	for i, hash := range path {
		copy(args.Path[i+1][:], hash)
	}

	nonce := sha256.Sum256([]byte(txHash))
	withdraw, _, err := solana.FindProgramAddress([][]byte{nonce[:]}, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.WithdrawNFTInstruction(programId, bridgeAdmin, mint, FeePayerKey.PublicKey(), withdraw, args)
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
