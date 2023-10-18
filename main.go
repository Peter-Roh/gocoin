package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getPrevHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{
		data:     data,
		hash:     "",
		prevHash: b.getPrevHash(),
	}
	newBlock.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(newBlock.data+newBlock.prevHash)))
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) printBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("Previous hash: %s\n", block.prevHash)
		fmt.Println("")
	}
}

func main() {
	chain := blockchain{}

	chain.addBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.printBlocks()
}
