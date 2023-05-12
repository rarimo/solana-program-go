package scripts

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contracts/upgrade"
)

func TransferOwnership(program, target, newKey, prvKey, payerPrivateKey string) {
	programId := solana.MustPublicKeyFromBase58(program)
	contract := solana.MustPublicKeyFromBase58(target)
	payer := solana.MustPrivateKeyFromBase58(payerPrivateKey)

	x, y := secp256k1.S256().ScalarBaseMult(hexutil.MustDecode(prvKey))
	prv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: secp256k1.S256(),
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(hexutil.MustDecode(prvKey)),
	}

	fmt.Println(hexutil.Encode(elliptic.Marshal(secp256k1.S256(), prv.X, prv.Y)))

	signature, err := crypto.Sign(crypto.Keccak256(hexutil.MustDecode(newKey)[1:]), prv)
	if err != nil {
		panic(err)
	}

	upgradeAdmin, err := GetUpgradeAdmin(contract, programId)
	if err != nil {
		panic(err)
	}

	args := upgrade.TransferOwnershipArgs{
		RecoveryId: signature[64],
	}

	copy(args.NewPublicKey[:], hexutil.MustDecode(newKey)[1:])
	copy(args.Signature[:], signature[:64])

	instruction, err := upgrade.TransferOwnershipInstruction(programId, upgradeAdmin, args)
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
		solana.TransactionPayer(payer.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.AddSignature(payer)
	if err != nil {
		panic(err)
	}

	binTx, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	Submit(binTx)
}
