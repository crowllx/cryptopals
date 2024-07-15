package set2

import (
	"bytes"
)

func Pksc7Pad(pt []byte, blockSize int) []byte {
	remainder := len(pt) % blockSize
	if remainder != 0 {
		padLen := blockSize - remainder
		pt = append(pt, bytes.Repeat([]byte{byte(padLen)}, padLen)...)
	}
	return pt
}
