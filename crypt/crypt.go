package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
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

// GenerateHashFromPin to encrypt and decrypt secret
func GenerateHashFromPin(pin string) string {
	s := fmt.Sprintf("%s", pin)

	sha256 := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", sha256)
}

// HashToString to string
func HashToString(hash string) []byte {
	return []byte(hash)
}

// Encrypt string to base64 crypto using AES
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// Decrypt from base64 to decrypted string
func Decrypt(encryptedString string, keyString string) (decryptedString string, _err bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return "", false
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
		return "", false
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	return fmt.Sprintf("%s", plaintext), true
}
