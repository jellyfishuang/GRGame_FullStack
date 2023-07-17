package Api_8joker

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%sn", s);
// }

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Output_test struct {
	Index int32
	Arr   [5]int32
	Bok   bool
}

func TestSOFile() int {
	println("hello cgo")

	return 0
}

func TestDLL() (out int) {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	test := lib.NewProc("TestGR")

	Hand := []int32{28, 2, 3, 4, 5, -1}
	Tile := int32(14)
	flo := float32(2.897)

	input := Output_test{
		Index: 178,
		Arr:   [5]int32{1, 2, 3, 4, 5},
		Bok:   true,
	}

	ret, _, _ := test.Call(
		uintptr(unsafe.Pointer(&Hand[0])),
		uintptr(Tile),
		uintptr(unsafe.Pointer(&input)),
		uintptr(unsafe.Pointer(&flo)),
	)

	result := int32(ret)
	fmt.Println("result", result)
	out = int(result)
	return out
}
