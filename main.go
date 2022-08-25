package main

import (
	"os"

	"gitlab.com/rarify-protocol/solana-program-go/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
