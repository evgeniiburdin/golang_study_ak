package main

import (
	"reflect"
	"unsafe"
)

func getStringHeader(s string) reflect.StringHeader {
	return *(*reflect.StringHeader)(unsafe.Pointer(&s))
}
