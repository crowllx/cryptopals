package main

import (
	"bytes"
	"cryptopals/set1"
	"cryptopals/set2"
	"encoding/base64"
	"encoding/hex"
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
func detectAesECB() {
	data, _ := os.ReadFile("./tests/aes-cipher2.txt")
	cts := strings.Split(string(data), "\n")

	for i, ct := range cts {
		decoded, _ := hex.DecodeString(string(ct))
		res := set1.DetectECB(string(decoded))
		if res {
			fmt.Printf("%v\nct: %s\nindex: %d\n", res, ct, i)
			return
		}
	}
	fmt.Println("err")
}
func aesDecrypt() {
	key := "YELLOW SUBMARINE"
	ct, _ := os.ReadFile("./tests/aes-cipher.txt")

	pt, _ := set1.ECBDecrypt([]byte(key), ct)
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

func set2tests() {
	key := []byte("YELLOW SUBMARINE")
	ct, _ := os.ReadFile("./tests/cbc-encrypted.txt")
	fmt.Println(len(ct))
	decoded, _ := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(string(ct))
	fmt.Println(len(decoded))
	iv := bytes.Repeat([]byte("\x00"), 16)
	decrypted, _ := set2.Decrypt(decoded, key, iv)
	fmt.Printf("%s", string(decrypted))
}

func identifyOracleMode() {
	pt := "some test data that is amazingthis is another block for whatever"
	fmt.Println(pt[:16])
	fmt.Println(pt[32:48])
	fmt.Println(pt[32:])
	fmt.Println(len(pt[32:]))
	pt += pt
	fmt.Println(len(pt))
	ct, err := set2.Oracle(pt)
	if err != nil {
		panic(err)
	}
	ecb := set1.DetectECB(ct)
	fmt.Printf("ecb: %v\n", ecb)
	if ecb {
		fmt.Println("mode used was ECB")
	} else {
		fmt.Println("mode used was probably CBC")
	}

}
func main() {
	singleCharXor()
	fmt.Println()
	breakVignere()
	// aesDecrypt()
	detectAesECB()

	fmt.Println()
	// set2tests()
	identifyOracleMode()
}
