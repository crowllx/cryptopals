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

func scoreChar(char byte) int {
    char_scores := map[byte]int{
        byte('e'): 13,
        byte('a'): 13,
        byte('r'): 12,
        byte('i'): 12,
        byte('o'): 11,
        byte('t'): 11,
        byte('n'): 10,
        byte('s'): 10,
        byte('l'): 9,
        byte('c'): 9,
        byte('u'): 8,
        byte('d'): 8,
        byte('p'): 7,
        byte('m'): 7,
        byte('h'): 6,
        byte('g'): 6,
        byte('b'): 5,
        byte('f'): 5,
        byte('y'): 4,
        byte('w'): 4,
        byte('k'): 3,
        byte('v'): 3,
        byte('x'): 2,
        byte('z'): 2,
        byte('j'): 1,
        byte('q'): 1,
    }
    return char_scores[char]
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
