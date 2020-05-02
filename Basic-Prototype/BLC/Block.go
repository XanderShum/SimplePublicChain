package BLC

import (
	"time"
)

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
	//nouce
	Nonce int64
}

//创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

//创建创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
