package set1

import (
	"encoding/hex"
	"strings"
)

func xor(input []byte, c byte) []byte {
	var xor []byte
	for _, b := range input {
		xor = append(xor, b^c)
	}
	return xor
}

func scoreXor(xor []byte) int {
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
	score := 0
	xor = []byte(strings.ToLower(string(xor)))
	for _, b := range xor {
		score += char_scores[b]
	}
	return score
}
func ScoreList(strs []string) string {
	best := 0
	var best_str string
	for _, s := range strs {
		score := scoreXor([]byte(s))
		if score > best {
			best = score
			best_str = s
		}

	}
	return best_str
}
func SingleByteXor(input string) (key byte, output string) {

	best_score := 0
	best_xor := new(string)
	var best_key byte

	hex, _ := hex.DecodeString(input)
	for c := 0; c < 255; c++ {
		xor := xor(hex, byte(c))
		score := scoreXor(xor)
		if score > best_score {
			best_score = score
			*best_xor = string(xor)
			best_key = byte(c)
		}
	}
	return best_key, *best_xor
}
