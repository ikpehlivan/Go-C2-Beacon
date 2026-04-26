package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

var SecretKey = []byte("SECRE-KEY-!!!") // 32 byte and AES-GCM

func Encrypt(data []byte) ([]byte, error) {
	block, _ := aes.NewCipher(SecretKey)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	return gcm.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(data []byte) ([]byte, error) {
	block, _ := aes.NewCipher(SecretKey)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}