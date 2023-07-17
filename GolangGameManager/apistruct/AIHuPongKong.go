package ApiStruct

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_AIHuPongKong struct {
	Hand          []int64
	Tile          int64
	UnOpenPool    []int64
	Without       int64
	AlreadyHu     bool
	Pong          []int64
	Kong          []int64
	ConcealedKong []int64
	GameRule      TransformGolang.CppGameRule
	Strategies    TransformGolang.CppStrategy
}

type Output_AIHuPongKong struct {
	AIHuPongKong int
	Index        int64
}

func AIHuPongKong(data Input_AIHuPongKong) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	add := lib.NewProc("AIHuPongKong_Interface")
	Hand_Length := int64(len(data.Hand))
	UnOpenPool_Length := int64(len(data.UnOpenPool))
	Hand_Pong_Length := int64(len(data.Pong))
	Hand_Kong_Length := int64(len(data.Kong))
	Hand_ConcealedKong_Length := int64(len(data.ConcealedKong))
	ret, _, _ := add.Call(
		TransformGolang.IntArrayPtr(data.Hand),
		TransformGolang.Int64Ptr(Hand_Length),
		TransformGolang.Int64Ptr(data.Tile),
		TransformGolang.IntArrayPtr(data.UnOpenPool),
		TransformGolang.Int64Ptr(UnOpenPool_Length),
		TransformGolang.Int64Ptr(data.Without),
		TransformGolang.IntArrayPtr(data.Pong),
		TransformGolang.Int64Ptr(Hand_Pong_Length),
		TransformGolang.IntArrayPtr(data.Kong),
		TransformGolang.Int64Ptr(Hand_Kong_Length),
		TransformGolang.IntArrayPtr(data.ConcealedKong),
		TransformGolang.Int64Ptr(Hand_ConcealedKong_Length),
		TransformGolang.BoolPtr(data.AlreadyHu),

		TransformGolang.BoolPtr(data.GameRule.ChangeTileSameColor),
		TransformGolang.Int64Ptr(data.GameRule.ChangeTileCount),
		TransformGolang.BoolPtr(data.GameRule.CanEat),
		TransformGolang.BoolPtr(data.GameRule.CanPong),
		TransformGolang.BoolPtr(data.GameRule.CanKong),
		TransformGolang.IntArrayPtr(data.GameRule.XorTable),
		TransformGolang.IntArrayPtr(data.GameRule.FanCount),
		TransformGolang.IntArrayPtr(data.GameRule.JokerTurnIntos),
		TransformGolang.Int64Ptr(data.GameRule.LimitTai),

		TransformGolang.Int64Ptr(data.Strategies.HuDistance),
		TransformGolang.Float64Ptr(data.Strategies.FaanTaiWeight),
		TransformGolang.Float64Ptr(data.Strategies.DistanceWeight),
		TransformGolang.Float64Ptr(data.Strategies.PossibleInTilesWeight),
		TransformGolang.Int64Ptr(data.Strategies.DiscardCount),
		TransformGolang.Int64Ptr(data.Strategies.Rule),
	)
	var returnValue *int = (*int)(unsafe.Pointer(&ret))
	return *returnValue
}
