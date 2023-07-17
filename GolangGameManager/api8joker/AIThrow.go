package Api_8joker

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_AIThrow struct {
	Hand                  []int32
	Tile                  int32
	UnOpenPool            []int32
	Without               int32
	Pong                  []int32
	Kong                  []int32
	ConcealedKong         []int32
	SignleJokerAnGonCount int32
	GameRule              TransformGolang.CppGameRule32
	Strategies            TransformGolang.Strategies
	CheatTiles            []int32
}

func AIThrow(data Input_AIThrow) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	AIThrow := lib.NewProc("APIAIThrow")

	ret, _, _ := AIThrow.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(data.Tile),
		uintptr(unsafe.Pointer(&data.UnOpenPool[0])),
		uintptr(data.Without),
		uintptr(unsafe.Pointer(&data.Pong[0])),
		uintptr(unsafe.Pointer(&data.Kong[0])),
		uintptr(unsafe.Pointer(&data.ConcealedKong[0])),
		uintptr(data.SignleJokerAnGonCount),

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

	result := int32(ret)

	//fmt.Println("AIThrow: ", result)
	return int(result)
}
