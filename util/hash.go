package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	key := hex.EncodeToString(h.Sum(nil))
	return key
}
