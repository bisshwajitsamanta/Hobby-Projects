package encrypt

func EncryptData(data string) string {
	encryptedData := ""

	for _, char := range data {
		asciiCode := int(char)
		character := string(rune(asciiCode + 3))
		encryptedData += character
	}

	return encryptedData
}
