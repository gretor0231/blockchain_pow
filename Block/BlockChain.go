package main

import (
	"fmt"
	"log"
	"time"
)

type BlockChain struct {
	chain           []*Block
	transactionPool []string
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func NewBlockchain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *BlockChain) Print() {
	for _, b := range bc.chain {
		b.Print()
	}
}

func init() {
	log.SetPrefix("Blockchain: ")
}

//Bock is a block in the blockchain
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

//NewBlock creates a new block
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

//Print prints the block
func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("nonce           %d\n", b.nonce)
	fmt.Printf("previous_hash   %s\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}
