package main

import (
	"cryptopals/set1"
	"fmt"
)

func main() {
	result := set1.SingleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    for _, s := range result {
        fmt.Println(s)
    }
}
