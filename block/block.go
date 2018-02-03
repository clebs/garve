package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents the data necessary to define a block in the chain.
type Block struct {
	Index     int
	Timestamp string
	Content   fmt.Stringer
	Hash      string
	PrevHash  string
}

// New creates a new block from the previous one and links their hashes.
func New(previous Block, content fmt.Stringer) (Block, error) {
	var newBlock Block

	newBlock.Index = previous.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Content = content
	newBlock.PrevHash = previous.Hash
	newBlock.Hash = newBlock.calculateHash()

	return newBlock, nil
}

// HasValidHash recalculates the hash of the block and checks it against the current hash.
func (b Block) HasValidHash() bool {
	return b.calculateHash() == b.Hash
}

// calculateHash generates the hash of the given block with its data.
func (b Block) calculateHash() string {
	record := string(b.Index) + b.Timestamp + b.Content.String() + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
