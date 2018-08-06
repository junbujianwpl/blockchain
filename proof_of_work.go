package block_chain

import (
	"math/big"
	"encoding/gob"
	"bytes"
	"fmt"
	"math"
	"crypto/sha256"
	)

const targetBits=24

type ProofOfWork struct {

	block *Block
	targetBit *big.Int

}

func NewProofOfWork(block *Block) *ProofOfWork {

	var intTargit= new(big.Int).SetInt64(100)

	intTargit.Lsh(intTargit,uint(256-targetBits))


	return &ProofOfWork{block,intTargit }

}

func (pow *ProofOfWork) PrepareRawData(nonce int64) []byte {
	block:=pow.block
	block.Nonce=nonce
	block.SetHash()
	var network bytes.Buffer
	enc:=gob.NewEncoder(&network)
	err:=enc.Encode(block)
	CheckErr("newblockchain",err)
	//fmt.Println(network)

	return network.Bytes()
}

func (pow ProofOfWork) Run() (int64, []byte)  {
	var nonce int64
	var hash [32]byte
	var hashInt big.Int
	fmt.Println("开始 挖圹了")
	for nonce < math.MaxInt64{
		data:=pow.PrepareRawData(nonce)
		hash=sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.targetBit)<0{
			fmt.Printf("found hash :%x  hashInt :%s \n",hash,hashInt)
			break
		}else {
			nonce++
		}
	}
	return nonce,hash[:]
}

func (pow *ProofOfWork) IsValid()bool  {
	data:=pow.PrepareRawData(pow.block.Nonce)
	hash :=sha256.Sum256(data)
	var intHash big.Int
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.targetBit)<0
}
