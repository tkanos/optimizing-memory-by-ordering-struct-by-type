package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myInt   bool    // 1 bit
	myFloat float64 // 8 bytes
	myBool  int32   // 4 bytes
}

type myStructOptimized struct {
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
	myBool  bool    // 1 bytes
}

func main() {
	a := myStruct{}
	b := myStructOptimized{}

	fmt.Println(unsafe.Sizeof(a)) // unordered 24 bytes
	fmt.Println(unsafe.Sizeof(b)) // ordered 16 bytes
}
