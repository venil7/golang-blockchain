package golang_blockchain

// import (
// 	b "bytes"
// )

// func (block *Block) Serialize() []byte {
// 	datalen := uint32(len(block.data))
// 	noncelen := uint32(len(block.nonce))
// 	var buffer b.Buffer
// 	buffer.Write([]byte(datalen))
// 	buffer.Write(block.data)
// 	buffer.Write(noncelen)
// 	buffer.Write(block.nonce)
// 	buffer.Write(block.nonce)
// 	buffer.Write(block.sum[:])
// 	buffer.WriteByte(block.difficulty)
// 	return buffer.Bytes()
// }
