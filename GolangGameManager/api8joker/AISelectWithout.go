package Api_8joker

import (
	"syscall"
	"unsafe"
)

type Input_AISelectWithout struct {
	Hand       []int32
	CheatTiles []int32
}

func AISelectWithOut(data Input_AISelectWithout) int {
	lib := syscall.NewLazyDLL(DLL_LOCATION)
	SelectWithout := lib.NewProc("AISelectWithout")

	ret, _, _ := SelectWithout.Call(
		uintptr(unsafe.Pointer(&data.Hand[0])),
		uintptr(unsafe.Pointer(&data.CheatTiles[0])),
	)

	result := int32(ret)

	//fmt.Println("AISelectWithout: ", result)
	return int(result)
}
