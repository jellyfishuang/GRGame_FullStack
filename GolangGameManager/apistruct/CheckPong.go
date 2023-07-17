package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_CheckPong struct {
	Hand      []int64
	Tile      int64
	Without   int64
	AlreadyHu bool
}

type Output_CheckPong struct {
	CanPong    int
	Error_code int
}

func CheckPong(data Input_CheckPong) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("CheckPong_Interface")
	Hand_Length := int64(len(data.Hand))
	ret, _, _ := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(Hand_Length),
		TransformGolang.Int64Ptr(data.Tile),
		TransformGolang.Int64Ptr(data.Without),
		TransformGolang.BoolPtr(data.AlreadyHu),
	)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
