package main

import "block_chain"

func main() {

	bc := block_chain.NewBlockChain()
	cli := block_chain.CLI{bc}
	cli.Run()
	//for {
	//	block:=it.Next()
	//
	//	fmt.Println("block num",i)
	//	fmt.Println("version",block.Version)
	//	fmt.Println("Data",block.Data)
	//	fmt.Println("PrevBlockHash",block.PrevBlockHash)
	//	fmt.Println("timestamp",block.TimeStamp)
	//	fmt.Println("Nonce",block.Nonce)
	//	fmt.Println("targetbit",block.TargetBits)
	//	fmt.Println("Merkel root",block.MerKelRoot)
	//	pow:=block_chain.NewProofOfWork(block)
	//	fmt.Println(pow.IsValid())
	//}

}
