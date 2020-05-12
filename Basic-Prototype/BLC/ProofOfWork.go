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
			pow.Block.PreBlockHash,
			pow.Block.HashTransactions(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(nonce),
			IntToHex(pow.Block.Height),
		},
		[]byte{},
	)
	return data
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
