package golang_blockchain_test

import (
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
