package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_AISelectWithout struct {
	Hand []int64
}

type Output_AISelectWithout struct {
	AISelectWithout int
	Index           int64
}

func AISelectWithout(data Input_AISelectWithout) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("AISelectWithout_Interface")
	Hand_Length := int64(len(data.Hand))
	ret, _, _ := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(Hand_Length),
	)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
