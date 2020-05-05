package main

import (
	"go_project/SimplePublicChain/Basic-Prototype/BLC"
)

func main() {
	// 创世区块
	blockchain := BLC.CreateBlockChainWithGensisBlock()
	defer blockchain.DB.Close()

	//新区块
	blockchain.AddBlockToBlockChain("Send 100RMB To sxz")

	blockchain.AddBlockToBlockChain("Send 200RMB To zyy")

	blockchain.AddBlockToBlockChain("Send 300RMB To xtu")

	blockchain.AddBlockToBlockChain("Send 50RMB To wnm")

	blockchain.PrintChain()
}
