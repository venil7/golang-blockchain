package golang_blockchain_test

import (
	"reflect"
	"testing"

	bchain "github.com/venil7/golang-blockchain"
)

func TestNonce1(t *testing.T) {
	nonce := bchain.Nonce{}
	res := nonce.Next()
	expected := bchain.Nonce{0}
	if !reflect.DeepEqual(res, expected) {
		t.Fail()
	}
}
func TestNonce2(t *testing.T) {
	nonce := bchain.Nonce{255}
	res := nonce.Next()
	expected := bchain.Nonce{0, 0}
	if !reflect.DeepEqual(res, expected) {
		t.Fail()
	}
}

func TestNonce3(t *testing.T) {
	nonce := bchain.Nonce{1, 2, 3}
	res := nonce.Next()
	expected := bchain.Nonce{2, 2, 3}
	if !reflect.DeepEqual(res, expected) {
		t.Fail()
	}
}

func TestNonce4(t *testing.T) {
	nonce := bchain.Nonce{255, 255, 0}
	res := nonce.Next()
	expected := bchain.Nonce{0, 0, 1}
	if !reflect.DeepEqual(res, expected) {
		t.Fail()
	}
}
