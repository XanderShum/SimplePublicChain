package main

import (
	"fmt"
	"go_project/SimplePublicChain/Basic-Prototype/BLC"
)

func main() {
	BlockChain := BLC.CreateBlockChainWithGensisBlock()

	BlockChain.AddBlockToBlockChain("Send 100RMB To sxz", BlockChain.Blocks[len(BlockChain.Blocks)-1].Height+1, BlockChain.Blocks[len(BlockChain.Blocks)-1].Hash)

	BlockChain.AddBlockToBlockChain("Send 200RMB To zyy", BlockChain.Blocks[len(BlockChain.Blocks)-1].Height+1, BlockChain.Blocks[len(BlockChain.Blocks)-1].Hash)

	BlockChain.AddBlockToBlockChain("Send 300RMB To cnm", BlockChain.Blocks[len(BlockChain.Blocks)-1].Height+1, BlockChain.Blocks[len(BlockChain.Blocks)-1].Hash)

	BlockChain.AddBlockToBlockChain("Send 50RMB To abc", BlockChain.Blocks[len(BlockChain.Blocks)-1].Height+1, BlockChain.Blocks[len(BlockChain.Blocks)-1].Hash)

	fmt.Println(BlockChain)
	fmt.Println(BlockChain.Blocks)
}
