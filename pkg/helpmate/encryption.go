package helpmate

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var (
	initialVector = "1234567890123456"
)

func AESEncryptCBC(content, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("key error %s", err)
	}
	if content == nil {
		return nil, fmt.Errorf("can't blank")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	//content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	encrypted := make([]byte, len(content))
	ecb.CryptBlocks(encrypted, content)

	return encrypted, nil
}

func AESDecryptCBC(crypt, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Bank Account No encrypt KEY error")
	}
	if len(crypt) == 0 {
		return nil, fmt.Errorf("content can't empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
