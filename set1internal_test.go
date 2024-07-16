package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEditDistance(t *testing.T) {
	dist := editDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	expected := 37
	if dist != expected {
		t.Fatalf(`
			dist: %d,
			expected: %d
			`, dist, expected)
	}

	t.Logf(`
		dist: %d
		expected: %d
		`, dist, expected)
}
func TestChunkify(t *testing.T) {
	start := "aaaabbbbccccdddde"
	expected := make([][]byte, 4)
	for c := 'a'; c <= 'd'; c++ {
		chunk := bytes.Repeat([]byte{byte(c)}, 4)
		expected = append(expected, chunk)
	}
	result := chunkify([]byte(start), 4)

	t.Logf("expected: %v\nresult: %v", expected, result)

}

func TestTranspose(t *testing.T) {
	start := [][]byte{
		[]byte("aaaaaa"),
		[]byte("bbbbbb"),
		[]byte("cccccc"),
	}

	result := transpose(start)
	t.Logf("transpose: %v", result)
}
func TestFindKeySize(t *testing.T) {
	pt1 := `somepass very unique and lengthy. i hope I'm making progress. I think
	this needs to be longer so here we go adding some more text`
	key1 := "Rich in fiber"
	key2 := "ICE"
	ct1 := RepeatingXor([]byte(pt1), []byte(key1))
	ct2 := RepeatingXor([]byte(pt1), []byte(key2))
	fmt.Printf("%s\n%s\n", ct1, ct2)
	keySize1 := findKeysize([]byte(ct1))
	keySize2 := findKeysize([]byte(ct2))

	if keySize1 != len(key1) || keySize2 != len(key2) {
		t.Fatalf(`
Incorrect key size: expected result
			key1:	%d		 %d
			key2:	%d		 %d`, len(key1), keySize1, len(key2), keySize2)
	}
	t.Logf(`
Success 		expected result
		key1:	%d		 %d
		key2:	%d		 %d`, len(key1), keySize1, len(key2), keySize2)

}
