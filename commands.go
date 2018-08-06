package block_chain

import "fmt"

func (cli *CLI) AddBlock(data string)  {
	cli.Bc.AddBlock(data)
	fmt.Println("add block succeed")

}

func (cli *CLI) PrintChain()  {
	bc:=cli.Bc
	it:=bc.Iterator()
	for ; ;  {
		block:=it.Next()//取回当前hash指向的block，将hash值前移

		//fmt.Println("block num",i)
		fmt.Println("version",block.Version)
		fmt.Println("Data",block.Data)
		fmt.Println("PrevBlockHash",block.PrevBlockHash)
		fmt.Println("timestamp",block.TimeStamp)
		fmt.Println("Nonce",block.Nonce)
		fmt.Println("targetbit",block.TargetBits)
		fmt.Println("Merkel root",block.MerKelRoot)
		pow:=NewProofOfWork(block)
		fmt.Println(pow.IsValid())

		if len(block.PrevBlockHash)==0{
			break
		}

	}

}
