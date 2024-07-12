package set1

import (
	"encoding/hex"
)

func xor(input []byte, c byte) []byte {
	var xor []byte
	for _, b := range input {
		xor = append(xor, b^c)
	}
	return xor
}
func SingleByteXor(input string) ([]rune, []string) {
	var xors []string
	var keys []rune

	hex, _ := hex.DecodeString(input)
	for c := 'a'; c < 'z'; c++ {
		xors = append(xors, string(xor(hex, byte(c))))
		keys = append(keys, c)
	}

	return result
}
