package crypto

import (
	"fmt"
	"testing"
)

func TestGetHasher(t *testing.T) {
	sm3, err := GetHasher("sm3")
	if err != nil {
		fmt.Println(err)
	}
	msg := []byte("Hello")
	hashedMsg, err := sm3.Hash(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashedMsg)

	message := "Hello"
	secret := "world"

	hmacMsg, err := sm3.GetHMAC(secret, message)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hmacMsg)
	fmt.Println("name is", sm3.Name())
}
