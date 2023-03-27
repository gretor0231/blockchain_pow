package main

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
	b := NewBlock(0, "init hash")
	b.Print()
}
