package golang_blockchain

import sha "crypto/sha256"

func bytes(block *Block, difficulty byte, nonce Nonce) []byte {
	parentsum := Hash{}
	if block.Parent != nil {
		parentsum = block.Parent.Sum
	}
	difficultylen := 1
	parentsumlen := len(parentsum)
	noncelen := len(nonce)
	datalen := len((block.Data))
	bytes := make([]byte, difficultylen+parentsumlen+noncelen+datalen)
	numbytes1 := copy(bytes, []byte{difficulty})
	numbytes2 := copy(bytes[numbytes1:], parentsum[:])
	numbytes3 := copy(bytes[numbytes1+numbytes2:], nonce)
	numbytes4 := copy(bytes[numbytes1+numbytes2+numbytes3:], block.Data)
	if numbytes1+numbytes2+numbytes3+numbytes4 != len(bytes) {
		panic("lengths dont match")
	}
	return bytes
}

func hash(block *Block, difficulty byte, nonce Nonce) Hash {
	bytes := bytes(block, difficulty, nonce)
	return sha.Sum256(bytes)
}

func (hash *Hash) MatchesDifficulty(difficulty byte) bool {
	for _, b := range hash[0:difficulty] {
		if b != 0 {
			return false
		}
	}
	return true
}
