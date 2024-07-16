package main

func RepeatingXor(plaintext []byte, key []byte) []byte {
	key_len := len(key)
	key_index := 0
	var xor []byte
	for _, b := range plaintext {
		xor = append(xor, b^key[key_index])
		key_index = (key_index + 1) % key_len
	}
	return xor
}
