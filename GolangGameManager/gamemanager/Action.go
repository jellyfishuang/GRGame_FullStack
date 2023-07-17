package gamemanager

type Action struct {
	Index int
	Tiles uint8
}

const (
// _BuHua         = 1 // 補花
// _ChangeThree   = 2 // 換三張
// _ChooseWithout = 3 // 定缺

// _DrawTile = 4 // 進牌
// _Discard  = 5 // 丟牌

// _DealCard = 6 // 發牌

// _CountPoints = 100 // 計算番數(血流的話不能放在結算階段, 盡量在胡牌之後就計算番數存起來)
)
