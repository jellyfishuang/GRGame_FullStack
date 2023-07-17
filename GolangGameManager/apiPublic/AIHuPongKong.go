package Api_Public

import (
	TransformGolang "GolangGameManager/transform"
	"fmt"
	"syscall"
	"unsafe"
)

type Input_AIHuPongKong struct {
	Hand                  []int32
	Tile                  int32
	Chow                  []int32
	Pong                  []int32
	Kong                  []int32
	ConcealedKong         []int32
	UnOpenPool            []int32
	Flower                []int32
	IsZimo                int32
	Uid                   int32
	Without               int32
	SingleJokerAnGonCount int32
	DoorSize              []int32
	ThrowSeq              []int32
	CheatTiles            []int32
	AlreadyHuTiles        []int32
	AlreadyHu             bool
	AlreadyMing           bool
	AlreadyCrit           bool
	IsTingCard            bool
	Strategies            TransformGolang.Strategies
}

func AIHuPongKong(data Input_AIHuPongKong) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	AIHuPongKong := lib.NewProc("APIAIHuPongKong")

	ret, _, _ := AIHuPongKong.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(data.Tile),
		uintptr(unsafe.Pointer(&data.Chow[0])),
		uintptr(unsafe.Pointer(&data.Pong[0])),
		uintptr(unsafe.Pointer(&data.Kong[0])),
		uintptr(unsafe.Pointer(&data.ConcealedKong[0])),
		uintptr(unsafe.Pointer(&data.UnOpenPool[0])),
		uintptr(unsafe.Pointer(&data.Flower[0])),
		uintptr(data.IsZimo),
		uintptr(data.Uid),
		uintptr(data.Without),
		uintptr(data.SingleJokerAnGonCount),
		uintptr(unsafe.Pointer(&data.DoorSize[0])),
		uintptr(unsafe.Pointer(&data.ThrowSeq[0])),
		uintptr(unsafe.Pointer(&data.CheatTiles[0])),
		uintptr(unsafe.Pointer(&data.AlreadyHuTiles[0])),
		uintptr(unsafe.Pointer(&data.AlreadyHu)),
		uintptr(unsafe.Pointer(&data.AlreadyMing)),
		uintptr(unsafe.Pointer(&data.AlreadyCrit)),
		uintptr(unsafe.Pointer(&data.IsTingCard)),

		// 傳入 Strategies
		uintptr(unsafe.Pointer(&data.Strategies)),
		// uintptr(data.Strategies.HuDistance),
		// uintptr(data.Strategies.DistanceWeight),
		// uintptr(data.Strategies.FaanTaiWeight),
		// uintptr(data.Strategies.PossibleInTilesWeight),
		// uintptr(data.Strategies.TotalExpectValueWeight),
		// uintptr(data.Strategies.MaxExpectValueWeight),
		// uintptr(data.Strategies.PonActionExtraBounsWeight),
		// uintptr(data.Strategies.GonActionExtraBounsWeight),
		// uintptr(data.Strategies.AnGonActionExtraBounsWeight),
		// uintptr(data.Strategies.DiscardCount),
		// uintptr(data.Strategies.Rule),
	)

	// 0: 不動作 1: 胡 2: 碰 3: 槓 6: 暴擊胡
	result := int32(ret)

	fmt.Println("AIHuPongKong_result: ", result)

	return int(result)
}
