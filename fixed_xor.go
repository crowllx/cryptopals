package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func FixedXor(input1 string, input2 string) (string, error) {

	if (len(input1) != len(input2)) && len(input1)/2 != 0 {
		return "", errors.New("invalid input")
	}

	input1_bytes, err := hex.DecodeString(input1)
	if err != nil {
		return "", err
	}

	input2_bytes, err := hex.DecodeString(input2)
	if err != nil {
		return "", err
	}

	fmt.Println(input1_bytes)
	fmt.Println(input2_bytes)
    result := make([]byte, len(input1_bytes))
    for i, b := range input1_bytes {
        result[i] = b ^ input2_bytes[i]
    }

	return hex.EncodeToString(result), nil
}
