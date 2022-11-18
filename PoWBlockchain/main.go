package main 

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"
	"encoding/binary"
)

type Block struct {
	PreviousHash []byte
	Data         []byte
	Hash         []byte
	Nonce        int
}

type Blockchain struct  {
	Blocks []*Block
}

func createNewBlock(data string, previousHash []byte) *Block {
	block := &Block {previousHash, []byte(data), []byte{}, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
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

//--------------------------------Proof Of Work-------------------------------------//

type ProofOfWork struct {
	Block *Block
	target *big.Int
}

const Difficulty = 10

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}

func ToBytes(number int64) []byte {
	var buffer = new(bytes.Buffer)
	error := binary.Write(buffer, binary.BigEndian, number)
	if error != nil {
		log.Fatal(error)
	}
	return buffer.Bytes()
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join([][]byte{pow.Block.Data, pow.Block.PreviousHash, ToBytes(int64(nonce)), ToBytes(int64(Difficulty))}, []byte{})
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var initHash *big.Int
	var hash [32]byte
	var nonce = 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("%x", hash)
		initHash.SetBytes(hash[:])
		if initHash.Cmp(pow.target) == -1 {
			break;
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}



