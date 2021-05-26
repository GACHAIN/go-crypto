// MIT License
//
// Copyright (c) 2016-2021 GACHAIN
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ecies

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/GACHAIN/go-crypto/crypto"
)

// HexToBytes converts the hexadecimal representation to []byte
func HexToBytes(hexdata string) ([]byte, error) {
	return hex.DecodeString(hexdata)
}

func TestEccencryptoKey(t *testing.T) {
	plainText := []byte("ecc hello")

	privateHex := "5d2275c0888d1576e15a45b7eeb870b26a45ceb89f37e586ee21b07c14b0541a"
	pubkeyHex := "0463fbbfefe076637384717297f9f09951e8a2a02480b14cfbd1ed4050ff07d2882a67212dce487ed5cee93fcc3126e9197b73eea02d2a73c64a4906ece24fad67"

	privateKeyBytes, err := HexToBytes(privateHex)
	if err != nil {
		fmt.Println(err)
	}

	publicKeyBytes, err := crypto.HexToPub(pubkeyHex)
	if err != nil {
		fmt.Println(err)
	}

	pub, err2 := crypto.GetPublicKeys(publicKeyBytes)
	if err2 != nil {
		fmt.Println(err2)
	}

	pri, err2 := crypto.GetPrivateKeys(privateKeyBytes)
	if err2 != nil {
		fmt.Println(err2)
	}

	cryptText, _ := EccPubEncrypt(plainText, pub)
	fmt.Println("ECC：", hex.EncodeToString(cryptText))

	msg, err := EccPriDeCrypt(cryptText, pri)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ECC：", string(msg))

}
