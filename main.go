package main

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric/protos/common"
	"io/ioutil"
	"log"
	"os"
)

func genDataHash(block *cb.Block) string {
	calcDataHash := base64.StdEncoding.EncodeToString(block.Data.Hash())

	return calcDataHash
}

func getBlock(blockData []byte) *cb.Block {
	block := &cb.Block{}

	err := proto.Unmarshal(blockData, block)
	if err != nil {
		log.Panic(err)
	}

	return block
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	return data

}
func main() {
	if len(os.Args) != 2 {
		panic("Wrong argument number.")
	}

	args := os.Args[1:]

	path := args[0]
	blockData := readFile(path)
	block := getBlock(blockData)

	calcDataHash := genDataHash(block)
	fmt.Println("Calculated DataHash: ", calcDataHash)

	expectedDataHash := base64.StdEncoding.EncodeToString(block.Header.DataHash)
	fmt.Println("Expected DataHash: ", expectedDataHash)
}
