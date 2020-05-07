package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type BlockChain struct {
	Tip []byte //最新的区块的Hash
	DB  *bolt.DB
}

// 迭代器
func (blockchain *BlockChain) Iterator() *BlockChainIterator {

	return &BlockChainIterator{blockchain.Tip, blockchain.DB}
}

// 判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

// 遍历输出所有区块的信息
func (blc *BlockChain) PrintChain() {

	fmt.Println("PrintchainPrintchainPrintchainPrintchain")
	blockchainIterator := blc.Iterator()

	for {
		block := blockchainIterator.Next()

		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PreBlockHash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Printf("Timestamp：%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n", block.Hash)
		fmt.Printf("Nonce：%d\n", block.Nonce)

		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)

		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

func CreateBlockChainWithGenesisBlock() *BlockChain {
	if dbExists() {
		fmt.Println("创世区块已经存在......")

		db, err := bolt.Open(dbName, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}

		var blockchain *BlockChain

		err = db.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte(blockTableName))

			hash := b.Get([]byte("l"))

			blockchain = &BlockChain{hash, db}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockchain
	}

	// 创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {

		//  获取表
		b := tx.Bucket([]byte(blockTableName))

		if b == nil {
			// 创建数据库表
			b, err = tx.CreateBucket([]byte(blockTableName))

			if err != nil {
				log.Panic(err)
			}
		}

		if b != nil {
			// 创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.......")
			// 将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// 存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}

		return nil
	})

	// 返回区块链对象
	return &BlockChain{blockHash, db}
}

func (blc *BlockChain) AddBlockToBlockChain(data string) {
	err := blc.DB.Update(func(tx *bolt.Tx) error {

		//1. 获取表
		b := tx.Bucket([]byte(blockTableName))
		//2. 创建新区块
		if b != nil {

			//先获取最新区块
			blockBytes := b.Get(blc.Tip)
			// 反序列化
			block := DeserializeBlock(blockBytes)

			//3. 将区块序列化并且存储到数据库中
			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//4. 更新数据库里面"l"对应的hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5. 更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
