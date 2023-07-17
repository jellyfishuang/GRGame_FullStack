package Api_Public

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Input_CheckHu struct {
	Hand                  []int32
	Tile                  int32
	Chow                  []int32
	Pong                  []int32
	Kong                  []int32
	ConcealedKong         []int32
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
}

type Output_CheckHu struct {
	Faans      [FaanArraySize]int32
	Counts     [FaanArraySize]int32
	Error_Code int32
}

func DoCheckHu(data Input_CheckHu) Output_CheckHu {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	CheckHu := lib.NewProc("APICheckHu")
	Free := lib.NewProc("FreeCheckHu")

	ret, _, _ := CheckHu.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(data.Tile),
		uintptr(unsafe.Pointer(&data.Chow[0])),
		uintptr(unsafe.Pointer(&data.Pong[0])),
		uintptr(unsafe.Pointer(&data.Kong[0])),
		uintptr(unsafe.Pointer(&data.ConcealedKong[0])),
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

		// FanCount 跟 xortable 直接在 C++ API端宣告，而不透過 golang 每次呼叫 AI 傳入
		// 透過 golang 傳入 xortable (31684*4 byte) 會導致 stack overflow 和 memory leak 的問題

		//uintptr(unsafe.Pointer(&data.GameRule.FanCount[0])),
		//uintptr(unsafe.Pointer(&data.GameRule.JokerTurnIntos[0])),
		//uintptr(unsafe.Pointer(&data.GameRule.XorTable[0])),
	)

	result := *(*Output_CheckHu)(unsafe.Pointer(ret))

	fmt.Println("returnValue.Faans", result.Faans)
	fmt.Println("returnValue.Counts", result.Counts)
	fmt.Println("returnValue.Error_Code", result.Error_Code)

	// free 掉回傳的 struct 指標
	Free.Call(ret)

	return result
}
