package Api_8joker

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_AISelectChangeTile struct {
	Hand       []int32
	GameRule   TransformGolang.CppGameRule32
	Strategies TransformGolang.Strategies
	CheatTiles []int32
}

type Output_AISelectChangeTiles struct {
	Array      [ArrayStructSize]int32
	Error_code int32
}

func AISelectChangeTile(data Input_AISelectChangeTile) [10]int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	SelectChangeTiles := lib.NewProc("APIAISelectChangeTile")
	free := lib.NewProc("FreeAISelectChangeTile")

	// fmt.Println("data.cheattile: ", data.CheatTiles)
	// fmt.Println("data.Strategies: ", data.Strategies)
	ret, _, _ := SelectChangeTiles.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(unsafe.Pointer(&data.CheatTiles[0])),

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
	)

	result := *(*Output_AISelectChangeTiles)(unsafe.Pointer(ret))
	//fmt.Println("result:", result)

	var returnArray [ArrayStructSize]int
	for i := 0; i < ArrayStructSize; i++ {
		returnArray[i] = int(result.Array[i])
	}

	free.Call(ret)

	return returnArray
}
