package main

import (
	"fmt"
	"time"
	"crypto/sha256"
)

type Block struct {
	TimeStamp    time.Time
	Data []string
	PreviousHash []byte
	Hash         []byte
}

func createBlock(data []string, previousHash []byte) *Block {
	
	time := time.Now()

	return &Block {
	TimeStamp   : time,
	Data        : data,
	PreviousHash: previousHash,
	Hash		: NewHash(time, previousHash, data),
	}
}

func NewHash(time time.Time, previousHash []byte, data []string) []byte {

	input := append(previousHash, time.String()...)

	for data := range data {
		input = append(input, string(rune(data))...)
	}

	hash := sha256.Sum256(input)

	return hash[:]
}

func printBlock(block *Block) {

	fmt.Printf("Time of creation: %s\n", block.TimeStamp.String())
	fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
	fmt.Printf("Block Hash: %x\n", block.Hash)
}

func main() {
	block1data := []string{"Block 1 Data"}
	block1 := createBlock(block1data, []byte{})
	fmt.Println("Block 1")
	printBlock(block1)

	block2data := []string{"Block 2 Data"}
	block2 := createBlock(block2data, block1.Hash)
	fmt.Println("Block 2")
	printBlock(block2)
	
	block3data := []string{"Block 3 Data"}
	block3 := createBlock(block3data, block2.Hash)
	fmt.Println("Block 3")
	printBlock(block3)
}