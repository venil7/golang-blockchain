package main

import (
	"crypto/rand"
	"fmt"
	"time"

	bchain "github.com/venil7/golang-blockchain"
)

func main() {
	data := make([]byte, 10)
	rand.Read(data) //populate with random data
	head := bchain.Genesis(&data)
	for i := 0; i <= 3; i++ {
		start := time.Now()
		fmt.Printf("Mining next block with difficulty %v \n", i)
		rand.Read(data)
		head = head.Append(&data, byte(i))
		duration := time.Since(start)
		fmt.Printf("Hash found %v, took %v \n", head.Hash(), duration)
	}
	fmt.Printf("Verifying chain, valid: %v \n", head.VerifyIntegrity())
}
