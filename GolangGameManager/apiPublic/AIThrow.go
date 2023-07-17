package Api_Public

import (
	TransformGolang "GolangGameManager/transform"
	"fmt"
	"syscall"
	"unsafe"
)

type Input_AIThrow struct {
	Hand                  []int32
	Tile                  int32
	Chow                  []int32
	Pong                  []int32
	Kong                  []int32
	ConcealedKong         []int32
	UnOpenPool            []int32
	Strategies            TransformGolang.Strategies
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
}

func AIThrow(data Input_AIThrow) int {
	fmt.Println("calling Public AIThorw")
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	APIAIThrow := lib.NewProc("APIAIThrow")
	fmt.Println("+++++++NewProc:", APIAIThrow, "+++++++")
	ret, _, _ := APIAIThrow.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(data.Tile),
		uintptr(unsafe.Pointer(&data.Chow[0])),
		uintptr(unsafe.Pointer(&data.Pong[0])),
		uintptr(unsafe.Pointer(&data.Kong[0])),
		uintptr(unsafe.Pointer(&data.ConcealedKong[0])),
		uintptr(unsafe.Pointer(&data.UnOpenPool[0])),

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

		// 傳入 GameRule
		// uintptr(unsafe.Pointer(&data.GameRule.ChangeTileSameColor)),
		// uintptr(data.GameRule.ChangeTileCount),
		//uintptr(unsafe.Pointer(&data.GameRule.CanEat)),
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
		uintptr(unsafe.Pointer(&data.Strategies)),
		//uintptr(data.Strategies.HuDistance),
		// uintptr(data.Strategies.FaanTaiWeight),
		// uintptr(data.Strategies.DistanceWeight),
		// uintptr(data.Strategies.PossibleInTilesWeight),
		// uintptr(data.Strategies.MaxExpectValueWeight),
		// uintptr(data.Strategies.TotalExpectValueWeight),
		// uintptr(data.Strategies.EatActionExtraBounsWeight),
		// uintptr(data.Strategies.PonActionExtraBounsWeight),
		// uintptr(data.Strategies.GonActionExtraBounsWeight),
		// uintptr(data.Strategies.AnGonActionExtraBounsWeight),
		// uintptr(data.Strategies.HuActionExtraBounsWeight),
		// uintptr(data.Strategies.AIThrowStrategy),
		// uintptr(data.Strategies.DiscardCount),
		// uintptr(data.Strategies.Rule),
		//uintptr(unsafe.Pointer(&data.Strategies.UseRuleBasedStrategy)),
		//uintptr(unsafe.Pointer(&data.Strategies.PhaseParameter[0])),
	)

	result := int32(ret)

	//fmt.Println("AIThrow: ", result)

	return int(result)

}
