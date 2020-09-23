package bchain_test

import (
	"testing"

	"github.com/venil7/bchain"
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
