package Api_8joker

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_AIHuPongKong struct {
	Hand                  []int32
	Tile                  int32
	UnOpenPool            []int32
	Without               int32
	Pong                  []int32
	Kong                  []int32
	ConcealedKong         []int32
	SignleJokerAnGonCount int32
	AlreadyHu             bool
	AlreadyMing           bool
	AlreadyCrit           bool
	GameRule              TransformGolang.CppGameRule32
	Strategies            TransformGolang.Strategies
	CheatTiles            []int32
}

func AIHuPongKong(data Input_AIHuPongKong) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	AIHuPongKong := lib.NewProc("APIAIHuPongKong")

	ret, _, _ := AIHuPongKong.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(data.Tile),
		uintptr(unsafe.Pointer(&data.UnOpenPool[0])),
		uintptr(data.Without),
		uintptr(unsafe.Pointer(&data.Pong[0])),
		uintptr(unsafe.Pointer(&data.Kong[0])),
		uintptr(unsafe.Pointer(&data.ConcealedKong[0])),
		uintptr(data.SignleJokerAnGonCount),
		uintptr(TransformGolang.BoolPtr(data.AlreadyHu)),
		uintptr(TransformGolang.BoolPtr(data.AlreadyMing)),
		uintptr(TransformGolang.BoolPtr(data.AlreadyCrit)),

		uintptr(unsafe.Pointer(&data.CheatTiles[0])),

		// 傳入 GameRule
		// uintptr(unsafe.Pointer(&data.GameRule.ChangeTileSameColor)),
		// uintptr(data.GameRule.ChangeTileCount),
		// uintptr(unsafe.Pointer(&data.GameRule.CanEat)),
		// uintptr(unsafe.Pointer(&data.GameRule.CanPong)),
		// uintptr(unsafe.Pointer(&data.GameRule.CanKong)),
		// uintptr(unsafe.Pointer(&data.GameRule.CanJokersingleGon)),
		// uintptr(unsafe.Pointer(&data.GameRule.CanJokerGon)),
		// uintptr(unsafe.Pointer(&data.GameRule.CanMing)),
		// uintptr(unsafe.Pointer(&data.GameRule.LogAICsv)),
		// uintptr(unsafe.Pointer(&data.GameRule.CanCrit)),
		// uintptr(data.GameRule.LimitTai),
		// uintptr(unsafe.Pointer(&data.GameRule.KongAfterHuUseSimpleRule)),
		// uintptr(unsafe.Pointer(&data.GameRule.Debug)),
		// uintptr(unsafe.Pointer(&data.GameRule.GuoShouHu)),
		// uintptr(unsafe.Pointer(&data.GameRule.LogAIParameter)),
		// uintptr(data.GameRule.JokerNumbers),
		// uintptr(data.GameRule.GameMode),
		// uintptr(unsafe.Pointer(&data.GameRule.CanTing)),
		// uintptr(unsafe.Pointer(&data.GameRule.SpecialMode)),
		// uintptr(unsafe.Pointer(&data.GameRule.JokerTurnIntos[0])),

		// 傳入 Strategies
		uintptr(data.Strategies.HuDistance),
		uintptr(data.Strategies.DistanceWeight),
		uintptr(data.Strategies.FaanTaiWeight),
		uintptr(data.Strategies.PossibleInTilesWeight),
		uintptr(data.Strategies.TotalExpectValueWeight),
		uintptr(data.Strategies.MaxExpectValueWeight),
		uintptr(data.Strategies.PonActionExtraBounsWeight),
		uintptr(data.Strategies.GonActionExtraBounsWeight),
		uintptr(data.Strategies.AnGonActionExtraBounsWeight),
		uintptr(data.Strategies.DiscardCount),
		uintptr(data.Strategies.Rule),
	)

	// 0: 不動作 1: 胡 2: 碰 3: 槓 6: 暴擊胡
	result := int32(ret)

	//fmt.Println("AIHuPongKong_result: ", result)

	return int(result)
}
