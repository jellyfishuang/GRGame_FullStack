package gamemanager

import (
	DrawTool "GolangGameManager/drawtool"
	GameRule "GolangGameManager/gamerule"
	TransformGolang "GolangGameManager/transform"
	"math/rand"
	"sort"
)

type Player struct {
	// 呼叫 AI 相關API時使用
	strategies TransformGolang.Strategies // Strategies (AI 策略參數)
	//ApiDataPack
	//ApiDataPack TransformGolang.ApiDataPack // ApiDataPack (AI 資料)
	UnOpenPool            []int32 // 尚未顯現出的牌，包含牌海和玩家未打出的手牌
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
	// 玩家資料
	DealerWind   int                    // 自風
	WithoutTile  int                    // 定缺
	HandTiles    []int                  // 現在手上的牌
	FlowerTiles  []int                  // 抽到的花牌
	OnDrawTile   int                    // 摸到的牌
	DiscardTiles []int                  // 已經打過的牌
	ShowTiles    []int                  // 要亮在桌上的牌 (扣除被別家吃碰槓的牌)
	MeldTiles    []TransformGolang.Meld // 刻子
	Point        int                    // 目前番數
	Gamerule     GameRule.GameRule      // 遊戲規則
	// 提示玩家可以進行的動作
	CanChow             bool    // 可以吃
	CanChowSet          [][]int //可吃牌的組合
	CanPon              bool    // 可以碰
	CanPonSet           []int   //可碰牌的組合
	CanKong             bool    // 可以槓
	CanKongSet          [][]int //可槓牌的組合
	CanConcealedKong    bool    //暗槓
	CanConcealedKongSet [][]int
	CanAddKong          bool    //加槓
	CanAddKongSet       []int   //可加槓的組合
	CanHu               bool    // 可以胡
	CanHuFaanList       []int32 // 可以胡的番型
	CanHuFaanListStr    []string
	CanHuTile           int  // 可以胡的牌
	CanReadyHand        bool // 可以聽牌
	CanSelfHu           bool // 可以自摸
	CanSelfHuTile       int  // 可以自摸的牌
	DoKongTile          int  // AI選擇執行槓的牌
	ResetSelfHuTile     int
	CanBaoJi            bool // 可以使用暴擊
	CanMingPai          bool // 可以明牌
	// 接收玩家進行的動作
	Donothing       bool // 不做任何動作
	DoChow          bool // 吃
	DoPon           bool // 碰
	DoKong          bool // 槓
	DoConcealedKong bool //暗槓
	DoAddKong       bool //加槓
	DoHu            bool // 胡
	DoReadyHand     bool // 聽牌
	DoSelfHu        bool // 自摸
	DoBaoJi         bool // 暴擊
	DoMingPai       bool // 明牌
	AfterHu         bool // 胡牌後確認
	AfterSelfhu     bool // 自摸後丟牌確認
	// 確認聽牌(聽牌後不能做其他動作)
	IsZimo         int32 // 是否已經胡牌或自摸"過"
	IsReadyHand    bool  // 是否已經聽牌
	IsAlreadyHu    bool  // 是否已經胡過牌
	IsRiichi       bool  // 是否已經立直
	IsAlreadyBaoJi bool  // 是否使用過暴擊
	IsBaoJi        bool  // 是否使用暴擊
	IsMingPai      bool  // 是否明牌

	SignleJokerAnGonCount int32 //紅中單槓次數 (for 8joker玩法)

	//GameLog
	CsvChowTiles          []int
	CsvPongTiles          []int
	CsvKongTIles          []int
	CsvConcealedKongTiles []int
	CsvAddKongTiles       []int
}

func Remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// 初始化玩家
func (_player *Player) InitPlayer(count int, _gamerule GameRule.GameRule) {
	_player.strategies.SetStrategy(TransformGolang.SetCPPStrategy()) // strategies 初始化
	// ApiDataPack 初始化
	DrawTool.InitTileSea(&_player.UnOpenPool, _gamerule, false)
	_player.UnOpenPool = append(_player.UnOpenPool, -1)
	_player.Uid = 0
	_player.Without = 0
	_player.SingleJokerAnGonCount = 0
	_player.DoorSize = make([]int32, 0)
	_player.ThrowSeq = make([]int32, 0)
	_player.CheatTiles = make([]int32, 0)
	_player.AlreadyHuTiles = make([]int32, 0)
	_player.AlreadyHu = false
	_player.AlreadyMing = false
	_player.AlreadyCrit = false
	_player.IsTingCard = false
	_player.DealerWind = count
	_player.WithoutTile = 0 //-1
	_player.HandTiles = make([]int, 0)
	_player.OnDrawTile = -1
	_player.DiscardTiles = make([]int, 0)
	_player.FlowerTiles = make([]int, 0)
	_player.ShowTiles = make([]int, 0)
	_player.MeldTiles = make([]TransformGolang.Meld, 0)
	_player.Point = 0
	_player.CanChow = false
	_player.CanChowSet = make([][]int, 0)
	_player.CanPon = false
	_player.CanPonSet = make([]int, 0)
	_player.CanKong = false
	_player.CanKongSet = make([][]int, 0)
	_player.CanHu = false
	_player.CanHuFaanList = make([]int32, 0)
	_player.CanHuFaanListStr = make([]string, 0)
	_player.CanHuTile = -1
	_player.CanReadyHand = false
	_player.CanSelfHu = false
	_player.CanSelfHuTile = -1

	_player.DoChow = false
	_player.DoPon = false
	_player.DoKong = false
	_player.DoKongTile = -1
	_player.DoHu = false
	_player.DoReadyHand = false
	_player.DoSelfHu = false
	_player.AfterHu = false
	_player.AfterSelfhu = false

	_player.IsReadyHand = false
	_player.IsAlreadyHu = false
	_player.IsRiichi = false
	_player.IsAlreadyBaoJi = false
	_player.IsBaoJi = false

	_player.CsvChowTiles = make([]int, 0)
	_player.CsvPongTiles = make([]int, 0)
	_player.CsvKongTIles = make([]int, 0)
	_player.CsvConcealedKongTiles = make([]int, 0)
	_player.CsvAddKongTiles = make([]int, 0)

}

// 更新時機：初始發牌、抽牌、丟牌、吃、碰、槓
// 更新 UnOpenPool，將傳入的 _tiles 從 UnOpenPool 移除
func (_player *Player) UpdateUnOpenPool(_tiles []int) {
	_unopenpool := &(_player.UnOpenPool)

	for _, tile := range _tiles {
		for idx, pooltile := range *_unopenpool {
			if tile == int(pooltile) {
				//從UnOpenPool中刪除出現過的牌
				(*_unopenpool)[idx] = (*_unopenpool)[len((*_unopenpool))-2]
				(*_unopenpool)[len((*_unopenpool))-2] = -1
				*_unopenpool = (*_unopenpool)[:(len((*_unopenpool)) - 1)]
				break
			}
		}
	}
}

// 玩家換三張(交換階段)
func (_player *Player) ChangeThreeTiles(tiles []int) {
	_player.HandTiles = append(_player.HandTiles, tiles...)
}

// 真人玩家定缺
func (_player *Player) ChooseWithout(without int) {
	_player.WithoutTile = without
}

// 檢查玩家可用什麼手牌吃
// func (_player *Player) CheckChow(_tiles int64) [10]int {
// 	// Verify if the player can Chow the tile

// 	var inputValue ApiStruct.Input_CheckChow

// 	inputValue.Hand = TransformGolang.CppTiles(_player.HandTiles)
// 	inputValue.Tile = _tiles
// 	inputValue.Without = int64(_player.WithoutTile)
// 	inputValue.AlreadyHu = _player.IsAlreadyHu

// 	return ApiStruct.CheckChow(inputValue)
// }

// 檢查玩家是否可碰
// func (_player *Player) CheckPong(_tile int64) bool {
// 	//Verify if the player can Pong the tile

// 	var inputValue ApiStruct.Input_CheckPong

// 	inputValue.Hand = TransformGolang.CppTiles(_player.HandTiles)
// 	inputValue.Tile = _tile
// 	inputValue.Without = int64(_player.WithoutTile)
// 	inputValue.AlreadyHu = false // 有問題

// 	// CheckPong 回傳值若是負數範圍為參數錯誤，所以加入 == operator 防止此類型錯誤
// 	return ApiStruct.CheckPong(inputValue) == 1
// }

// 檢查玩家是否可槓
// func (_player *Player) CheckKong(_tile int64) bool {
// 	// Verify if the player can Kong the tile

// 	var inputValue ApiStruct.Input_CheckKong

// 	inputValue.Hand = TransformGolang.CppTiles(_player.HandTiles)
// 	inputValue.Tile = _tile
// 	inputValue.Without = int64(_player.WithoutTile)
// 	inputValue.AlreadyHu = false // 有問題

// 	// CheckKong 回傳值若是負數範圍為參數錯誤，所以加入 == operator 防止此類型錯誤
// 	return ApiStruct.CheckKong(inputValue) == 1
// }

// // 檢查所有玩家是否可胡牌，以及胡牌之後的番型與番數
// func (_player *Player) CheckHu(_tile int32, _isZimo int32, _gamerule GameRule.GameRule) (ok bool) {
// 	// Verify if the player can Hu the tile with different gamerules
// 	ok = true
// 	if _gamerule.GameMode == GameRule.Mode8Joker {
// 		// println("8Joker_CheckHu")
// 		var parm AI_8JOKER.Input_CheckHu

// 		// 把手牌轉成int32
// 		for _, tile := range _player.HandTiles {
// 			parm.Hand = append(parm.Hand, int32(tile))
// 		}
// 		parm.IsZimo = _isZimo
// 		parm.Tile = _tile
// 		parm.Without = int32(_player.WithoutTile)
// 		for _, tile
// 		parm.Pong, parm.Kong, parm.ConcealedKong = []int32{-1}, []int32{-1}, []int32{-1}
// 		parm.SignleJokerAnGonCount = 0
// 		parm.GameRule.SetGameRule32(_gamerule)

// 		Output_CheckHu := AI_8JOKER.DoCheckHu(parm)

// 		var faanidx int

// 		for faanidx = 0; faanidx < len(Output_CheckHu.Faans); faanidx++ {
// 			if Output_CheckHu.Faans[faanidx] == -1 {
// 				break
// 			}
// 			_player.CanHuFaanList = append(_player.CanHuFaanList, Output_CheckHu.Faans[faanidx])
// 		}
// 		// fmt.Println("Output_CheckHu_Faans:", Output_CheckHu.Faans)
// 		// fmt.Println("Output_CheckHu_Counts:", Output_CheckHu.Counts)
// 		// fmt.Println("Output_CheckHu_Error_Code:", Output_CheckHu.Error_Code)

// 		if faanidx == 0 {
// 			// 代表沒有胡任何番型 = 沒有胡
// 			return false
// 		} else {
// 			// 代表有胡
// 			for i := 0; i < faanidx; i++ {
// 				// 存取番型名稱
// 				_player.CanHuFaanListStr = append(_player.CanHuFaanListStr, AI_8JOKER.GetFaanName(int(_player.CanHuFaanList[i])))
// 				// fmt.Println(AI_8JOKER.GetFaanName(int(_player.CanHuFaanList[i])))
// 			}
// 			return true
// 		}
// 	}

// 	return false
// }

// 玩家吃
func (_player *Player) Chow(_tiles []int, _ontabletile int) {
	// to do Chow action in this player

	// MeldTiles
	meldtile := []int{_tiles[0], _tiles[1], _ontabletile}
	sort.Slice(meldtile, func(i, j int) bool {
		return meldtile[i] < meldtile[j]
	})
	_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 0, Tiles: meldtile})
	// HandTiles
	for i := 0; i < 3; i++ {
		for idx, tile := range _player.HandTiles {
			if tile == _tiles[i] {
				_player.HandTiles = Remove(_player.HandTiles, idx)
				break
			}
		}
	}
}

// 玩家碰
func (_player *Player) Pong(_tile int) {
	// to do Pong action in this player

	// MeldTiles
	_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 1, Tiles: []int{_tile}})
	// HandTiles
	for times := 0; times < 2; times++ {
		for idx, tile := range _player.HandTiles {
			if tile == _tile {
				_player.HandTiles = Remove(_player.HandTiles, idx)
				break
			}
		}
	}
}

// 玩家槓
func (_player *Player) Kong(_tile int) {
	// to do Kong action in this player

	// MeldTiles
	/*if isConcealedKong {
		_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 3, Tiles: []int{_tile}})
	} else {
		_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 2, Tiles: []int{_tile}})
	}*/
	_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 2, Tiles: []int{_tile}})
	// HandTiles
	for times := 0; times < 3; times++ {
		for idx, tile := range _player.HandTiles {
			if tile == _tile {
				_player.HandTiles = Remove(_player.HandTiles, idx)
				break
			}
		}
	}
}

func (_player *Player) ConcealedKong(_tile int) {
	_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 3, Tiles: []int{_tile}})

	removetimes := 0
	for times := 0; times < 4; times++ {
		for idx, tile := range _player.HandTiles {
			if tile == _tile {
				_player.HandTiles = Remove(_player.HandTiles, idx)
				removetimes++
				break
			}
		}
	}

	// 針對暗槓掉沒有在進牌的狀況
	// 把進牌放進手牌
	if removetimes == 4 {
		_player.HandTiles = append(_player.HandTiles, _player.OnDrawTile)
		_player.OnDrawTile = -1
	}
}

func (_player *Player) AddKong(_tile int) {
	_player.MeldTiles = append(_player.MeldTiles, TransformGolang.Meld{Action: 4, Tiles: []int{_tile}})

	for times := 0; times < 1; times++ {
		for idx, tile := range _player.HandTiles {
			if tile == _tile {
				_player.HandTiles = Remove(_player.HandTiles, idx)
				break
			}
		}
	}
}

// 玩家丟牌, 對玩家資訊的修改
func (_player *Player) PlayerDiscardTile(_tile int) (ok bool) {

	// fmt.Println("PlayerDiscardTile", _tile)
	// fmt.Println("PlayerOnDrawTile(前)", _player.OnDrawTile)
	// fmt.Println("PlayerHandTiles(前)", _player.HandTiles)

	// 丟牌之前, 可以把手上剛進的牌(如果有)放進手牌, 好用來再接著判斷丟牌
	if _player.OnDrawTile != -1 {
		_player.HandTiles = append(_player.HandTiles, _player.OnDrawTile)
		_player.OnDrawTile = -1
	}

	// fmt.Println("PlayerOnDrawTile", _player.OnDrawTile)
	// fmt.Println("PlayerHandTiles", _player.HandTiles)

	//fmt.Println("玩家", _player.DealerWind, "丟牌: ", _tile)

	// 檢查是否有此牌
	ok = false
	for idx, tile := range _player.HandTiles {
		if tile == _tile {
			// 丟牌
			_player.HandTiles[idx] = _player.HandTiles[len(_player.HandTiles)-1]
			_player.HandTiles = _player.HandTiles[:len(_player.HandTiles)-1]
			_player.DiscardTiles = append(_player.DiscardTiles, _tile)
			ok = true
			break
		}
	}

	return ok
}

func (_player *Player) ResetAction() {
	_player.CanChow = false
	_player.CanChowSet = nil
	_player.CanPon = false
	_player.CanPonSet = nil
	_player.CanKong = false
	_player.CanKongSet = nil
	_player.CanConcealedKong = false
	_player.CanConcealedKongSet = nil
	_player.CanAddKong = false
	_player.CanHu = false
	_player.CanHuFaanList = nil
	_player.CanHuFaanListStr = nil
	_player.CanHuTile = -1
	_player.CanReadyHand = false
	_player.CanSelfHu = false
	_player.CanSelfHuTile = -1
	_player.DoKongTile = -1

	_player.Donothing = false
	_player.DoChow = false
	_player.DoPon = false
	_player.DoKong = false
	_player.DoConcealedKong = false
	_player.DoAddKong = false
	_player.DoHu = false
	_player.DoReadyHand = false
	_player.DoSelfHu = false
}

func (_player *Player) CalculateFanCounts(_tile uint8) int {
	// to do calculate fan counts

	return 0
}

// ForAItest
func SetSeat(players []*Player) {
	n := len(players)
	for i := 0; i < n; i++ {
		newIdx := i + rand.Intn(n-i)
		players[i], players[newIdx] = players[newIdx], players[i]
	}
}
