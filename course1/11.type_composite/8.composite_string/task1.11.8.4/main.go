package main

func concatStrings(xs ...string) string {
	str := ""
	for _, x := range xs {
		str += x
	}
	return str
}
