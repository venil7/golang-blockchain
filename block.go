package golang_blockchain

type Hash [32]byte
type Nonce []byte

type Block struct {
	data       []byte
	nonce      Nonce
	sum        Hash
	difficulty byte
	parent     *Block
}

func Genesis(data *[]byte) *Block {
	return New(data, nil, 0)
}

func New(data *[]byte, parent *Block, difficulty byte) *Block {
	block := new(Block)
	block.difficulty = difficulty
	block.parent = parent
	block.data = make([]byte, len(*data))
	copy(block.data, *data)
	block.mine()
	return block
}

func (block *Block) mine() *Block {
	nonce := Nonce{}
	for {
		sum := hash(block, block.difficulty, nonce)
		if sum.MatchesDifficulty(block.difficulty) {
			block.sum = sum
			block.nonce = nonce
			return block
		}
	}
}

func (block *Block) Append(data *[]byte, difficulty byte) *Block {
	return New(data, block, difficulty)
}

func (block *Block) VerifyBlock() bool {
	return block.sum == hash(block, block.difficulty, block.nonce)
}

func (block *Block) VerifyIntegrity() bool {
	verifyBlock := block.VerifyBlock()
	if block.parent == nil {
		return verifyBlock
	}
	return verifyBlock && block.parent.VerifyIntegrity()
}
