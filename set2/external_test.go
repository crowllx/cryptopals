package set2_test

import (
	"bytes"
	"cryptopals/set2"
	"fmt"
	"testing"
)

func TestPksc7(t *testing.T) {
	blockSize := 20
	pt := "YELLOW SUBMARINE"
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	result := set2.Pksc7Pad([]byte(pt), blockSize)

	if !bytes.Equal(expected, result) {
		t.Fatalf(`
			failed
			result: %v
			expected: %v`, []byte(result), []byte(expected))

	}
}

func TestCBC(t *testing.T) {
	// test will not work if input is not divisisble by 16
	// currently decryption does not take into account pksc7 padding
	pt := []byte("some test text!!\ndefinately work another 16 byte")
	key := []byte("YELLOW SUBMARINE")
	iv := bytes.Repeat([]byte("\x00"), 16)
	ct, _ := set2.Encrypt(pt, key, iv)
	decrypted, _ := set2.Decrypt(ct, key, iv)
	fmt.Printf("pt: %s\n", pt)
	fmt.Printf("dec: %s\n", decrypted)
	if !bytes.Equal(pt, decrypted) {
		t.Fatalf(`
			encryption or decryption failed.
			pt: %s,
			ct: %s,
			decrypted: %s`, pt, ct, decrypted)
	}

}
