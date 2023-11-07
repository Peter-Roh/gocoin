// Package blockchain provides blockchain and functions needed.
package blockchain

import (
	"sync"

	"github.com/Peter-Roh/gocoin/db"
	"github.com/Peter-Roh/gocoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// Blockchain guarantees creating genesis block just once by using singleton pattern and returns the blockchain.
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				NewestHash: "",
				Height:     0,
			}
			b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks.")
		})
	}
	return b
}
