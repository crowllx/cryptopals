package set1

import (
	"encoding/hex"
	"fmt"
)

func diff(l []byte, r []byte) ([]byte, []byte, int) {
	diff := len(l) - len(r)
	if diff < 0 {
		diff = -diff
		return l, r, diff
	}
	return r, l, diff
}
func binary(s []byte) []byte {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return []byte(res)
}
func editDistance(l []byte, r []byte) int {
	l, r, diff := diff(binary(l), binary(r))
	for i, c := range l {
		if c != r[i] {
			diff++
		}
	}
	return diff
}

type keyData struct {
	key  int
	dist float32
}

func avg(arr [][]byte) float32 {
	score := 0
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			score += editDistance(arr[i], arr[j])
		}
	}
	return float32(score) / float32(len(arr))
}
func findKeysize(ct []byte) int {
	min := float32(len(ct))
	keyLen := len(ct)

	for key := 2; key <= 40; key++ {
		chunks := chunkify(ct, key)
		sample := chunks[:4]
		avg := avg(sample) / float32(key)
		if avg < min {
			min = avg
			keyLen = key
		}
	}
	return keyLen
}

func chunkify(ct []byte, key int) [][]byte {
	var chunks [][]byte
	prev := 0
	for i := key; i <= len(ct); i += key {
		chunks = append(chunks, ct[prev:i])
		prev = i
	}
	if len(ct)%key != 0 {
		paddLength := key - len(ct[prev:])
		tmp := ct[prev:]
		for i := 0; i < paddLength; i++ {
			tmp = append(tmp, byte(paddLength))
		}
		chunks = append(chunks, tmp)
	}
	return chunks
}
func transpose(chunks [][]byte) [][]byte {
	size := len(chunks[0])
	var result = make([][]byte, size)
	for _, c := range chunks {
		for i, b := range c {
			result[i] = append(result[i], b)
		}
	}
	return result
}
func BreakVigenere(ciphertext string) []byte {
	ct := []byte(ciphertext)
	keySize := findKeysize(ct)
	chunks := chunkify(ct, keySize)
	transposed := transpose(chunks)
	var key []byte
	for _, x := range transposed {
		c, _ := SingleByteXor(hex.EncodeToString(x))
		key = append(key, c)
	}

	return key
}
