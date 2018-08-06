package block_chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	TimeStamp     int64
	TargetBits    int64
	Hash          []byte ///为了方便实现而做的简化，正常比特币节点不包含Hash

	Nonce      int64
	MerKelRoot []byte
	Data       []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		TimeStamp:     time.Now().Unix(),
		TargetBits:    targetBits,
		Nonce:         0,
		MerKelRoot:    []byte{},
		Data:          []byte(data),
	}
	//block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block

}

func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		IntToByte(block.TargetBits),
		IntToByte(block.Nonce),
		block.MerKelRoot,
		block.Data,
	}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]

}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (block *Block) Serialize() []byte {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr("serialize", err)
	return buffer.Bytes()

}

func Deserialize(data []byte) *Block {
	decoder := gob.NewDecoder(bytes.NewReader(data))

	var block Block
	err := decoder.Decode(block)
	CheckErr("Deserialize", err)

	return &block

}
