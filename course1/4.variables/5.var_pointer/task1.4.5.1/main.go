package main

import "fmt"

func main() {

	var i int = 1
	var f float64 = 1.35
	var s string = "a"
	var b bool = false

	fmt.Println(i, " ", f, " ", s, " ", b)

	changeInt(&i)
	changeFloat(&f)
	changeString(&s)
	changeBool(&b)

	fmt.Println(i, " ", f, " ", s, " ", b)

}

func changeInt(a *int) {
	*a = 2
}

func changeFloat(a *float64) {
	*a = 2.35
}

func changeString(a *string) {
	*a = "b"
}

func changeBool(a *bool) {
	*a = true
}
