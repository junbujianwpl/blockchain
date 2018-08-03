package block_chain

import (
	"os"
	"github.com/boltdb/bolt"
)

const dbfile="blockChainDb.db"
const blockBuckit ="blockBuckit"
const lastHash ="lastHash"

type BlockChain struct {
	//Blocks [] *Block

	db *bolt.DB
	lastHash []byte

}

func NewBlockChain() *BlockChain {

	//return &BlockChain{ []*Block{NewGenesisBlock()},
	//}
	db,err:=bolt.Open(dbfile,0600,nil)
	CheckErr(err)

	db.Update(func(tx *bolt.Tx) error {

		var lastHash []byte
		bucket:=tx.Bucket([]byte(blockBuckit))
		if bucket !=nil{
			//读取hash即可
			lastHash=bucket.Get([]byte(lastHash))
		}else {
			//创建bucket
			//与数据
			genesis:=NewGenesisBlock()
			bucket,err:=tx.CreateBucket([]byte(blockBuckit))
			CheckErr(err)
			err=bucket.Put(genesis.Hash,genesis.Serialize())
			CheckErr(err)
			err=bucket.Put([]byte(lastHash),genesis.Hash)
			CheckErr(err)
			
		}
		return nil

	})



	return &BlockChain{db,lastHash}

}

func (bc *BlockChain)AddBlock(data string)  {
	if len(bc.Blocks)<=0{
		os.Exit(1)
	}
	prevBlockHash :=bc.Blocks[len(bc.Blocks)-1].Hash
	block:=NewBlock(data,prevBlockHash)
	bc.Blocks =append(bc.Blocks, block)

}
