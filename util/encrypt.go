package util

import (
	"crypto/rc4"
	"encoding/base64"
)

const tokenRc4Key = "WHdJKoPaBj4FEupnJYZyAU8pUEB6qvbjW5w0yd"

// Rc4 rc4 encryption
func Rc4(data []byte) ([]byte, error) {
	c, err := rc4.NewCipher([]byte(tokenRc4Key))
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	c.XORKeyStream(dst, data)
	return dst, nil
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
