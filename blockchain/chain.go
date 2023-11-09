// Package blockchain provides blockchain and functions needed.
package blockchain

import (
	"sync"

	"github.com/Peter-Roh/gocoin/db"
	"github.com/Peter-Roh/gocoin/utils"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}

	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newest := allBlocks[0]
	lastRecalculated := allBlocks[difficultyInterval-1]
	actualTime := (newest.Timestamp / 60) - (lastRecalculated.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

// Blockchain guarantees creating genesis block just once by using singleton pattern and returns the blockchain.
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				NewestHash:        "",
				Height:            0,
				CurrentDifficulty: defaultDifficulty,
			}
			persisted := db.Blockchain()
			if persisted == nil {
				b.AddBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks.")
			} else {
				b.restore(persisted)
			}
		})
	}
	return b
}
