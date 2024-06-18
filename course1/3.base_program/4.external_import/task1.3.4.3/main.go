package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func main() {
	fmt.Println(GenerateFakeData())
}

func GenerateFakeData() string {
	return fmt.Sprintf(
		"Name: %s\n"+
			"Address: %s\n"+
			"Phone: %s\n"+
			"Email: %s\n",
		fake.FullName(),
		fake.StreetAddress(),
		fake.Phone(),
		fake.EmailAddress())

}
