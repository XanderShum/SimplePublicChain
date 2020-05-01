package main

import (
	"fmt"
	"go_project/SimplePublicChain/Basic-Prototype/BLC"
)

func main() {
	genesisBlock := BLC.CreateGenesisBlock("Genesis Block.....")

	fmt.Println(genesisBlock)
}
