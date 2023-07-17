package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_VerifyPong struct {
	Hand      []int64
	PongTile  int64
	Without   int64
	AlreadyHu bool
}

type Output_VerifyPong struct {
	VerifyPong int
	Error_code int
}

func VerifyPong(data Input_VerifyPong) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("VerifyPong_Interface")
	Hand_Length := int64(len(data.Hand))
	ret, _, _ := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(Hand_Length),
		TransformGolang.Int64Ptr(data.PongTile),
		TransformGolang.Int64Ptr(data.Without),
		TransformGolang.BoolPtr(data.AlreadyHu),
	)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
