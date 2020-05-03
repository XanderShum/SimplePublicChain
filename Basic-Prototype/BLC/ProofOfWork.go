package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex(pow.Block.Height),
			pow.Block.PreBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(nonce),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)

	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	var nonce int64 = 0
	var hashInt big.Int
	var hash [32]byte

	for {
		dataBytes := pow.prepareData(nonce)

		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Block里面的target
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}
	return hash[:], nonce
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
