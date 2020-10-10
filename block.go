package golang_blockchain

import "encoding/hex"

type Hash [32]byte
type Nonce []byte

type Block struct {
	Data       []byte
	Nonce      Nonce
	Sum        Hash
	Difficulty byte
	Parent     *Block
}

func Genesis(data *[]byte) *Block {
	return New(data, nil, 0)
}

func New(data *[]byte, parent *Block, difficulty byte) *Block {
	block := new(Block)
	block.Difficulty = difficulty
	block.Parent = parent
	block.Data = make([]byte, len(*data))
	copy(block.Data, *data)
	block.mine()
	return block
}

func (block *Block) mine() *Block {
	nonce := Nonce{}
	for {
		sum := hash(block, block.Difficulty, nonce)
		if sum.MatchesDifficulty(block.Difficulty) {
			block.Sum = sum
			block.Nonce = nonce
			return block
		}
		nonce = nonce.Next()
	}
}

func (block *Block) Append(data *[]byte, difficulty byte) *Block {
	return New(data, block, difficulty)
}

func (block *Block) VerifyBlock() bool {
	sum := hash(block, block.Difficulty, block.Nonce)
	return block.Sum == sum && sum.MatchesDifficulty(block.Difficulty)
}

func (block *Block) VerifyIntegrity() bool {
	verifyBlock := block.VerifyBlock()
	if block.Parent == nil {
		return verifyBlock
	}
	return verifyBlock && block.Parent.VerifyIntegrity()
}

func (block *Block) Hash() string {
	return hex.EncodeToString(block.Sum[:])
}

func (block *Block) Bytes() []byte {
	return bytes(block, block.Difficulty, block.Nonce)
}
