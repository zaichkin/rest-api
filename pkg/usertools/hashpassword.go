package usertools

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashPassword(password string) string {
	ps := []byte(password)
	h := sha1.New()
	h.Write(ps)
	hashpasswd := hex.EncodeToString(h.Sum(nil))
	return hashpasswd
}
