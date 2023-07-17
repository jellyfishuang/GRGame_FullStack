package TransformGolang

import (
	GameRule "GolangGameManager/gamerule"
	"syscall"
	"unsafe"
)

type Meld struct {
	Action int // 刻子(0:吃, 1:碰, 2:槓, 3:暗槓)
	Tiles  []int
}

type CppGameRule struct {
	ChangeTileSameColor      bool
	ChangeTileCount          int64
	CanEat                   bool
	CanPong                  bool
	CanKong                  bool
	CanJokersingleGon        bool
	CanJokerGon              bool
	CanMing                  bool
	LogAICsv                 bool
	CanCrit                  bool
	XorTable                 []int64
	FanCount                 []int64
	JokerTurnIntos           []int64
	LimitTai                 int64
	KongAfterHuUseSimpleRule bool
	Debug                    bool
	GuoShouHu                bool
	LogAIParameter           bool
	CanUseJokerAsEye         bool
	JokerNumbers             int64
	GameMode                 int64
	CanTing                  bool
	SpecialMode              bool
}

type CppGameRule32 struct {
	ChangeTileSameColor      bool
	ChangeTileCount          int32
	CanEat                   bool
	CanPong                  bool
	CanKong                  bool
	CanJokersingleGon        bool
	CanJokerGon              bool
	CanMing                  bool
	LogAICsv                 bool
	CanCrit                  bool
	LimitTai                 int32
	KongAfterHuUseSimpleRule bool
	Debug                    bool
	GuoShouHu                bool
	LogAIParameter           bool
	CanUseJokerAsEye         bool
	JokerNumbers             int32
	GameMode                 int32
	CanTing                  bool
	SpecialMode              bool
	FanCount                 []int32
	JokerTurnIntos           []int32
	XorTable                 []int32
}

type CppStrategy struct {
	HuDistance            int64
	FaanTaiWeight         float64
	DistanceWeight        float64
	PossibleInTilesWeight float64
	DiscardCount          int64
	Rule                  int64
}

type CppStrategy32 struct {
	HuDistance                  int32
	FaanTaiWeight               float32
	DistanceWeight              float32
	PossibleInTilesWeight       float32
	MaxExpectValueWeight        float32
	TotalExpectValueWeight      float32
	EatActionExtraBounsWeight   float32
	PonActionExtraBounsWeight   float32
	GonActionExtraBounsWeight   float32
	AnGonActionExtraBounsWeight float32
	HuActionExtraBounsWeight    float32
	AIThrowStrategy             int32
	DiscardCount                int32
	Rule                        int32
	UseRuleBasedStrategy        bool
	PhaseParameter              []int32
}

// 針對 Call C++ API的部分把策略參數轉成int後傳入, 再由C++ API內轉成float後使用
type Strategies struct {
	HuDistance                  int32
	FaanTaiWeight               int32
	DistanceWeight              int32
	PossibleInTilesWeight       int32
	MaxExpectValueWeight        int32
	TotalExpectValueWeight      int32
	EatActionExtraBounsWeight   int32
	PonActionExtraBounsWeight   int32
	GonActionExtraBounsWeight   int32
	AnGonActionExtraBounsWeight int32
	HuActionExtraBounsWeight    int32
	AIThrowStrategy             int32
	DiscardCount                int32
	Rule                        int32
	UseRuleBasedStrategy        bool
	PhaseParameter              []int32
}

func (_strategy *Strategies) SetStrategy(CppStrategy CppStrategy32) {

	_strategy.HuDistance = CppStrategy.HuDistance
	_strategy.FaanTaiWeight = int32(CppStrategy.FaanTaiWeight * 100)
	_strategy.DistanceWeight = int32(CppStrategy.DistanceWeight * 100)
	_strategy.PossibleInTilesWeight = int32(CppStrategy.PossibleInTilesWeight * 100)
	_strategy.MaxExpectValueWeight = int32(CppStrategy.MaxExpectValueWeight * 100)
	_strategy.TotalExpectValueWeight = int32(CppStrategy.TotalExpectValueWeight * 100)
	_strategy.EatActionExtraBounsWeight = int32(CppStrategy.EatActionExtraBounsWeight * 100)
	_strategy.PonActionExtraBounsWeight = int32(CppStrategy.PonActionExtraBounsWeight * 100)
	_strategy.GonActionExtraBounsWeight = int32(CppStrategy.GonActionExtraBounsWeight * 100)
	_strategy.AnGonActionExtraBounsWeight = int32(CppStrategy.AnGonActionExtraBounsWeight * 100)
	_strategy.HuActionExtraBounsWeight = int32(CppStrategy.HuActionExtraBounsWeight * 100)
	_strategy.AIThrowStrategy = CppStrategy.AIThrowStrategy
	_strategy.DiscardCount = CppStrategy.DiscardCount
	_strategy.Rule = CppStrategy.Rule
	_strategy.UseRuleBasedStrategy = CppStrategy.UseRuleBasedStrategy
}

func (gr *CppGameRule) SetGameRule(_gamerule GameRule.GameRule) {
	gr.ChangeTileCount = int64(_gamerule.ChangeTileCount)
	gr.ChangeTileSameColor = _gamerule.ChangeTileSameColor
	gr.CanEat = _gamerule.CanEat
	gr.CanKong = _gamerule.CanKong
	gr.CanPong = _gamerule.CanPong
	gr.CanJokersingleGon = _gamerule.CanJokersingleGon
	gr.CanJokerGon = _gamerule.CanJokerGon
	gr.CanMing = _gamerule.CanMing
	gr.LogAICsv = _gamerule.LogAICsv
	gr.CanCrit = _gamerule.CanCrit

	gr.XorTable = make([]int64, len(_gamerule.XorTable))
	for i, v := range _gamerule.XorTable {
		gr.XorTable[i] = int64(v)
	}
	//gr.XorTable = _gamerule.XorTable
	gr.FanCount = make([]int64, len(_gamerule.FanCount))
	for i, v := range _gamerule.FanCount {
		gr.FanCount[i] = int64(v)
	}
	//gr.FanCount = _gamerule.FanCount
	gr.JokerTurnIntos = make([]int64, len(_gamerule.JokerTurnIntos))
	for i, v := range _gamerule.JokerTurnIntos {
		gr.JokerTurnIntos[i] = int64(v)
	}
	//gr.JokerTurnIntos = _gamerule.JokerTurnIntos

	gr.LimitTai = int64(_gamerule.LimitTai)
	gr.KongAfterHuUseSimpleRule = _gamerule.KongAfterHuUseSimpleRule
	gr.Debug = _gamerule.Debug
	gr.GuoShouHu = _gamerule.GuoShouHu
	gr.LogAIParameter = _gamerule.LogAIParameter
	gr.CanUseJokerAsEye = _gamerule.CanUseJokerAsEye
	gr.JokerNumbers = int64(_gamerule.JokerNumbers)
	gr.GameMode = int64(_gamerule.GameMode)
	gr.CanTing = _gamerule.CanTing
	gr.SpecialMode = _gamerule.SpecialMode

}

func (gr *CppGameRule32) SetGameRule32(_gamerule GameRule.GameRule) {
	gr.ChangeTileCount = int32(_gamerule.ChangeTileCount)
	gr.ChangeTileSameColor = _gamerule.ChangeTileSameColor
	gr.CanEat = _gamerule.CanEat
	gr.CanKong = _gamerule.CanKong
	gr.CanPong = _gamerule.CanPong
	gr.CanJokersingleGon = _gamerule.CanJokersingleGon
	gr.CanJokerGon = _gamerule.CanJokerGon
	gr.CanMing = _gamerule.CanMing
	gr.LogAICsv = _gamerule.LogAICsv
	gr.CanCrit = _gamerule.CanCrit

	gr.XorTable = make([]int32, len(_gamerule.XorTable))
	for i, v := range _gamerule.XorTable {
		gr.XorTable[i] = int32(v)
	}

	gr.FanCount = make([]int32, len(_gamerule.FanCount))
	for i, v := range _gamerule.FanCount {
		gr.FanCount[i] = int32(v)
	}

	gr.JokerTurnIntos = make([]int32, len(_gamerule.JokerTurnIntos))
	for i, v := range _gamerule.JokerTurnIntos {
		gr.JokerTurnIntos[i] = int32(v)
	}
	// gr.XorTable = _gamerule.XorTable
	// gr.FanCount = _gamerule.FanCount
	// gr.JokerTurnIntos = _gamerule.JokerTurnIntos

	gr.LimitTai = int32(_gamerule.LimitTai)
	gr.KongAfterHuUseSimpleRule = _gamerule.KongAfterHuUseSimpleRule
	gr.Debug = _gamerule.Debug
	gr.GuoShouHu = _gamerule.GuoShouHu
	gr.LogAIParameter = _gamerule.LogAIParameter
	gr.CanUseJokerAsEye = _gamerule.CanUseJokerAsEye
	gr.JokerNumbers = int32(_gamerule.JokerNumbers)
	gr.GameMode = int32(_gamerule.GameMode)
	gr.CanTing = _gamerule.CanTing
	gr.SpecialMode = _gamerule.SpecialMode
}

func SetCPPStrategy() (_strategy CppStrategy32) {
	_strategy.HuDistance = 2

	_strategy.DistanceWeight = 0.6
	_strategy.FaanTaiWeight = 4.0
	_strategy.PossibleInTilesWeight = 0.45

	_strategy.TotalExpectValueWeight = 0.89
	_strategy.MaxExpectValueWeight = 1.26

	_strategy.PonActionExtraBounsWeight = 1.0
	_strategy.GonActionExtraBounsWeight = 1.0
	_strategy.AnGonActionExtraBounsWeight = 1.0

	_strategy.DiscardCount = 0
	_strategy.Rule = 0

	return _strategy
}

// Golang轉換C++
func CppTiles[T int | int64](_tiles []T) (cpptiles []int64) {
	cpptiles = make([]int64, len(_tiles)+1)
	for idx, tile := range _tiles {
		cpptiles[idx] = int64(tile)
	}
	cpptiles[len(_tiles)] = -1

	return cpptiles
}

func CppTiles32[T int | int32](_tiles []T) (cpptiles []int32) {
	cpptiles = make([]int32, len(_tiles)+1)
	for idx, tile := range _tiles {
		cpptiles[idx] = int32(tile)
	}
	cpptiles[len(_tiles)] = -1

	return cpptiles
}

func CppMeldTiles(_meldTiles []Meld) (Chow []int64, Pong []int64, Kong []int64, ConcealedKong []int64) {
	for _, tile := range _meldTiles {
		switch tile.Action {
		case 0:
			Chow = append(Chow, int64(tile.Tiles[0]), int64(tile.Tiles[1]), int64(tile.Tiles[2]))
		case 1:
			Pong = append(Pong, int64(tile.Tiles[0]), int64(tile.Tiles[0]), int64(tile.Tiles[0]))
		case 2:
			Kong = append(Kong, int64(tile.Tiles[0]))
		case 3:
			ConcealedKong = append(ConcealedKong, int64(tile.Tiles[0]))
		}
	}
	Chow = CppTiles(Chow)
	Pong = CppTiles(Pong)
	Kong = CppTiles(Kong)
	ConcealedKong = CppTiles(ConcealedKong)

	return Chow, Pong, Kong, ConcealedKong
}

func CppMeldTiles32(_meldTiles []Meld) (Chow []int32, Pong []int32, Kong []int32, ConcealedKong []int32) {
	for _, tile := range _meldTiles {
		switch tile.Action {
		case 0:
			Chow = append(Chow, int32(tile.Tiles[0]), int32(tile.Tiles[1]), int32(tile.Tiles[2]))
		case 1:
			Pong = append(Pong, int32(tile.Tiles[0]), int32(tile.Tiles[0]), int32(tile.Tiles[0]))
		case 2:
			Kong = append(Kong, int32(tile.Tiles[0]))
		case 3:
			ConcealedKong = append(ConcealedKong, int32(tile.Tiles[0]))
		}
	}
	Chow = CppTiles32(Chow)
	Pong = CppTiles32(Pong)
	Kong = CppTiles32(Kong)
	ConcealedKong = CppTiles32(ConcealedKong)

	return Chow, Pong, Kong, ConcealedKong
}

func IntPtr(n int) uintptr {
	return uintptr(unsafe.Pointer(&n))
}

func Int32Ptr(n int32) uintptr {
	return uintptr(unsafe.Pointer(&n))
}

func Int64Ptr(n int64) uintptr {
	return uintptr(unsafe.Pointer(&n))
}

func Float64Ptr(n float64) uintptr {
	return uintptr(unsafe.Pointer(&n))
}

func IntArrayPtr(s []int64) uintptr {
	return uintptr(unsafe.Pointer(&s[0]))
}

func Int32ArrayPtr(s []int32) uintptr {
	return uintptr(unsafe.Pointer(&s[0]))
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func BoolPtr(s bool) uintptr {
	return uintptr(unsafe.Pointer(&s))
}
