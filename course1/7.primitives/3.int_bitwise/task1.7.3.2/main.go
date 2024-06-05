package main

import "fmt"

var Permissions = map[int]string{
	0: "-,-,-",
	1: "-,-,Execute",
	2: "-,Write,-",
	3: "-,Write,Execute",
	4: "Read,-,-",
	5: "Read,-,Execute",
	6: "Read,Write,-",
	7: "Read,Write,Execute",
}

func main() {
	fmt.Println(getFilePermissions(755))
}

func getFilePermissions(flag int) string {
	return fmt.Sprintf("Owner:%s Group:%s Other:%s",
		Permissions[flag/100],
		Permissions[flag%100/10],
		Permissions[flag%10])
}
