package main

import (
	"log"
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

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
}
