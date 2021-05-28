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

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/GACHAIN/go-crypto/consts"
	log "github.com/sirupsen/logrus"
)

func SignString(privateKeyHex, data string) ([]byte, error) {
	privateKey, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.WithFields(log.Fields{"type": consts.ConversionError, "error": err}).Error("decoding private key from hex")
		return nil, err
	}
	return getCryptoer().sign(privateKey, []byte(data))
}

// GetPrivateKeys return
func GetPrivateKeys(privateKey []byte) (ret *ecdsa.PrivateKey, err error) {
	var pubkeyCurve elliptic.Curve

	switch ellipticSize {
	case elliptic256:
		pubkeyCurve = elliptic.P256()
	default:
		log.WithFields(log.Fields{"type": consts.CryptoError}).Error(ErrUnsupportedCurveSize.Error())
		return
	}

	bi := new(big.Int).SetBytes(privateKey)
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = pubkeyCurve
	priv.D = bi

	return priv, nil
}

// GetPublicKeys return
func GetPublicKeys(public []byte) (*ecdsa.PublicKey, error) {

	pubkey := new(ecdsa.PublicKey)

	if len(public) != consts.PubkeySizeLength {
		log.WithFields(log.Fields{"size": len(public), "size_match": consts.PubkeySizeLength, "type": consts.SizeDoesNotMatch}).Error("invalid public key")
		return pubkey, fmt.Errorf("invalid parameters len(public) = %d", len(public))
	}

	var pubkeyCurve elliptic.Curve
	switch ellipticSize {
	case elliptic256:
		pubkeyCurve = elliptic.P256()
	default:
		log.WithFields(log.Fields{"type": consts.CryptoError}).Error(ErrUnsupportedCurveSize.Error())
	}

	pubkey.Curve = pubkeyCurve
	pubkey.X = new(big.Int).SetBytes(public[0:consts.PrivkeyLength])
	pubkey.Y = new(big.Int).SetBytes(public[consts.PrivkeyLength:])

	return pubkey, nil
}
