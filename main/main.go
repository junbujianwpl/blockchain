package main

import "fmt"
import "../../block_chain"

func main()  {
	fmt.Println("hello")
	block_chain.NewBlockChain()
	bc:=block_chain.NewBlockChain()

	bc.AddBlock("班长给老师一枚BTC")
	bc.AddBlock("班长又给老师一枚BTC")


	for i,block := range bc.Blocks {

		fmt.Println("block num",i)
		fmt.Println("version",block.Version)
		fmt.Println("Data",block.Data)
		fmt.Println("PrevBlockHash",block.PrevBlockHash)
		fmt.Println("timestamp",block.TimeStamp)
		fmt.Println("Nonce",block.Nonce)
		fmt.Println("targetbit",block.TargetBits)
		fmt.Println("Merkel root",block.MerKelRoot)
		pow:=block_chain.NewProofOfWork(block)
		fmt.Println(pow.IsValid())
	}

}
