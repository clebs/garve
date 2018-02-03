package chain

import (
	"github.com/clebs/garve/block"
)

// Chain is a sequence of blocks
type Chain []block.Block

// Blockchain is the central chain and source of truth
var Blockchain Chain

// replaceChain replaces the current Blockchain with the given chain if it is longer
func replaceChain(newBlocks Chain) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

// CanAppend checks if the given block can be appended at the end of the chain.
func (ch Chain) CanAppend(candidate block.Block) bool {
	lastBlock := ch.Top()

	if lastBlock.Index+1 != candidate.Index {
		return false
	}

	if lastBlock.Hash != candidate.PrevHash {
		return false
	}

	if !candidate.HasValidHash() {
		return false
	}
	return true
}

// Append appends the given block into the chain
func (ch *Chain) Append(b block.Block) {
	*ch = append(*ch, b)
}

// Top returns the top block in the chain.
func (ch *Chain) Top() block.Block {
	if len(*ch) == 0 {
		ch.genesis()
	}
	return (*ch)[len(*ch)-1]
}

// genesis is the origin of a new chain
func (ch *Chain) genesis() {
	ch.Append(block.Block{})
}
