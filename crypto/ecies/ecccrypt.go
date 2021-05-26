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
	"crypto/ecdsa"
	"crypto/rand"
	"log"
	"runtime"

	"github.com/GACHAIN/go-crypto/crypto"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

//
//Ecc
/*func EccEnCrypt(plainText []byte,prv2 *ecies.PrivateKey)(crypText []byte,err error){

	ct, err := ecies.Encrypt(rand.Reader, &prv2.PublicKey, plainText, nil, nil)
	return ct, err
}
//
func EccDeCrypt(cryptText []byte,prv2 *ecies.PrivateKey) ([]byte, error) {
	pt, err := prv2.Decrypt(cryptText, nil, nil)
	return pt, err
}*/

//
func EccPubEncrypt(plainText []byte, pub *ecdsa.PublicKey) (cryptText []byte, err error) { //

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "check key ")
			default:
				log.Println("error:", err)
			}
		}
	}()

	publicKey := ImportECDSAPublic(pub)
	//
	crypttext, err := Encrypt(rand.Reader, publicKey, plainText, nil, nil)

	return crypttext, err

}

//
func EccPriDeCrypt(cryptText []byte, priv *ecdsa.PrivateKey) (msg []byte, err error) { //
	privateKey := ImportECDSA(priv)

	//
	plainText, err := privateKey.Decrypt(cryptText, nil, nil)

	return plainText, err
}

func EccCryptoKey(plainText []byte, publickey string) (cryptoText []byte, err error) {
	pubbuff, err := crypto.HexToPub(publickey)
	if err != nil {
		return nil, err
	}
	pub, err := crypto.GetPublicKeys(pubbuff)
	if err != nil {
		return nil, err
	}
	return EccPubEncrypt(plainText, pub)
}

func EccDeCrypto(cryptoText []byte, prikey []byte) ([]byte, error) {
	pri, err := crypto.GetPrivateKeys(prikey)
	if err != nil {
		return nil, err
	}
	return EccPriDeCrypt(cryptoText, pri)
}
