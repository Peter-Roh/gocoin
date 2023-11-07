package main

import "github.com/Peter-Roh/gocoin/blockchain"

func main() {
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
	blockchain.Blockchain().AddBlock("Fourth")
}
