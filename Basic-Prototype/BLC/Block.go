package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

//设置hash
func (block *Block) SetHash() {
	//调用utils的函数将区块高度转化为byte数组
	heightBytes := IntToHex(block.Height)
	//将时间戳转化为string再转成byte
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timebytes := []byte(timeString)
	//将数据拼接
	blockBytes := bytes.Join([][]byte{heightBytes, block.PreBlockHash, block.Data, timebytes}, []byte{})
	//转换为sha256
	hash := sha256.Sum256(blockBytes)
	//设置hash
	block.Hash = hash[:]
}

//创建新的区块
func NewBlock(data []byte, height int64, prevBlockHash []byte) *Block {
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
	//设置Hash
	block.SetHash()
	return block
}

//创建创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock([]byte(data), 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
