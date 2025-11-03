package Decrypt

func Nimbus(str string) string {
	decryptedData := ""

	for _, char := range str {
		asciiCode := int(char)
		character := string(rune(asciiCode - 3))
		decryptedData += character
	}
	return decryptedData
}
