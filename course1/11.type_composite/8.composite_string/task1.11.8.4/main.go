package main

import "strings"

func concatStrings(xs ...string) string {
	str := strings.Builder{}
	for _, x := range xs {
		str.WriteString(x)
	}
	return str.String()
}
