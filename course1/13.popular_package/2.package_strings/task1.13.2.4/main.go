package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func generateActivationKey() string {
	key := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		keyPart := make([]byte, 0, 4)
		for j := 0; j < 4; j++ {
			bigX, _ := rand.Int(rand.Reader, big.NewInt(26))
			intX := bigX.Int64() + 48
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
