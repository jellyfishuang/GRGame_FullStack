package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_VerifyHu struct {
	Hand          []int64
	Tile          int64
	Pong          []int64
	Kong          []int64
	ConcealedKong []int64
	Without       int64
	IsZimo        bool
	GameRule      TransformGolang.CppGameRule
}

type Output_VerifyHu struct {
	VerifyHu   int
	Error_code int
}

func VerifyHu(data Input_VerifyHu) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("VerifyHu_Interface")
	Hand_Length := int64(len(data.Hand))
	Hand_Pong_Length := int64(len(data.Pong))
	Hand_Kong_Length := int64(len(data.Kong))
	Hand_ConcealedKong_Length := int64(len(data.ConcealedKong))
	ret, _, _ := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(Hand_Length),
		TransformGolang.Int64Ptr(data.Tile),
		TransformGolang.IntArrayPtr(data.Pong),
		TransformGolang.Int64Ptr(Hand_Pong_Length),
		TransformGolang.IntArrayPtr(data.Kong),
		TransformGolang.Int64Ptr(Hand_Kong_Length),
		TransformGolang.IntArrayPtr(data.ConcealedKong),
		TransformGolang.Int64Ptr(Hand_ConcealedKong_Length),
		TransformGolang.Int64Ptr(data.Without),
		TransformGolang.BoolPtr(data.IsZimo),

		TransformGolang.BoolPtr(data.GameRule.ChangeTileSameColor),
		TransformGolang.Int64Ptr(data.GameRule.ChangeTileCount),
		TransformGolang.BoolPtr(data.GameRule.CanEat),
		TransformGolang.BoolPtr(data.GameRule.CanPong),
		TransformGolang.BoolPtr(data.GameRule.CanKong),
		TransformGolang.IntArrayPtr(data.GameRule.XorTable),
		TransformGolang.IntArrayPtr(data.GameRule.FanCount),
		TransformGolang.IntArrayPtr(data.GameRule.JokerTurnIntos),
		TransformGolang.Int64Ptr(data.GameRule.LimitTai),
	)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
