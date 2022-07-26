package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncryptPassword(str string) string {
	h := sha1.New()
	_, err := h.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
