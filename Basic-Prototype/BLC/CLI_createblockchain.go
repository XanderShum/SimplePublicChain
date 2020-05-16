package BLC

// 创建创世区块
func (cli *CLI) createGenesisBlockchain(address string) {

	blockchain := CreateBlockChainWithGenesisBlock(address)
	defer blockchain.DB.Close()
}
