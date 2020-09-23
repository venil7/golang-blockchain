package golang_blockchain

import (
	sha "crypto/sha256"
)

type Block struct {
	data   []byte
	sum    [32]byte
	parent *Block
}

func Genesis(data *[]byte) *Block {
	return New(data, nil)
}

func New(data *[]byte, parent *Block) *Block {
	block := new(Block)
	block.parent = parent
	copy(*data, block.data)
	block.hash()
	return block
}

func (block *Block) hash() *Block {
	block.sum = block.calcHash()
	return block
}

func (block *Block) calcHash() [32]byte {
	var data []byte
	parent := block.parent
	if parent != nil {
		data = append(parent.sum[:], block.data...)
	} else {
		data = block.data
	}
	return sha.Sum256(data)
}

func (block *Block) Append(data *[]byte) *Block {
	return New(data, block)
}

func (block *Block) VerifyBlock() bool {
	return block.sum == block.calcHash()
}

func (block *Block) VerifyIntegrity() bool {
	verifyBlock := block.VerifyBlock()
	if block.parent == nil {
		return verifyBlock
	}
	return verifyBlock && block.parent.VerifyIntegrity()
}
