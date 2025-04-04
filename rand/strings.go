package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 32

func Read(b []byte) (n int, err error) {
	return rand.Read(b)
}

// Bytes will help us generate n random bytes, or will
// return an error if there was one. This uses the
// crypto/rand package so it is safe to use with things
// like remember tokens.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Strings will generate a byte slice of size nBytes and then
// return a string that is the base64 URL encoded version
// of that byte slice
func Strings(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func RememberToken() (string, error) {
	return Strings(RememberTokenBytes)
}
