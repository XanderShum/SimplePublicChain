package BLC

import "time"

type Block struct {
	//区块高度
	Height int64
	//上一个区块的hash
	PreBlockHash []byte
	//交易数据
	Data []byte
	//时间戳
	Timestamp int64
	//本区块hash
	Hash []byte
}

//创建新的区块
func NewBlock(data []byte, height int64, prevBlockHash []byte) *Block {
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}

	//设置Hash
	//	block.SetHash()
	return block
}
