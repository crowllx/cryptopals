package set2

import (
	"cryptopals/set1"
	"fmt"
	"math/rand"
)

func generateRandBytes(size int) []byte {
	key := make([]byte, size)
	for i := range size {
		key[i] = byte(rand.Intn(256))
	}
	return key
}

func Oracle(input string) (string, error) {
	pt := []byte(input)

	// append 5-10 bytes to beginning and end of input
	prefixLen := rand.Intn(5) + 1
	suffixLen := rand.Intn(5) + 1
	prefix := generateRandBytes(prefixLen)
	suffix := generateRandBytes(suffixLen)
	pt = append(prefix, append(pt, suffix...)...)

	// generate random encryption key
	key := generateRandBytes(16)
	// half the time use ECB, the other half use CBC encryption
	mode := rand.Intn(2)
	fmt.Println(mode)
	var ct []byte
	fmt.Println(mode)
	if mode == 1 {
		fmt.Println("CBC")
		res, err := Encrypt(pt, key, generateRandBytes(16))

		if err != nil {
			panic(err)
		}
		ct = res
	} else {
		fmt.Println("ECB")
		res, err := set1.ECBEncrypt(key, pt)
		if err != nil {
			panic(err)
		}
		ct = res
	}
	return string(ct), nil

}
