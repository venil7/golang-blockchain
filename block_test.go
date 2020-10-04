package golang_blockchain_test

import (
	"testing"

	bchain "github.com/venil7/golang-blockchain"
)

func TestBlock1(t *testing.T) {
	data := []byte("some random data")
	block1 := bchain.Genesis(&data)
	if !block1.VerifyBlock() {
		t.Fatal("genesis block failed to self verify")
	}
}

func TestBlock2(t *testing.T) {
	data1 := []byte("first block data")
	data2 := []byte("second block data")
	block1 := bchain.Genesis(&data1)
	block2 := block1.Append(&data2, 0)
	if !block2.VerifyIntegrity() {
		t.Fatal("chain failed to self verify")
	}
}
