package BLC

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChainWithGensisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Block.....")

	return &BlockChain{[]*Block{genesisBlock}}
}

func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	blc.Blocks = append(blc.Blocks, newBlock)
}
