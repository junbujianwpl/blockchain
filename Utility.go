package block_chain

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(num int64)[]byte{
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)
	if err !=nil{
		fmt.Print("Int to byte err occur",err)
		os.Exit(1)

	}
	return  buffer.Bytes()


}

func CheckErr(err error) {
	if err !=nil{
		fmt.Print("err occur:",err)
		os.Exit(1)
	}
}
