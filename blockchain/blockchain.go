// Package blockchain provides blockchain and functions needed.
package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// Block represents a single block.
type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.Data+b.PrevHash)))
}

func getPrevHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{
		Data:     data,
		Hash:     "",
		PrevHash: getPrevHash(),
	}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

// GetBlockchain guarantees creating genesis block just once by using singleton pattern and returns the blockchain.
func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks.")
		})
	}
	return b
}
