package main

import (
	"cryptopals/set1"
	"fmt"
	"os"
	"strings"
)

func main() {
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
