package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "0"
	s2 := "1"
	s3 := "1"
	s4 := "2"
	s5 := "3"
	s6 := "5"
	s7 := "8"
	s8 := "13"

	var builder strings.Builder

	builder.WriteString(s1)
	builder.WriteString(s2)
	builder.WriteString(s3)
	builder.WriteString(s4)
	builder.WriteString(s5)
	builder.WriteString(s6)
	builder.WriteString(s7)
	builder.WriteString(s8)

	fmt.Println(builder.String())
}
