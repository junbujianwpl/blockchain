package block_chain

import (
	"fmt"
	)

func main()  {

	bc:=NewBlockChain()

	bc.AddBlock("班长给老师一枚BTC")
	bc.AddBlock("班长又给老师一枚BTC")

	for i,block := range bc.Blocks {

		fmt.Print("block num",i)
		fmt.Print("version",block.Version)
		fmt.Print("Data",block.Data)
		fmt.Print("PrevBlockHash",block.PrevBlockHash)
		fmt.Print("timestamp",block.TimeStamp)
		fmt.Print("Nonce",block.Nonce)
		fmt.Print("targetbit",block.TargetBits)
		fmt.Print("Merkel root",block.MerKelRoot)
	}

}
