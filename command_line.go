package block_chain

import (
	"flag"
	"fmt"
	"os"
)

const Usage = `
	./addBlock --data Data "add a block to block chain"
	./block printchain 				"print all blocks"
`

type CLI struct {
	Bc *BlockChain
}

func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("too few args", Usage)
		os.Exit(1)
	}

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data", "", "block info ")
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("addBlock", err)
		if addBlockCmd.Parsed() {
			if *addBlockCmdPara == "" {
				//cli.Bc.AddBlock(addBlockCmdPara)
				cli.AddBlock(*addBlockCmdPara) //todo
			}
		}

	case "printChain":
		err := printCmd.Parse(os.Args[2:])
		CheckErr("printChain", err)
		if printCmd.Parsed() {
			cli.PrintChain() //todo

		}

	default:
		fmt.Println("invalid cmd", Usage)
	}

}
