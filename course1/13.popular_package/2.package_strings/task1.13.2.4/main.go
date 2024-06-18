package main

import (
	"crypto/rand"
	"math/big"

	"fmt"

	"strings"
)

const (
	licenceKeyParts       = 4
	licenceKeyPartSymbols = 4
)

func generateActivationKey() string {
	key := make([]string, 0, licenceKeyParts)
	for i := 0; i < licenceKeyParts; i++ {
		keyPart := make([]byte, 0, licenceKeyPartSymbols)
		for j := 0; j < licenceKeyPartSymbols; j++ {
			bigX, _ := rand.Int(rand.Reader, big.NewInt(26)) // one of the 26 letters of the english alphabet

			// adding 48 to match one of ASCII codes for english digits and letters
			intX := bigX.Int64() + 48

			// if the number is between ASCII codes for digits and letters,
			// subtracting 7 to make it match one of the ASCII codes for digits
			if intX >= 58 && intX <= 64 {
				intX -= 7
			}
			keyPart = append(keyPart, byte(intX))
		}
		key = append(key, string(keyPart))
	}
	return strings.Join(key, "-")
}

func main() {
	fmt.Println(generateActivationKey())
}
