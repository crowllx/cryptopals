package set1

import (
	"crypto/aes"
	"encoding/base64"
)

func Decrypt(key []byte, data string) string {
	cipher, _ := aes.NewCipher(key)
	ct, _ := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(data)
	pt := make([]byte, len(ct))
	size := len(key)

	for bs, be := 0, size; bs < len(ct); bs, be = bs+size, be+size {
		cipher.Decrypt(pt[bs:be], ct[bs:be])
	}
	return string(pt)
}
