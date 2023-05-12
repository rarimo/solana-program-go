package scripts

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/olegfomenko/solana-go"
	"github.com/olegfomenko/solana-go/programs/system"
	"github.com/olegfomenko/solana-go/rpc"
	"gitlab.com/rarimo/solana-program-go/contracts/upgrade"
)

func UpgradeContract(program, target, buffer, signature, payerPrivateKey string) {
	programId := solana.MustPublicKeyFromBase58(program)
	contract := solana.MustPublicKeyFromBase58(target)
	payer := solana.MustPrivateKeyFromBase58(payerPrivateKey)

	upgradeAdmin, err := GetUpgradeAdmin(contract, programId)
	if err != nil {
		panic(err)
	}

	sig := hexutil.MustDecode(signature)

	args := upgrade.UpgradeArgs{
		RecoveryId: sig[64],
	}
	copy(args.Signature[:], sig[:64])

	spill := solana.NewWallet()

	instruction, err := upgrade.UpgradeInstruction(programId, upgradeAdmin, contract, solana.MustPublicKeyFromBase58(buffer), spill.PublicKey(), args)
	if err != nil {
		panic(err)
	}

	fmt.Println(spill.PrivateKey.String())
	const sz = uint64(100000)

	lamports, err := Client.GetMinimumBalanceForRentExemption(context.TODO(), sz, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	fmt.Println(lamports)

	instruction0, err := system.NewCreateAccountInstruction(lamports, sz, solana.BPFLoaderUpgradeableProgramID, payer.PublicKey(), spill.PublicKey()).ValidateAndBuild()
	if err != nil {
		panic(err)
	}

	blockhash, err := Client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			instruction0,
			instruction,
		},
		blockhash.Value.Blockhash,
		solana.TransactionPayer(payer.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.AddSignature(payer, spill.PrivateKey)
	if err != nil {
		panic(err)
	}

	binTx, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	Submit(binTx)
}
