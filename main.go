package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain : contains All the blocks of the chain
type BlockChain struct {
	Blocks []*Block
}

// Block : contains Hash, Data, PrevHash
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash : Create Hash from data and previous Hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// AddBlock : Add a Block to the BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// CreateBlock : Create a Block with data and the Hash of the previous Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Genesis : Create first Block of the blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain : Initialize Blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	blockchain := InitBlockChain()

	blockchain.AddBlock("First Block after Genesis")
	blockchain.AddBlock("Second Block after Genesis")
	blockchain.AddBlock("Third Block after Genesis")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Previsous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
