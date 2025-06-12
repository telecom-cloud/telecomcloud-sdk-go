package main

import (
	c "crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/emmansun/gmsm/sm4"
)

// define ecb encrypt
type ecbCipher struct {
	b         c.Block
	blockSize int
}

func (x *ecbCipher) BlockSize() int { return x.blockSize }

func (x *ecbCipher) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		return
	}

	if len(dst) < len(src) {
		return
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b c.Block) c.BlockMode {
	return &ecbCipher{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// decrypt SM4-ECB-PKCS5Padding
func decrypt(data []byte, key []byte) (string, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(data))
	ecb := NewECBDecrypter(block)
	ecb.CryptBlocks(ciphertext, data)

	out := string(pkcs5Unpadding(ciphertext))
	if out == "" {
		return out, fmt.Errorf("decrypt failed, ciphertext is empty")
	}
	return out, nil
}

// pkcs5Unpadding PKCS5 remove padding
func pkcs5Unpadding(data []byte) []byte {
	if len(data) == 0 {
		return make([]byte, 0)
	}

	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

// Decode the data
func Decode(data, key string) (string, error) {
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	return decrypt(dataBytes, keyBytes)
}
