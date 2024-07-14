package main

import (
	"cryptopals/set1"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func singleCharXor() {
	key, xor := set1.SingleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println("char: ", key, "\nxor:", xor)

	dat, err := os.ReadFile("./tests/single-char-xor.txt")
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(dat), "\n")

	var xors []string
	for _, s := range inputs {
		_, xor = set1.SingleByteXor(s)
		xors = append(xors, xor)
	}
	best := set1.ScoreList(xors)
	fmt.Println(best)
}

func aesDecrypt() {
	key := "YELLOW SUBMARINE"
	ct, _ := os.ReadFile("./tests/aes-cipher.txt")

	pt := set1.Decrypt([]byte(key), string(ct))
	fmt.Println(pt)
}
func breakVignere() []byte {
	dat, err := os.ReadFile("./tests/vignere.txt")
	if err != nil {
		panic(err)
	}
	decoded, _ := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(string(dat))
	fmt.Println(len(decoded))
	key := set1.BreakVigenere(string(decoded))
	fmt.Println(string(key))
	// fmt.Println(set1.RepeatingXor(decoded, key))
	return key
}
func main() {
	singleCharXor()
	fmt.Println()
	breakVignere()
	aesDecrypt()

}
