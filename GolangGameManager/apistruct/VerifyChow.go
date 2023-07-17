package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_VerifyChow struct {
	Hand      []int64
	ChowTile  int64
	Sequence  []int64
	Without   int64
	AlreadyHu bool
}

type Output_VerifyChow struct {
	VerifyChow int
	Error_code int
}

func VerifyChow(data Input_VerifyChow) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("VerifyChow_Interface")
	Hand_Length := len(data.Hand)
	ret, _, _ := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(int64(Hand_Length)),
		TransformGolang.Int64Ptr(data.ChowTile),
		TransformGolang.IntArrayPtr(data.Sequence),
		TransformGolang.Int64Ptr(data.Without),
		TransformGolang.BoolPtr(data.AlreadyHu),
	)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
