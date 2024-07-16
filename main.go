package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func singleCharXor() {

	key, xor := SingleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println("char: ", key, "\nxor:", xor)

	dat, err := os.ReadFile("./tests/single-char-xor.txt")
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(dat), "\n")

	var xors []string
	for _, s := range inputs {
		_, xor = SingleByteXor(s)
		xors = append(xors, xor)
	}
	best := ScoreList(xors)
	fmt.Println(best)
}
func detectAesECB() {
	data, _ := os.ReadFile("./tests/aes-cipher2.txt")
	cts := strings.Split(string(data), "\n")

	for i, ct := range cts {
		decoded, _ := hex.DecodeString(string(ct))
		res := DetectECB(string(decoded))
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

	pt, _ := ECBDecrypt([]byte(key), ct)
	fmt.Println(pt)
}
func breakVignere() []byte {
	dat, err := os.ReadFile("./tests/vignere.txt")
	if err != nil {
		panic(err)
	}
	decoded, _ := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(string(dat))
	fmt.Println(len(decoded))
	key := BreakVigenere(string(decoded))
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
	decrypted, _ := Decrypt(decoded, key, iv)
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
	ct, err := Oracle(pt)
	if err != nil {
		panic(err)
	}
	ecb := DetectECB(ct)
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
