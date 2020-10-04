package golang_blockchain

import sha "crypto/sha256"

func hash(block *Block, difficulty byte, nonce Nonce) Hash {
	parentsum := Hash{}
	if block.parent != nil {
		parentsum = block.parent.sum
	}
	difflen := 1
	parentsumlen := len(parentsum)
	noncelen := len(nonce)
	datalen := len((block.data))
	data := make([]byte, difflen+parentsumlen+noncelen+datalen)
	numbytes1 := copy(data, []byte{difficulty})
	numbytes2 := copy(data[numbytes1:], parentsum[:])
	numbytes3 := copy(data[numbytes1+numbytes2:], nonce)
	numbytes4 := copy(data[numbytes1+numbytes2+numbytes3:], block.data)
	if numbytes1+numbytes2+numbytes3+numbytes4 != len(data) {
		panic("lengths dont match")
	}
	return sha.Sum256(data)
}

func (hash *Hash) MatchesDifficulty(difficulty byte) bool {
	for _, b := range hash[0:difficulty] {
		if b != 0 {
			return false
		}
	}
	return true
}
