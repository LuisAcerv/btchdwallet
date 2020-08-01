package crypt

import (
	"crypto/rand"
	"encoding/hex"
)

// CreateHash returns random 32 byte array and it's hex string to be stored
func CreateHash() string {
	key := make([]byte, 8)

	_, err := rand.Read(key)
	if err != nil {
		// handle error here
		panic(err)
	}

	str := hex.EncodeToString(key)

	return str
}
