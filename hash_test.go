package golang_blockchain_test

import (
	"fmt"
	"testing"

	bchain "github.com/venil7/golang-blockchain"
)

func TestHash1(t *testing.T) {
	hash := bchain.Hash{0, 0, 1}
	if !hash.MatchesDifficulty(2) {
		t.Fail()
	}
}
func TestHash2(t *testing.T) {
	hash := bchain.Hash{0, 0, 1}
	if hash.MatchesDifficulty(3) {
		t.Fail()
	}
}

func TestHash3(t *testing.T) {
	data := []byte("aaa")
	block1 := bchain.Genesis(&data)
	block2 := block1.Append(&data, 3)
	fmt.Print(block2.VerifyIntegrity())
	// if hash.MatchesDifficulty(3) {
	// 	t.Fail()
	// }
}
