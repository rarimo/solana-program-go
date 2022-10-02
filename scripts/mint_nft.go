package scripts

//
//func MintNFTCollection(adminSeed, program, payerPrvKey string) {
//	seed := getSeedFromString(adminSeed)
//
//	args := contract.MintNFTArgs{
//		Data: metaplex.DataV2{
//			Name:   "Test NFT",
//			Symbol: "TNFT",
//			URI:    "google.com",
//		},
//		Seeds:  seed,
//		Verify: false,
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
//	fmt.Println("NFT token mint: " + mint.PublicKey().String())
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
//	instruction, err := contract.MintNFTInstruction(programId, bridgeAdmin, mint.PublicKey(), payer.PublicKey(), args)
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
//
//func MintNFT(adminSeed, program, coll, payerPrvKey string) {
//	seed := getSeedFromString(adminSeed)
//
//	collection, err := solana.PublicKeyFromBase58(coll)
//	if err != nil {
//		panic(err)
//	}
//
//	args := contract.MintNFTArgs{
//		Data: metaplex.DataV2{
//			Name:   "Test NFT from collection",
//			Symbol: "TNFT",
//			URI:    "google.com",
//			Collection: &metaplex.Collection{
//				Address: collection,
//			},
//		},
//		Seeds:  seed,
//		Verify: true,
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
//	fmt.Println("NFT token mint: " + mint.PublicKey().String())
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
//	instruction, err := contract.MintNFTVerifiedInstruction(programId, bridgeAdmin, mint.PublicKey(), payer.PublicKey(), collection, args)
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
