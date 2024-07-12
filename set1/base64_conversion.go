package set1

import (
	"encoding/hex"
)

func make_b64_map() []rune {
	var char_map []rune
	for c := 'A'; c <= 'Z'; c++ {
		char_map = append(char_map, c)
	}
	for c := 'a'; c <= 'z'; c++ {
		char_map = append(char_map, c)
	}
	for x := range 10 {
		char_map = append(char_map, rune(x+48))
	}
	char_map = append(char_map, '+')
	char_map = append(char_map, '\\')
	return char_map
}
func Hex2base64(input string) (string, error) {
	// var base64str string
    var result string
	bytes, _ := hex.DecodeString(input)
	num_bytes := len(bytes)
	char_map := make_b64_map()
	for i := 0; i < num_bytes; i += 3 {
        var chunk_result string = ""
		chunk := int(bytes[i])
		chunk <<= 8
		chunk += int(bytes[i+1])
		chunk <<= 8
		chunk += int(bytes[i+2])
        
		for i := 0; i < 4; i++ {
			c := chunk & 0b111111
			chunk >>= 6

            b64char := char_map[c]
            chunk_result = string(b64char) + chunk_result
		}
        result += chunk_result
	}
	return result, nil
}
