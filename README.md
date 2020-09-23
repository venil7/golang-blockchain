# A Simple blockchain datastructure implementation in `Golang`

## how to use

```go
  // create a genesis block, pass some bytes
  block1 := bchain.Genesis(&[]byte{1, 2, 3})
  // keep adding blocks, hashes calculate automatically
  block2 := block1.Append(&[]byte{3, 4, 5})
```

## verifications

Hashes calculate automatically, as blocks get added,
at any point is is possiblo to `VerifyBlock()` or `VerifyIntegrity()` of a whole chain

```go
  if !block1.VerifyBlock() {
		t.Fatal("genesis block failed to self verify")
  }

  if !block2.VerifyIntegrity() {
    t.Fatal("chain failed to self verify")
  }
```

## Tests

`go test ./lib_test.go`
