package scripts

import (
	"context"
	"crypto/elliptic"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	xcrypto "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

func WithdrawNFT(adminSeed, program, txHash, eventId, token, col, networkFrom, privateKey, ownerPrivateKey string) {
	seed := getSeedFromString(adminSeed)
	programId, err := solana.PublicKeyFromBase58(program)
	if err != nil {
		panic(err)
	}

	mint, err := solana.PublicKeyFromBase58(token)
	if err != nil {
		panic(err)
	}

	collection, err := solana.PublicKeyFromBase58(col)
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
		Data: operation.NewTransferFullMetaOperation(
			hexutil.Encode(collection.Bytes()),
			hexutil.Encode(mint.Bytes()),
			"", "CryptoPunk", "CRY", "", 9).GetContent(),
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
		Amount:     1,
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

	instruction, err := contract.WithdrawNFTInstruction(programId, bridgeAdmin, mint, owner.PublicKey(), withdraw, args)
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
