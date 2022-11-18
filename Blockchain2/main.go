package main 

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PreviousHash []byte
	Data         []byte
	Hash         []byte
}

type Blockchain struct  {
	Blocks []*Block
}

func (b *Block) createHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createNewBlock(data string, previousHash []byte) *Block {
	block := &Block {previousHash, []byte(data), []byte{}}
	block.createHash()
	return block
}

func Chain() *Blockchain {
	return &Blockchain{[]*Block{createNewBlock("Genesis", []byte{})}}
}

func (chain *Blockchain) addNewBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := createNewBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func main() {
	chain := Chain()
	chain.addNewBlock("Second Block")
	chain.addNewBlock("Third Block")
	chain.addNewBlock("Fourth Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Println("---------------------------------------------------------------------")
	}
}