package main

import "fmt"

var (
	name,
	surname,
	age,
	city string
)

var (
	ageInt,
	year,
	month,
	day int
)

func main() {

	name = "Christian"
	surname = "Smith"
	age = "27"
	city = "Ohio"

	ageInt = 27
	year = 1997
	month = 03

	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(age)
	fmt.Println(city)

	fmt.Println(ageInt)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

}
