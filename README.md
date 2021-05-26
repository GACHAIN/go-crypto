1. SM2算法的接入实现:
* 代码位置: crypto/cryptoer.go
* 接口设计:
```go
type Cryptoer interface {
	Name() string
	// GenBytesKeys generates a random pair of private and public binary keys.
	GenBytesKeys() ([]byte, []byte, error)
	// Sign in signing data with private key
	Sign(privateKey, data []byte) ([]byte, error)
	// Verify is the corresponding function for CheckSign in crypto.go
	CheckSign(public, data, signature []byte) (bool, error)
}
```
* 接口实现举例:
```go
type SM2 struct {
	name string
	Cryptoer
}
func (s *SM2) Name() string {
	return s.name
}
```
* 调用方法:
```go
func GetCryptoer(t string) (Cryptoer, error) {
}
```
* 调用与测试代码: crypto/cryptoer_test.go


2. SM3算法的接入实现:
* 代码位置: crypto/hasher.go
* 接口设计:
```go
type Hasher interface {
	Name() string
	// GetHMAC returns HMAC hash
	GetHMAC(secret string, message string) ([]byte, error)
	// Hash returns hash of passed bytes
	Hash(msg []byte) ([]byte, error)
	// DoubleHash returns double hash of passed bytes
	DoubleHash(msg []byte) ([]byte, error)
}
```
* 接口实现举例:
```go
type SM3 struct {
	name string
	Hasher
}
func (s *SM3) Name() string {
	return s.name
}
```
* 调用方法:
```go
func GetHasher(t string) (Cryptoer, error) {
}
```
* 调用与测试代码: crypto/hasher_test.go