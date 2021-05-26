package crypto

import (
	"crypto/hmac"
	"fmt"

	"github.com/tjfoc/gmsm/sm3"
)

type Hasher interface {
	// GetHMAC returns HMAC hash
	getHMAC(secret string, message string) ([]byte, error)
	// Hash returns hash of passed bytes
	hash(msg []byte) ([]byte, error)
	// DoubleHash returns double hash of passed bytes
	doubleHash(msg []byte) ([]byte, error)
}

type SM3 struct {
	Hasher
}

type SHA256 struct {
	Hasher
}

func (s *SM3) getHMAC(secret string, message string) ([]byte, error) {
	mac := hmac.New(sm3.New, []byte(secret))
	mac.Write([]byte(message))
	return mac.Sum(nil), nil
}

func (s *SM3) hash(msg []byte) ([]byte, error) {
	return sm3.Sm3Sum(msg), nil
}
func (s *SM3) doubleHash(msg []byte) ([]byte, error) {
	firstHash := sm3.Sm3Sum(msg)
	return sm3.Sm3Sum(firstHash[:]), nil
}

func (s *SHA256) getHMAC(secret string, message string) ([]byte, error) {
	return getHMAC(secret, message)
}
func (s *SHA256) hash(msg []byte) ([]byte, error) {
	return _Hash(msg)
}
func (s *SHA256) doubleHash(msg []byte) ([]byte, error) {
	return doubleHash(msg)
}

type Hval struct {
	name string
}

const (
	hSM3    = "SM3"
	hSHA256 = "SHA256"
)

var hal Hval
var Hal = &hal

func (h Hval) String() string {
	return h.name
}

func getHasher() (Hasher, error) {
	switch hal.name {
	case hSM3:
		return &SM3{}, nil
	case hSHA256:
		return &SHA256{}, nil
	}
	return nil, fmt.Errorf("not support hash")
}

func InitHash(s string) {
	switch s {
	case hSM3:
		hal.name = hSM3
		return
	case hSHA256:
		hal.name = hSHA256
		return
	}
	panic(fmt.Errorf("not support hash"))
}

func GetHMAC(secret string, message string) ([]byte, error) {
	h, err := getHasher()
	if err != nil {
		return nil, err
	}
	return h.getHMAC(secret, message)
}

func Hash(msg []byte) ([]byte, error) {
	h, err := getHasher()
	if err != nil {
		return nil, err
	}
	return h.hash(msg)
}

func DoubleHash(msg []byte) ([]byte, error) {
	h, err := getHasher()
	if err != nil {
		return nil, err
	}
	return h.doubleHash(msg)
}
