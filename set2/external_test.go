package set2_test

import (
	"bytes"
	"cryptopals/set2"
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
