package block_chain

import (
		"github.com/boltdb/bolt"
	"os"
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
	CheckErr("newblockchain",err)

	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBuckit))
		if bucket !=nil{
			//读取hash即可
			lastHash=bucket.Get([]byte(lastHash))
		}else {
			//创建bucket
			//与数据
			genesis:=NewGenesisBlock()
			bucket,err:=tx.CreateBucket([]byte(blockBuckit))
			CheckErr("newblockchain",err)
			err=bucket.Put(genesis.Hash,genesis.Serialize())	//todo
			CheckErr("newblockchain",err)
			err=bucket.Put([]byte(lastHash),genesis.Hash)
			CheckErr("newblockchain",err)
			
		}
		return nil

	})



	return &BlockChain{db,lastHash}

}

func (bc *BlockChain)AddBlock(data string)  {
	var prevBlockHash []byte
	err:=bc.db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(blockBuckit))
		lastHash:=bucket.Get([]byte(lastHash))


		prevBlockHash=lastHash

		return  nil

	})

	CheckErr("",err)

	block:=NewBlock(data,prevBlockHash)

	err=bc.db.Update(func(tx *bolt.Tx) error {

		bucket:=tx.Bucket([]byte(blockBuckit))
		err:=bucket.Put(block.Hash,block.Serialize())
		CheckErr("",err)

		err=bucket.Put([]byte(lastHash),block.Hash)
		CheckErr("AddBlock3",err)

		bc.lastHash=block.Hash

		return nil


	})
	CheckErr("",err)

}


type BlockChainIterator struct {
	db *bolt.DB
	currentHash []byte
}

func (bc *BlockChain) Iterator() *BlockChainIterator {

	return &BlockChainIterator{bc.db,bc.lastHash}
}

func (it *BlockChainIterator) Next() *Block {
	var block *Block
	err:=it.db.View(func(tx *bolt.Tx) error {

		bucket:=tx.Bucket([]byte(blockBuckit))

		if  bucket==nil{
			os.Exit(1)

		}
		blockTmp:=bucket.Get(it.currentHash)
		block=Deserialize(blockTmp)
		it.currentHash=block.PrevBlockHash
		return nil
	})
	CheckErr("Next",err)
	return block
}
