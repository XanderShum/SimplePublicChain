package main

import (
	"fmt"
	"go_project/SimplePublicChain/Basic-Prototype/BLC"
)

func main() {
	block := BLC.NewBlock([]byte("Genenis Block"), 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)
}
