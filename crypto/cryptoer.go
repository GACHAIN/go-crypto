package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/GACHAIN/go-crypto/consts"
	"github.com/GACHAIN/go-crypto/converter"

	log "github.com/sirupsen/logrus"
	"github.com/tjfoc/gmsm/sm2"
)

type Cryptoer interface {
	// GenBytesKeys generates a random pair of private and public binary keys.
	genBytesKeys() ([]byte, []byte, error)
	// Sign in signing data with private key
	sign(privateKey, data []byte) ([]byte, error)
	// Verify is the corresponding function for CheckSign in crypto.go
	checkSign(public, data, signature []byte) (bool, error)
}

type SM2 struct {
	Cryptoer
}

func (s *SM2) genBytesKeys() ([]byte, []byte, error) {
	priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		return nil, nil, err
	}
	return priv.D.Bytes(), append(converter.FillLeft(priv.PublicKey.X.Bytes()), converter.FillLeft(priv.PublicKey.Y.Bytes())...), nil
}

func (s *SM2) sign(privateKey, data []byte) ([]byte, error) {
	if len(data) == 0 {
		log.WithFields(log.Fields{"type": consts.CryptoError}).Debug(ErrSigningEmpty.Error())
	}
	fmt.Println("SM2 Sign")
	return signSM2(privateKey, data)
}

func signSM2(privateKey, data []byte) ([]byte, error) {
	// var pubkeyCurve elliptic.Curve
	pubkeyCurve := sm2.P256Sm2()
	// pubkeyCurve.ScalarBaseMult(data)
	bi := new(big.Int).SetBytes(privateKey)
	fmt.Println("bi is:", bi)
	priv := new(sm2.PrivateKey)
	// fmt.Println("priv is 0:", priv) // nil for now
	priv.PublicKey.Curve = pubkeyCurve
	fmt.Println("priv is 1:", priv)
	priv.D = bi
	fmt.Println("priv is 2:", priv)
	priv.PublicKey.X, priv.PublicKey.Y = pubkeyCurve.ScalarBaseMult(bi.Bytes())
	ret, err := priv.Sign(rand.Reader, data, nil)
	return ret, err
}

func (s *SM2) checkSign(public, data, signature []byte) (bool, error) {
	if len(public) == 0 {
		log.WithFields(log.Fields{"type": consts.CryptoError}).Debug(ErrCheckingSignEmpty.Error())
	}
	return checkSM2(public, data, signature)
}

func checkSM2(public, data, signature []byte) (bool, error) {
	if len(data) == 0 {
		log.WithFields(log.Fields{"type": consts.CryptoError}).Error("data is empty")
		return false, fmt.Errorf("invalid parameters len(data) == 0")
	}
	if len(public) != consts.PubkeySizeLength {
		log.WithFields(log.Fields{"size": len(public), "size_match": consts.PubkeySizeLength, "type": consts.SizeDoesNotMatch}).Error("invalid public key")
		return false, fmt.Errorf("invalid parameters len(public) = %d", len(public))
	}
	if len(signature) == 0 {
		log.WithFields(log.Fields{"type": consts.CryptoError}).Error("invalid signature")
		return false, fmt.Errorf("invalid parameters len(signature) == 0")
	}

	pubkeyCurve := sm2.P256Sm2()
	pubkey := new(sm2.PublicKey)
	pubkey.Curve = pubkeyCurve
	pubkey.X = new(big.Int).SetBytes(public[0:consts.PrivkeyLength])
	pubkey.Y = new(big.Int).SetBytes(public[consts.PrivkeyLength:])
	verifystatus := pubkey.Verify(data, signature)
	if !verifystatus {
		return false, ErrIncorrectSign
	}
	return true, nil
}

type ECDSA struct {
	Cryptoer
}

func (s *ECDSA) genBytesKeys() ([]byte, []byte, error) {
	return genBytesKeys()
}

func (s *ECDSA) sign(privateKey, data []byte) ([]byte, error) {
	return sign(privateKey, data)
}

func (s *ECDSA) checkSign(public, data, signature []byte) (bool, error) {
	return checkSign(public, data, signature)
}

type Oval struct {
	name string
}

const (
	cSM2   = "SM2"
	cECDSA = "ECDSA"
)

var Curve = &curve

var curve Oval

func (o Oval) String() string {
	return o.name
}

func getCryptoer() (Cryptoer, error) {
	switch curve.name {
	case cSM2:
		return &SM2{}, nil
	case cECDSA:
		return &ECDSA{}, nil
	}
	return nil, fmt.Errorf("not support crypto")
}

func GenBytesKeys() ([]byte, []byte, error) {
	c, err := getCryptoer()
	if err != nil {
		return nil, nil, err
	}
	return c.genBytesKeys()
}

func Sign(privateKey, data []byte) ([]byte, error) {
	c, err := getCryptoer()
	if err != nil {
		return nil, err
	}
	return c.sign(privateKey, data)
}

func CheckSign(public, data, signature []byte) (bool, error) {
	c, err := getCryptoer()
	if err != nil {
		return false, err
	}
	return c.checkSign(public, data, signature)
}

func InitCurve(s string) {
	switch s {
	case cECDSA:
		curve.name = cECDSA
		return
	case cSM2:
		curve.name = cSM2
		return
	}
	panic(fmt.Errorf("not support crypto"))
}
