package scripts

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func WithdrawFT(adminSeed, program, txHash, token, eventId, networkFrom string, amount uint64, privateKey string, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	owner, err := solana.PrivateKeyFromBase58(ownerPrivateKey)
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

	amountBinary := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBinary, amount)

	targetContent := HashContent{
		TxHash:         txHash,
		CurrentNetwork: networkFrom,
		EventId:        eventId,
		TargetAddress:  []byte{},
		TargetId:       mint.Bytes(),
		Receiver:       owner.PublicKey().Bytes(),
		TargetNetwork:  "Solana",
		Amount:         amountBinary,
		ProgramId:      programId.Bytes(),
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

	r, s, err := ecdsa.Sign(rand.Reader, pk, t.Root())
	if err != nil {
		panic(err)
	}

	var rb [32]byte
	var sb [32]byte

	copy(rb[:], r.Bytes())
	copy(sb[:], s.Bytes())

	signature := append(rb[:], sb[:]...)

	fmt.Printf("Signature %s\n", base58.Encode(signature[:64]))

	recoveredKey, err := secp256k1.RecoverPubkey(t.Root(), append(signature, 1))
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey[1:]))

	args := contract.WithdrawArgs{
		Amount:     amount,
		Path:       make([][32]byte, len(path)),
		RecoveryId: 1,
		Seeds:      seed,
	}

	copy(args.Signature[:], signature[:64])
	copy(args.OriginHash[:], targetContent.OriginHash())

	fmt.Println("Content hash: " + base58.Encode(targetContent.CalculateHash()))

	for i, hash := range path {
		copy(args.Path[i][:], hash)
	}

	withdraw, _, err := solana.FindProgramAddress([][]byte{targetContent.OriginHash()}, programId)
	if err != nil {
		panic(err)
	}

	instruction, err := contract.WithdrawFTInstruction(programId, bridgeAdmin, mint, owner.PublicKey(), withdraw, args)
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
