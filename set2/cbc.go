package set2

import (
	"crypto/aes"
	"errors"
)

func xor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("xor error: lengths are not equivalent")
	}
	buf := make([]byte, len(a))
	for i := range a {
		buf[i] = a[i] ^ b[i]
	}
	return buf, nil
}

func Encrypt(pt []byte, key []byte, iv []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("invalid key size")
	}
	pt = Pksc7Pad(pt, 16)
	ct := make([]byte, len(pt))
	block, _ := xor(iv, pt[:16])
	cipher, _ := aes.NewCipher(key)
	cipher.Encrypt(ct[:16], block)
	for bs, be := 16, 32; bs < len(pt); bs, be = bs+16, be+16 {
		block, _ = xor(block, pt[bs:be])
		cipher.Encrypt(ct[bs:be], block)
	}
	return ct, nil
}

func Decrypt(ct []byte, key []byte, iv []byte) ([]byte, error) {
	if len(key) != 16 || len(iv) != 16 {
		return nil, errors.New("invalid key or iv size")
	}
	cipher, _ := aes.NewCipher(key)
	block := make([]byte, 16)
	copy(block, iv)
	var pt []byte
	for bs, be := 0, 16; bs < len(ct); bs, be = bs+16, be+16 {
		cipher.Decrypt(ct[bs:be], ct[bs:be])
		tmp, _ := xor(ct[bs:be], block)
		pt = append(pt, tmp...)
		block = ct[bs:be]
	}
	return pt, nil
}
