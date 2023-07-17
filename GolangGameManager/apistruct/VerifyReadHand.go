package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"fmt"
	"syscall"
	"unsafe"
)

type Input_VerifyReadHand struct {
	Hand          []int64
	Pong          []int64
	Kong          []int64
	ConcealedKong []int64
	Without       int64
}

type Output_VerifyReadHand struct {
	VerifyReadHand int
	Error_code     int
}

func VerifyReadHand(data Input_VerifyReadHand) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("VerifyReadHand_Interface")
	Hand_Length := int64(len(data.Hand))
	Hand_Pong_Length := int64(len(data.Pong))
	Hand_Kong_Length := int64(len(data.Kong))
	Hand_ConcealedKong_Length := int64(len(data.ConcealedKong))
	ret, _, err := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(Hand_Length),
		TransformGolang.IntArrayPtr(data.Pong),
		TransformGolang.Int64Ptr(Hand_Pong_Length),
		TransformGolang.IntArrayPtr(data.Kong),
		TransformGolang.Int64Ptr(Hand_Kong_Length),
		TransformGolang.IntArrayPtr(data.ConcealedKong),
		TransformGolang.Int64Ptr(Hand_ConcealedKong_Length),
		TransformGolang.Int64Ptr(data.Without),
	)
	fmt.Println(err)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
