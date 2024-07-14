package set1_test

import (
	"cryptopals/set1"
	"encoding/hex"
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

func TestRepeatingXor(t *testing.T) {
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	xor := set1.RepeatingXor([]byte(input), []byte(key))

	if hex.EncodeToString([]byte(xor)) != expected {
		t.Fatalf(`
			Result does not match
			result: %s,
			expected: %s
			`, xor, expected)
	}
	t.Logf(`
		result: %s,
		expected: %s
		`, hex.EncodeToString([]byte(xor)), expected)

}
