package main

import (
	"encrypt/Decrypt"
	"encrypt/encrypt"
	"fmt"
)

func main() {
	originalText := "Hello, World!"
	encryptedText := encrypt.EncryptData(originalText)
	fmt.Printf("Original Text: %s\n", originalText)
	fmt.Printf("Encrypted Text: %s\n", encryptedText)
	decryptedText := Decrypt.Nimbus(encryptedText)
	fmt.Printf("Decrypted Text: %s\n", decryptedText)
}
