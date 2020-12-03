package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {

	fmt.Println("flux HMAC-SHA256 encrypt ")

	//secretKey = "ABCDEFGHIJMNBHDS"
	secretKey := []byte("ABCDEFGHIJMNBHDS")

	message := "HyperLedger Fabric : HMAC-SHA256 Encrypt 개발 ..."
	salt := generateSalt()
	fmt.Println("Message: " + message)
	fmt.Println("\nSalt: " + salt)

	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\nHMAC-Sha256: %x", hash.Sum(nil))

	hash = hmac.New(sha512.New, []byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\n\nHMAC-sha512: %x", hash.Sum(nil))

	//rsaPriKey, rsaPubKey := GenerateKeyPair(1024)

}
