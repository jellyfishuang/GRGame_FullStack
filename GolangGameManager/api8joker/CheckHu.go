package Api_8joker

import (
	TransformGolang "GolangGameManager/transform"
	"syscall"
	"unsafe"
)

type Input_CheckHu struct {
	Hand                  []int32
	Tile                  int32
	Pong                  []int32
	Kong                  []int32
	ConcealedKong         []int32
	SignleJokerAnGonCount int32
	Without               int32
	IsZimo                int32
	GameRule              TransformGolang.CppGameRule32
}

type Output_CheckHu struct {
	Faans      [FaanArraySize]int32
	Counts     [FaanArraySize]int32
	Error_Code int32
}

func DoCheckHu(data Input_CheckHu) Output_CheckHu {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	CheckHu := lib.NewProc("APICheckHu")
	free := lib.NewProc("FreeCheckHu")

	ret, _, _ := CheckHu.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(data.Tile),
		uintptr(unsafe.Pointer(&data.Pong[0])),
		uintptr(unsafe.Pointer(&data.Kong[0])),
		uintptr(unsafe.Pointer(&data.ConcealedKong[0])),
		uintptr(data.SignleJokerAnGonCount),
		uintptr(data.Without),
		uintptr(data.IsZimo),

		//uintptr(unsafe.Pointer(&data.GameRule)),

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

		// FanCount 跟 xortable 直接在 C++ API端宣告，而不透過 golang 每次呼叫 AI 傳入
		// 透過 golang 傳入 xortable (31684*4 byte) 會導致 stack overflow 和 memory leak 的問題

		//uintptr(unsafe.Pointer(&data.GameRule.FanCount[0])),
		//uintptr(unsafe.Pointer(&data.GameRule.JokerTurnIntos[0])),
		//uintptr(unsafe.Pointer(&data.GameRule.XorTable[0])),
	)

	result := *(*Output_CheckHu)(unsafe.Pointer(ret))

	// fmt.Println("returnValue.Faans", result.Faans)
	// fmt.Println("returnValue.Counts", result.Counts)
	// fmt.Println("returnValue.Error_Code", result.Error_Code)

	// free 掉回傳的 struct 指標
	free.Call(ret)

	return result
}
