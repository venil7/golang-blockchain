package golang_blockchain

// type Nonce []byte

func (nonce *Nonce) Next() Nonce {
	len := len(*nonce)
	if len == 0 {
		return Nonce{0}
	}
	next := make(Nonce, len)
	copy(next, *nonce)
	index := 0
	for {
		if next[index] < 255 {
			next[index]++
			return next
		}
		next[index] = 0
		if index == len-1 {
			return append(next, 0)
		}
		index++
	}
}
