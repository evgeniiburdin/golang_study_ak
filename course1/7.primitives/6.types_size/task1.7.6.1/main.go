package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		b   bool
		n   int
		n8  int8
		n16 int16
		n32 int32
		n64 int64
		un  uint
		un8 uint8
	)
	fmt.Printf("bool size: %d\n", sizeOfBool(b))
	fmt.Printf("int size: %d\n", sizeOfInt(n))
	fmt.Printf("in8 size: %d\n", sizeOfInt8(n8))
	fmt.Printf("int16 size: %d\n", sizeOfInt16(n16))
	fmt.Printf("int32 size: %d\n", sizeOfInt32(n32))
	fmt.Printf("int64 size: %d\n", sizeOfInt64(n64))
	fmt.Printf("uint size: %d\n", sizeOfUint(un))
	fmt.Printf("uint8 size: %d\n", sizeOfUint8(un8))
}

func sizeOfBool(b bool) int   { return int(unsafe.Sizeof(b)) }
func sizeOfInt(n int) int     { return int(unsafe.Sizeof(n)) }
func sizeOfInt8(n int8) int   { return int(unsafe.Sizeof(n)) }
func sizeOfInt16(n int16) int { return int(unsafe.Sizeof(n)) }
func sizeOfInt32(n int32) int { return int(unsafe.Sizeof(n)) }
func sizeOfInt64(n int64) int { return int(unsafe.Sizeof(n)) }
func sizeOfUint(n uint) int   { return int(unsafe.Sizeof(n)) }
func sizeOfUint8(n uint8) int { return int(unsafe.Sizeof(n)) }
