package cli

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/mr-tron/base58"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/rarify-protocol/solana-program-go/scripts"
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	/*	cfg := config.New(kv.MustFromEnv())
		log = cfg.Log()*/

	app := kingpin.New("solana-program-go", "")

	runCmd := app.Command("run", "run command")
	genSeedCmd := runCmd.Command("gen-seed", "generate admins seed")
	genSeedProgramId := genSeedCmd.Flag("program", "program address").String()

	initAdminCmd := runCmd.Command("init-admin", "initialize bridge admin")
	initAdminProgramId := genSeedCmd.Flag("program", "program address").String()
	initAdminSeed := genSeedCmd.Flag("seed", "admin seed").String()
	initAdminPublicKey := genSeedCmd.Flag("public-key", "ECDSA admin public key").String()

	depositNativeCmd := runCmd.Command("deposit-native", "deposit native token")
	sender := depositNativeCmd.Flag("sender", "sender solana address").String()

	withdrawNativeCmd := runCmd.Command("withdraw-native", "withdraw native token")

	depositFTCmd := runCmd.Command("deposit-ft", "deposit fungible token")

	withdrawFTCmd := runCmd.Command("withdraw-ft", "withdraw fungible token")

	depositNFTCmd := runCmd.Command("deposit-nft", "deposit non-fungible token")

	withdrawNFTCmd := runCmd.Command("withdraw-nft", "withdraw non-fungible token")

	mintFTCmd := runCmd.Command("mint-ft", "mint fungible token")

	mintNFTCmd := runCmd.Command("mint-nft", "mint non-fungible token")

	// custom commands go here...

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case genSeedCmd.FullCommand():
		fmt.Println("gen-seed command")
		key, seed := scripts.GenSeed(*genSeedProgramId)
		fmt.Println("BridgeAdmin: " + key.String())
		fmt.Println("Seed: " + base58.Encode(seed[:]))
	case initAdminCmd.FullCommand():
		fmt.Println("init-admin command")
		scripts.InitBridgeAdmin(*initAdminSeed, *initAdminProgramId, *initAdminPublicKey)
	case depositNativeCmd.FullCommand():
		fmt.Println("deposit-native command")
		fmt.Println(*sender)
	case withdrawNativeCmd.FullCommand():
		fmt.Println("withdraw-native command")
	case depositFTCmd.FullCommand():
		fmt.Println("deposit-ft command")
	case withdrawFTCmd.FullCommand():
		fmt.Println("withdraw-ft command")
	case depositNFTCmd.FullCommand():
		fmt.Println("deposit-nft command")
	case withdrawNFTCmd.FullCommand():
		fmt.Println("withdraw-nft command")
	case mintFTCmd.FullCommand():
		fmt.Println("mint-ft command")
	case mintNFTCmd.FullCommand():
		fmt.Println("mint-nft command")
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}

	return true
}
