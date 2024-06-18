package main

import (
	"crypto/rand"
	"math/big"

	"fmt"

	"strings"
)

func GenerateRandomString(length int) string {
	var str strings.Builder
	for i := 0; i < length; i++ {
		randomNumber, _ := rand.Int(rand.Reader, big.NewInt(26))
		str.WriteByte(byte(randomNumber.Int64() + 65))
	}
	return str.String()
}

func main() {
	randomString := GenerateRandomString(10)
	fmt.Println(randomString)
}
