package main

import (
	"go_project/SimplePublicChain/Basic-Prototype/BLC"
)

func main() {
	blockchain := BLC.CreateBlockChainWithGenesisBlock()

	cli := BLC.CLI{blockchain}

	cli.Run()
}
