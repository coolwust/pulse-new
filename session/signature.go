package session

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"
	"strings"
)

var keys = [][]byte{[]byte("hello world"), []byte("foo bar")}

var ErrInvalidSignature = errors.New("session: cookie signature is invalid")

func Sign(unsigned string) string {
	return sign(unsigned, keys[len(keys)-1])
}

func sign(unsigned string, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(unsigned))
	h := mac.Sum(nil)
	b := make([]byte, base64.StdEncoding.EncodedLen(len(h)))
	base64.StdEncoding.Encode(b, h)
	return unsigned + "." + string(bytes.Join(regexp.MustCompile(`\w+`).FindAll(b, -1), nil))
}

func Unsign(signed string) (string, error) {
	return unsign(signed, keys)
}

func unsign(signed string, keys [][]byte) (string, error) {
	s := strings.Split(signed, ".")
	if len(s) != 2 {
		return "", ErrInvalidSignature
	}
	for i := len(keys) - 1; i >= 0; i-- {
		if sign(s[0], keys[i]) == signed {
			return s[0], nil
		}
	}
	return "", ErrInvalidSignature
}
