package scripts

//
//func MintFT(adminSeed, program, payerPrvKey string, amount uint64) {
//	seed := getSeedFromString(adminSeed)
//
//	args := contract.MintFTArgs{
//		Data: metaplex.DataV2{
//			Name:   "Tether USD",
//			Symbol: "USDT",
//			URI:    "",
//		},
//		Amount:   amount,
//		Seeds:    seed,
//		Decimals: 9,
//	}
//
//	programId, err := solana.PublicKeyFromBase58(program)
//	if err != nil {
//		panic(err)
//	}
//
//	mint, err := solana.NewRandomPrivateKey()
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("FT token mint: " + mint.PublicKey().String())
//
//	bridgeAdmin, err := getBridgeAdmin(seed, programId)
//	if err != nil {
//		panic(err)
//	}
//
//	payer, err := solana.PrivateKeyFromBase58(payerPrvKey)
//	if err != nil {
//		panic(err)
//	}
//
//	instruction, err := contract.MintFTInstruction(programId, bridgeAdmin, mint.PublicKey(), payer.PublicKey(), args)
//	if err != nil {
//		panic(err)
//	}
//
//	blockhash, err := Client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
//	if err != nil {
//		panic(err)
//	}
//
//	tx, err := solana.NewTransaction(
//		[]solana.Instruction{
//			instruction,
//		},
//		blockhash.Value.Blockhash,
//		solana.TransactionPayer(payer.PublicKey()),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = tx.AddSignature(payer, mint)
//	if err != nil {
//		panic(err)
//	}
//
//	binTx, err := tx.MarshalBinary()
//	if err != nil {
//		panic(err)
//	}
//
//	Submit(binTx)
//}
