package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

func main() {
	testEmails := []string{
		"test@example.com",
		"invalid-email",
		"another.test@domain.co",
		"wrong@domain,com",
		"user@domain",
	}

	for _, email := range testEmails {
		if isValidEmail(email) {
			fmt.Printf("%s is a valid email address.\n", email)
		} else {
			fmt.Printf("%s is not a valid email address.\n", email)
		}
	}
}
