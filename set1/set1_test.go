package set1_test

import (
	"cryptopals/set1"
	"testing"
)

func TestHexConversion(t *testing.T) {
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	result, err := set1.Hex2base64(input)
	if expected != result || err != nil {
		t.Fatalf(`
        input: %s,
        output: %s
        `, input, result)

	}
	t.Log("expected: ", expected)
	t.Log("result: ", result)
}

func TestFixedXor(t *testing.T) {
	expected := "746865206b696420646f6e277420706c6179"
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	result, err := set1.FixedXor(input1, input2)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if result != expected {
		t.Fatalf(`
            Result does not match
            result: %s,
            expected: %s
            `, result, expected)
	}

	t.Log("expected: ", expected)
	t.Log("result: ", result)
}
