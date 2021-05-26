package main

import (
	"fmt"

	"github.com/GACHAIN/go-crypto/crypto"
)

func main() {
	fmt.Println(crypto.GenBytesKeys())
	fmt.Println(crypto.Curve.String())
	fmt.Println(crypto.Hash([]byte("Hello")))
	fmt.Println(crypto.Hal.String())
}

func init() {
	crypto.InitCurve("SM2")
	crypto.InitHash("SM3")
	//crypto.InitCurve("s")
	// crypto.InitCurve(crypto.SM2.String())
}
