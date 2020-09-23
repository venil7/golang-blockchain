package golang_blockchain_test

import (
	"testing"

	bchain "github.com/venil7/golang-blockchain"
)

func TestBlock1(t *testing.T) {
	block1 := bchain.Genesis(&[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	if !block1.VerifyBlock() {
		t.Fatal("genesis block failed to self verify")
	}
}

func TestBlock2(t *testing.T) {
	block1 := bchain.Genesis(&[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	block2 := block1.Append(&[]byte{1, 2, 3})
	if !block2.VerifyIntegrity() {
		t.Fatal("chain failed to self verify")
	}
}
