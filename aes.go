package main

import (
	"crypto/aes"
	"errors"
	"fmt"
)

func ECBDecrypt(key []byte, data []byte) ([]byte, error) {
	if len(data)%16 != 0 {
		return nil, errors.New("invalid input text length.")
	}
	cipher, _ := aes.NewCipher(key)
	// ct, _ := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(data)
	pt := make([]byte, len(data))
	size := len(key)
	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(pt[bs:be], data[bs:be])
	}
	return pt, nil
}

func ECBEncrypt(key []byte, data []byte) ([]byte, error) {
	// if len(data)%16 != 0 {
	// 	return nil, errors.New("invalid input text length.")
	// }
	pt := Pksc7Pad(data, 16)
	cipher, _ := aes.NewCipher(key)
	ct := make([]byte, len(pt))
	size := len(key)
	for bs, be := 0, size; bs < len(pt); bs, be = bs+size, be+size {
		cipher.Encrypt(ct[bs:be], pt[bs:be])
	}
	return ct, nil
}

func DetectECB(data string) bool {
	blocks := chunkify([]byte(data), aes.BlockSize)
	m := make(map[string]bool)
	for _, b := range blocks {
		fmt.Println(b)
		if m[string(b)] {
			return true
		}
		m[string(b)] = true
	}
	return false

}
