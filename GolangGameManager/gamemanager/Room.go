package gamemanager

import (
	DrawTool "GolangGameManager/drawtool"
	GameRule "GolangGameManager/gamerule"
	TransformGolang "GolangGameManager/transform"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

type State int

const (
	// 遊戲狀態
	StateConnect      = -1 // 連線
	StateCreateRoom   = 0  // 創建房間,設定玩法,發牌
	StateDealTiles    = 1  // 發牌
	StateChangeTiles  = 2  // 換牌
	StateWithoutTiles = 3  // 定缺
	StateBuHua        = 4  // 補花
	StateOpenGame     = 5  // 開門(發第一張牌給莊家)
	StateDraw         = 6  // 摸牌
	StateAction       = 7  // 動作
	StateDiscard      = 8  // 丟牌
	StateGameOver     = 9  // 遊戲結束
)

type doAction struct {
	PriviousPlayer int    //上一個玩家
	Player         int    //玩家
	Action         string //執行的動作(ex:吃碰槓)
	Tiles          []int  //執行動作的牌(ex:吃,[3,4,5])
}

type Room struct {
	RoomId         int64
	TilesSea       []int // 牌海(還沒發出去的牌)
	PrevailingWind int   // 場風
	PlayerNow      int
	RealPlayer     int // 真人玩家所代表的位置
	ContinueCount  int // 連莊次數
	RoundCount     int // 第幾輪(每丟一張牌就+1)
	Gameover       bool
	DoAction       doAction  //玩家實際執行的動作
	OntableTile    int       //被丟出來的牌(不為-1時，表示為PlayerNow丟出的牌)
	Players        [4]Player // 玩家
	NothingDo      bool      // 所有玩家沒有動作
	GameRule       GameRule.GameRule
	GameState      int // 遊戲狀態
	GameLog        [][]string
}

func NewRoom(_roomID int64) *Room {
	return &Room{RoomId: _roomID}
}

func (gm *GameManager) FindRoom(_roomID int64) *Room {
	for _, room := range gm.Rooms {
		if room.RoomId == _roomID {
			return room
		}
	}
	return nil
}

// 創建房間
func (_room *Room) InitRoom(_gamerule string, AItest bool, AutoTestRound int) (ok bool) {
	ok = true

	// 真人玩與自動測試的房間ID不同
	if AItest {
		_room.RoomId = int64(AutoTestRound)
		_room.RealPlayer = -1
	} else {
		_room.RoomId = 9999
		_room.RealPlayer = 0
	}

	_room.PrevailingWind = 28 // 場風預設 東風

	_room.PlayerNow = 0
	_room.ContinueCount = 0
	_room.RoundCount = 0
	_room.Gameover = false
	_room.OntableTile = -1
	_room.GameState = StateCreateRoom
	_room.NothingDo = false

	_room.GameRule.SetGameRule(_gamerule)
	//fmt.Println("GameMode", _room.GameRule.GameMode)

	// 初始化玩家
	for player := 0; player < 4; player++ {
		_room.Players[player].InitPlayer(player, _room.GameRule)
	}

	//fmt.Print("房間初始化完成\n")
	DrawTool.InitTileSea(&_room.TilesSea, _room.GameRule, true)
	//fmt.Println("牌海:(洗牌後) ", _room.TilesSea)

	_room.GameLog = nil

	return ok
}

func (_room *Room) DealTiles() {

	_room.Players[0].HandTiles, _room.Players[1].HandTiles, _room.Players[2].HandTiles,
		_room.Players[3].HandTiles = DrawTool.DealTile(_room.PlayerNow, &_room.TilesSea, int(_room.GameRule.TileUpperLimit))

	for idx := range _room.Players {
		_room.Players[idx].UpdateUnOpenPool(_room.Players[idx].HandTiles)
	}

	// 顯示發牌
	for player := 0; player < 4; player++ {
		//fmt.Println("玩家", player, "手牌: ", _room.Players[player].HandTiles)
	}

	// 花牌進行補花並更新手牌
	if _room.GameRule.IsbuHua {
		for player := 0; player < 4; player++ {
			// 搜尋玩家的手牌檢查是否為花牌
			playerhaveflower := 0
			for _, tile := range _room.Players[player].HandTiles {
				if _room.IsFlowerTile(tile) {
					_room.Players[player].FlowerTiles = append(_room.Players[player].FlowerTiles, tile)
					playerhaveflower++
				}
			}
			// 確認花牌數量大於 0 ，從手牌取出並放入花牌後，重新抽牌
			if playerhaveflower > 0 {
				_room.Players[player].HandTiles = deleteElements(_room.Players[player].HandTiles, _room.Players[player].FlowerTiles)

				// 補抽牌
				for i := 0; i < playerhaveflower; i++ {
					ok, _drawtile := DrawTool.DrawTile(player, &_room.TilesSea)
					if ok && _room.IsFlowerTile(_drawtile) {
						_room.Players[player].FlowerTiles = append(_room.Players[player].FlowerTiles, _drawtile)
						// 又抽到花牌 計數器 -1 重新再抽一張牌
						i--
					} else if ok && !_room.IsFlowerTile(_drawtile) {
						_room.Players[player].HandTiles = append(_room.Players[player].HandTiles, _drawtile)
					}
				}
			}

			//重新排列玩家手牌
			sort.Slice(_room.Players[player].HandTiles, func(i, j int) bool {
				return _room.Players[player].HandTiles[i] < _room.Players[player].HandTiles[j]
			})

			fmt.Println("玩家", player, "手牌", _room.Players[player].HandTiles, "花牌: ", _room.Players[player].FlowerTiles)
		}
	}

	for idx := range _room.Players {
		_room.Players[idx].UpdateUnOpenPool(_room.Players[idx].HandTiles)
		_room.Players[idx].UpdateUnOpenPool(_room.Players[idx].FlowerTiles)
	}
	//fmt.Println("牌海: ", _room.TilesSea)
}

// 玩家抽牌
func (_room *Room) DrawTiles() (ok bool, isFlower bool) {
	ok = true
	isFlower = false
	// 玩家順序進行輪換
	_room.PlayerNow = (_room.PlayerNow + 1) % 4

	// 根據玩家抽牌
	ok, _drawtile := DrawTool.DrawTile(_room.PlayerNow, &_room.TilesSea)

	if !ok || _room.TilesSea == nil {
		fmt.Println("牌海抽完")
		_room.Gameover = true

		// 寫入log
		_room.WriteCsv()

		return false, false
	}

	// 為花牌時，執行補花後重新抽牌
	// 重新抽牌交由外部API進行再一次call DrawTiles處理，避免又抽到花牌
	if _room.GameRule.IsbuHua && _room.IsFlowerTile(_drawtile) {
		_room.Players[_room.PlayerNow].FlowerTiles = append(_room.Players[_room.PlayerNow].FlowerTiles, _drawtile)
		fmt.Println("玩家", _room.PlayerNow, "抽到花牌: ", _drawtile)
		isFlower = true
		return ok, isFlower
	}

	// 抽牌後加入玩家手牌(暫存)
	_room.Players[_room.PlayerNow].OnDrawTile = _drawtile
	_room.Players[_room.PlayerNow].UpdateUnOpenPool([]int{_drawtile})
	_room.DoAction.Action = "Draw"
	_room.DoAction.Tiles = nil
	_room.DoAction.Tiles = append(_room.DoAction.Tiles, _drawtile)
	//fmt.Println("抽牌玩家: ", _room.PlayerNow, " ,抽牌: ", _room.Players[_room.PlayerNow].OnDrawTile)
	//fmt.Println("抽牌玩家手牌: ", _room.Players[_room.PlayerNow].HandTiles)

	//暗槓判斷
	_room.Players[_room.PlayerNow].CanConcealedKong, _room.Players[_room.PlayerNow].CanConcealedKongSet = CheckConcealedKong(_room.Players[_room.PlayerNow].HandTiles, _room.Players[_room.PlayerNow].OnDrawTile)

	//加槓判斷
	_room.Players[_room.PlayerNow].CanAddKong, _room.Players[_room.PlayerNow].CanAddKongSet = CheckAddKong(_room.Players[_room.PlayerNow].CsvPongTiles, _room.Players[_room.PlayerNow].OnDrawTile)

	//AI自摸判斷
	if _room.PlayerNow != _room.RealPlayer {
		_room.Players[_room.PlayerNow].AISelfHuKong(_room.GameRule, _drawtile)

		if _room.Players[_room.PlayerNow].CanSelfHu {
			_room.Players[_room.PlayerNow].IsZimo = 1
			_room.Players[_room.PlayerNow].CanSelfHuTile = _drawtile
			_room.Players[_room.PlayerNow].ResetSelfHuTile = _drawtile

			//fmt.Println("AIPlayer CanSelfHu: ", _room.PlayerNow)
		}

	} else {
		//真人玩家自摸判斷
		//fmt.Println("drawTile: ", int32(_drawtile))
		//fmt.Println("PlayerHand:", _room.Players[_room.PlayerNow].HandTiles)

		_room.Players[_room.PlayerNow].CanSelfHu, ok = _room.Players[_room.PlayerNow].CheckHuFaanJudge(int32(_drawtile), int32(_room.Players[_room.RealPlayer].IsZimo), _room.GameRule)
		fmt.Println("CanSelfHu: ", _room.Players[_room.PlayerNow].CanSelfHu)
		if _room.Players[_room.PlayerNow].CanSelfHu {
			_room.Players[_room.PlayerNow].IsZimo = 1
			_room.Players[_room.PlayerNow].CanSelfHuTile = _drawtile
			_room.Players[_room.PlayerNow].ResetSelfHuTile = _drawtile
			fmt.Println("RealPlayer CanSelfHuTiles: ", _room.Players[_room.PlayerNow].CanSelfHuTile)
			fmt.Println("RealPlayer CanSelfHu: ", _room.Players[_room.PlayerNow].CanSelfHu)
		}
	}

	_room.SetCsvData(_room.PlayerNow)
	return ok, isFlower
}

func (_room *Room) DoChangeThreeTiles(_player int, changetiles [3]int) (ok bool) {
	var changeTileSea []int

	ok = true

	// 玩家換三張(選擇階段)
	changeTileSea = append(changeTileSea, changetiles[0:3]...)

	_room.Players[_player].HandTiles = deleteElements(_room.Players[_player].HandTiles, changetiles[0:3])

	// 其他玩家(AI)換三張(選擇階段)
	for player := 0; player < 4; player++ {
		startIdx := ((player-1)*3 + 3) % 12
		if player != _player {
			//fmt.Println("玩家", player, "換三張: ", _room.Players[player].HandTiles)
			_room.Players[player].AIChangeThreeTiles(_room.GameRule, &changeTileSea)
			//fmt.Println(changeTileSea)
			_room.Players[player].HandTiles = deleteElements(_room.Players[player].HandTiles, changeTileSea[startIdx:startIdx+3])
		}
	}

	// 玩家換三張(交換階段)
	for i := 0; i < 4; i++ {
		startIdx := (i*3 + 3) % 12
		_room.Players[i].ChangeThreeTiles(changeTileSea[startIdx : startIdx+3])
		//fmt.Println("After_ChangeTiles: ", _room.Players[i].HandTiles)
	}

	//換三張內容放到DoAction
	_room.DoAction.Action = "ChangeThreeTiles"
	_room.DoAction.Player = _player
	startIdx := (_room.RealPlayer*3 + 3) % 12
	_room.DoAction.Tiles = changeTileSea[startIdx : startIdx+3]

	// log
	for i := 0; i < 4; i++ {
		_room.DoAction.Player = i
		startIdx = (i*3 + 3) % 12
		_room.DoAction.Tiles = changeTileSea[startIdx : startIdx+3]

		_room.SetCsvData(i)
	}

	_room.DoAction.Player = _player
	startIdx = (_room.RealPlayer*3 + 3) % 12
	_room.DoAction.Tiles = changeTileSea[startIdx : startIdx+3]

	return ok
}

func (_room *Room) DoWithoutTile(_player int, withoutTileType int) (ok bool) {
	ok = true

	// 指定對象做定缺
	//fmt.Println("定缺玩家: ", _player, " ,定缺: ", withoutTileType)
	_room.Players[_player].ChooseWithout(withoutTileType)
	//fmt.Println("定缺玩家: ", _player, " ,定缺: ", _room.Players[_player].WithoutTile)

	// 其他玩家(AI)定缺
	for player := 0; player < 4; player++ {
		if player != _player {
			_room.Players[player].AIChooseWithout(_room.GameRule)
		}
	}

	//選定缺內容放到DoAction
	_room.DoAction.Action = "WithoutTile"
	_room.DoAction.Player = _player
	_room.DoAction.Tiles = []int{_room.Players[_player].WithoutTile}

	for i := 0; i < 4; i++ {
		//fmt.Println(_room.Players[i].WithoutTile)
		_room.DoAction.Player = i
		_room.DoAction.Tiles = []int{_room.Players[i].WithoutTile}

		_room.SetCsvData(i)
	}
	_room.DoAction.Player = _player
	_room.DoAction.Tiles = []int{_room.Players[_player].WithoutTile}
	//fmt.Println("真人玩家的定缺:", _room.Players[_player].WithoutTile)
	return ok
}

func (_room *Room) DoDiscardTile(_player int, _discardtile int) (ok bool) {
	// 接收玩家進行丟牌的動作，檢查後
	ok = _room.Players[_player].PlayerDiscardTile(_discardtile)
	if !ok {
		fmt.Println("玩家", _player, "丟牌失敗")
		return ok
	}
	_room.DoAction.Action = "Discard"
	_room.DoAction.Tiles = nil
	_room.DoAction.Tiles = append(_room.DoAction.Tiles, _discardtile)
	// 把被丟出的牌放入桌上以供後續動作檢查
	_room.OntableTile = _discardtile

	// 回合遞增
	_room.RoundCount += 1

	// 檢查桌上這張牌對所有玩家的觸發動作後回傳前端
	ok = _room.CheckAllplayerAction(_player)

	for idx := range _room.Players {
		if idx != _player {
			_room.Players[idx].UpdateUnOpenPool([]int{_discardtile})
		}
	}

	// 整牌
	sort.Slice(_room.Players[_player].HandTiles, func(i, j int) bool {
		return _room.Players[_player].HandTiles[i] < _room.Players[_player].HandTiles[j]
	})

	_room.SetCsvData(_room.PlayerNow)

	// fmt.Println("玩家", _player, "丟牌: ", _discardtile)
	// fmt.Println("桌上牌: ", _room.OntableTile)
	// fmt.Println("玩家", _player, "手牌: ", _room.Players[_player].HandTiles)
	return ok
}

func (_room *Room) DoAIDiscardTile(_player int) (ok bool) {

	// 確認server要執行丟牌的玩家是不是AI

	// 丟牌前確認是否為自摸後丟牌
	if _room.Players[_player].AfterSelfhu {
		_room.Players[_player].AfterSelfhu = false

		// 執行自摸後丟出剛摸進的牌
		ok = _room.DoDiscardTile(_player, _room.Players[_room.PlayerNow].OnDrawTile)

		// } else if _room.Players[_player].AfterHu {
		// 	// AI胡牌後不須丟牌, 保持桌上牌不變過給下一家即可
		// 	_room.Players[_player].AfterHu = false
		// 	return true

	} else {

		// AI判斷要丟什麼牌
		_discardtile := _room.Players[_player].AIThrow(_room.GameRule, _room.Players[_room.PlayerNow].OnDrawTile)

		// 執行丟牌
		ok = _room.DoDiscardTile(_player, _discardtile)
	}
	return ok
}

func (_room *Room) OpenGame() (ok bool) {
	// 開局
	ok = true
	isFlower := true
	_room.PlayerNow = 3 // 開局由玩家(3 + 1) % 4 = 0開始

	_room.OntableTile = -1

	// 抽出如果為花牌則繼續抽牌
	for ok && isFlower {
		ok, isFlower = _room.DrawTiles()

		// 為花牌的話針對PlayerNow回朔
		if isFlower {
			_room.PlayerNow = 3
		}
	}

	//fmt.Println("開局抽牌: ", ok)

	return ok
}

func (_room *Room) CheckAllplayerAction(_player int) (ok bool) {
	ok = true

	// 檢查桌上的牌對所有玩家的觸發動作
	tile := _room.OntableTile
	//fmt.Println("_room.OntableTile: ", _room.OntableTile)
	// 並把所有玩家的動作結果賦值進玩家的CandoAction裡(通知真人玩家可以做&AI可以做的動作
	for i := 0; i < 4; i++ {
		if i != _room.PlayerNow {
			_room.Players[i].CanPon, _room.Players[i].CanPonSet = CheckPong(_room.Players[i].HandTiles, tile)
			_room.Players[i].CanKong, _room.Players[i].CanKongSet = CheckKong(_room.Players[i].HandTiles, tile)
			// fmt.Println("玩家", i, "CanPon: ", _room.Players[i].CanPon)
			// fmt.Println("玩家", i, "CanKong: ", _room.Players[i].CanKong)
		}
	}

	// 吃的判定只有丟牌的下家需要做
	nextplayerindex := _player + 1
	if nextplayerindex > 3 {
		nextplayerindex = 0
	}
	_room.Players[nextplayerindex].CanChow, _room.Players[nextplayerindex].CanChowSet = CheckChow(_room.Players[nextplayerindex].HandTiles, tile)
	if !_room.GameRule.CanEat {
		_room.Players[nextplayerindex].CanChow = false
	}

	// 檢查有人可不可以胡牌
	// 對AI 可以胡牌(Can && Do 皆為 true), 對真人玩家 只有Can為true
	for i := 0; i < 4; i++ {
		if i != _room.PlayerNow {
			_room.Players[i].CheckHuFaanJudge(int32(tile), _room.Players[i].IsZimo, _room.GameRule)

			// 非真人玩家
			if _room.Players[i].CanHu && (i != _room.RealPlayer) {
				_room.Players[i].DoHu = true
				fmt.Println("玩家", i, "CanHu: ", _room.Players[i].CanHu)
				fmt.Println("玩家", i, "DoHu: ", _room.Players[i].DoHu)
			}
		}

		//Check BaoJi
		_room.Players[i].CanBaoJi = CheckBaoJi(_room.Players[i].CanHu, _room.Players[i].IsBaoJi)
		_room.Players[i].CanMingPai = CheckMingPai(_room.Players[i].IsAlreadyHu, i, _room.PlayerNow)
	}

	// 排除當下丟牌的玩家，此玩家不能做任何動作
	_room.Players[_player].Donothing = true

	// 檢查其他玩家的動作，如果大家都沒有做動作，輪換下一家進行摸牌
	_room.NothingDo = true
	for i := 0; i < 4; i++ {
		if _room.Players[i].DoChow {
			_room.NothingDo = false
		}
		if _room.Players[i].DoPon {
			_room.NothingDo = false
		}
		if _room.Players[i].DoKong {
			_room.NothingDo = false
		}
		if _room.Players[i].DoConcealedKong {
			_room.NothingDo = false
		}
		if _room.Players[i].DoHu {
			_room.NothingDo = false
		}

	}

	return ok
}

func (_room *Room) CheckActionPriority() (ok bool) {
	ok = true
	// 檢查所有玩家的動作
	_room.NothingDo = true

	for i := 0; i < 4; i++ {
		if !_room.GameRule.CanEat {
			_room.Players[i].DoChow = false
		}
		if !_room.GameRule.CanPong {
			_room.Players[i].DoPon = false
		}
		if !_room.GameRule.CanKong {
			_room.Players[i].DoKong = false
		}
	}

	// 並依照優先順序把低優先順序的動作清除 (ex: player.DoPong = false)
	for i := 0; i < 4; i++ {
		if _room.Players[i].DoHu {
			_room.Players[(i+1)%4].DoKong = false
			_room.Players[(i+2)%4].DoKong = false
			_room.Players[(i+3)%4].DoKong = false

			_room.Players[(i+1)%4].DoConcealedKong = false
			_room.Players[(i+2)%4].DoConcealedKong = false
			_room.Players[(i+3)%4].DoConcealedKong = false

			_room.Players[(i+1)%4].DoAddKong = false
			_room.Players[(i+2)%4].DoAddKong = false
			_room.Players[(i+3)%4].DoAddKong = false

			_room.Players[(i+1)%4].DoPon = false
			_room.Players[(i+2)%4].DoPon = false
			_room.Players[(i+3)%4].DoPon = false

			_room.Players[(i+1)%4].DoChow = false
			_room.Players[(i+2)%4].DoChow = false
			_room.Players[(i+3)%4].DoChow = false
		}

		if _room.Players[i].DoSelfHu {
			_room.Players[(i+1)%4].DoKong = false
			_room.Players[(i+2)%4].DoKong = false
			_room.Players[(i+3)%4].DoKong = false

			_room.Players[(i+1)%4].DoConcealedKong = false
			_room.Players[(i+2)%4].DoConcealedKong = false
			_room.Players[(i+3)%4].DoConcealedKong = false

			_room.Players[(i+1)%4].DoAddKong = false
			_room.Players[(i+2)%4].DoAddKong = false
			_room.Players[(i+3)%4].DoAddKong = false

			_room.Players[(i+1)%4].DoPon = false
			_room.Players[(i+2)%4].DoPon = false
			_room.Players[(i+3)%4].DoPon = false

			_room.Players[(i+1)%4].DoChow = false
			_room.Players[(i+2)%4].DoChow = false
			_room.Players[(i+3)%4].DoChow = false
		}

		if _room.Players[i].DoKong || _room.Players[i].DoHu || _room.Players[i].DoConcealedKong || _room.Players[i].DoAddKong || _room.Players[i].DoSelfHu { //碰
			_room.Players[(i+1)%4].DoPon = false
			_room.Players[(i+2)%4].DoPon = false
			_room.Players[(i+3)%4].DoPon = false
		}

		if _room.Players[i].DoPon || _room.Players[i].DoKong || _room.Players[i].DoConcealedKong || _room.Players[i].DoAddKong || _room.Players[i].DoHu ||
			_room.Players[i].DoSelfHu { //吃
			_room.Players[(i+1)%4].DoChow = false
			_room.Players[(i+2)%4].DoChow = false
			_room.Players[(i+3)%4].DoChow = false
		}

	}

	// 只留下要被執行的動作	(ex: player.dopong = true)
	for i := 0; i < 4; i++ {
		if !_room.Players[i].DoChow && !_room.Players[i].DoPon &&
			!_room.Players[i].DoKong && !_room.Players[i].DoHu && !_room.Players[i].DoConcealedKong && !_room.Players[i].DoAddKong && !_room.Players[i].DoSelfHu {
			_room.Players[i].Donothing = true
		}
	}

	for i := 0; i < 4; i++ {
		if !_room.Players[i].Donothing {
			_room.NothingDo = false
		}
	}
	return ok
}

// 真人玩家
func (_room *Room) DoChowPongKong(_action string, _tiles []int) (ok bool) {
	ok = true
	// Verify Action
	// 待補

	_room.DoAction.PriviousPlayer = _room.PlayerNow
	//"Nothing", "Chow", "Pong", "Exposed Kong", "Concealed Kong", "Hu", "SelfDrawn"
	if _action == "Chow" {
		_room.Players[_room.RealPlayer].DoChow = true
		_room.DoAction.Action = "Chow"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles

	} else if _action == "Pong" {
		_room.Players[_room.RealPlayer].DoPon = true
		_room.DoAction.Action = "Pong"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles
	} else if _action == "ExposedKong" {
		_room.Players[_room.RealPlayer].DoKong = true
		_room.DoAction.Action = "ExposedKong"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles
	} else if _action == "ConcealedKong" {
		_room.Players[_room.RealPlayer].DoConcealedKong = true
		_room.Players[_room.RealPlayer].DoKongTile = _tiles[0]
		_room.DoAction.Action = "ConcealedKong"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles

	} else if _action == "AddKong" {
		_room.Players[_room.RealPlayer].DoAddKong = true
		_room.DoAction.Action = "AddKong"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles
	} else if _action == "Hu" {
		_room.Players[_room.RealPlayer].DoHu = true
		_room.DoAction.Action = "Hu"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles
	} else if _action == "SelfHu" {
		_room.Players[_room.RealPlayer].DoSelfHu = true
		_room.DoAction.Action = "SelfHu"
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles
		_room.Players[_room.RealPlayer].CanSelfHuTile = _room.Players[_room.RealPlayer].ResetSelfHuTile
		// 真人玩家確定自摸的情況下，清空桌面牌避免檢查其他玩家動作
		_room.OntableTile = -1
	} else if _action == "Nothing" {
		_room.Players[_room.RealPlayer].Donothing = true
		_room.DoAction.Action = ""
		_room.DoAction.Player = _room.RealPlayer
		_room.DoAction.Tiles = _tiles
	} else {
		ok = false
	}

	return ok
}

func (_room *Room) AIDoChowPongKong() (ok bool, tiles []int) {
	ok = true

	for idx, player := range _room.Players {
		if idx != _room.RealPlayer && idx != _room.PlayerNow {
			//fmt.Println("_room.OntableTile:", _room.OntableTile)

			// 玩家胡後，桌面牌清空不須對此桌牌做動作判斷
			if _room.OntableTile == -1 {
				break
			}
			ok = player.AIHuChowPongKong(_room.GameRule, _room.OntableTile)

			// 印出AI的動作
			if player.DoChow && ok {
				//fmt.Println("玩家", idx, "吃")
				_room.Players[idx].DoChow = true
				_room.DoAction.Action = "Chow"
				_room.DoAction.Player = idx
				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.OntableTile)
			} else if player.DoPon && ok {
				//fmt.Println("玩家", idx, "碰")
				_room.Players[idx].DoPon = true
				_room.DoAction.Action = "Pong"
				_room.DoAction.Player = idx
				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.OntableTile)
			} else if player.DoKong && ok {
				//fmt.Println("玩家", idx, "槓")
				_room.Players[idx].DoKong = true
				_room.DoAction.Action = "ExposedKong"
				_room.DoAction.Player = idx
				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.OntableTile)
			} else if player.DoHu && ok {
				//fmt.Println("玩家", idx, "胡")
				_room.Players[idx].DoHu = true
				_room.DoAction.Action = "Hu"
				_room.DoAction.Player = idx
				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.OntableTile)
				_room.Players[idx].CanHuTile = _room.OntableTile
				// _, ok = _room.Players[idx].CheckHuFaanJudge(int32(_room.OntableTile), _room.Players[_room.PlayerNow].IsZimo, _room.GameRule)
				// fmt.Println("玩家", idx, "胡牌番型: ", _room.Players[idx].CanHuFaanListStr)
			}
		}
	}
	tiles = _room.DoAction.Tiles
	return ok, tiles
}

// 執行動作
func (_room *Room) AllPlayerDoAction(tiles []int) (ok bool) {
	ok = true
	hutabletile := false

	// 針對一炮多響的狀況
	BeHuPlayer := -1

	for i := 0; i < 4; i++ {
		if _room.Players[i].DoSelfHu {
			// 自摸
			//fmt.Println("玩家", i, "自摸")
			_, ok = _room.Players[i].CheckHuFaanJudge(int32(_room.Players[i].CanSelfHuTile), _room.Players[i].IsZimo, _room.GameRule)
			//fmt.Println("玩家", i, "胡牌番型: ", _room.Players[_room.PlayerNow].CanHuFaanListStr)

			_room.DoAction.Action = "SelfHu"
			_room.DoAction.Player = i
			_room.DoAction.Tiles = []int{_room.Players[i].CanSelfHuTile}

			_room.Players[i].AfterSelfhu = true
			_room.Players[i].IsAlreadyHu = true

			_room.SelfHuSettlement(i)
			_room.SetCsvData(i)

			if i == _room.RealPlayer {
				_room.Players[i].AfterSelfhu = false
				_room.Players[i].DiscardTiles = append(_room.Players[i].DiscardTiles, _room.Players[i].ResetSelfHuTile)

				// // 存取 真人丟牌log, 並清空紀錄番型
				// _room.Players[i].CanHuFaanListStr = nil
				// _room.DoAction.Action = "Discard"
				// _room.SetCsvData(_room.PlayerNow)
			}

		} else if _room.Players[i].DoHu {
			// 胡牌
			//fmt.Println("玩家: ", i, " 胡牌")
			//fmt.Println("玩家: ", _room.PlayerNow, "被胡牌")

			// 針對一炮多響的狀況 紀錄被胡的玩家idx
			if BeHuPlayer == -1 {
				BeHuPlayer = _room.PlayerNow
			}
			_room.PlayerNow = i

			_, ok = _room.Players[i].CheckHuFaanJudge(int32(_room.OntableTile), _room.Players[i].IsZimo, _room.GameRule)
			//fmt.Println("ok:", ok, "玩家", i, "胡牌番型: ", _room.Players[_room.PlayerNow].CanHuFaanListStr)

			_room.Players[i].AfterHu = true

			_room.DoAction.Action = "Hu"
			_room.DoAction.Player = i
			// 把桌上的牌塞進去
			_room.DoAction.Tiles = []int{_room.OntableTile}
			hutabletile = true
			_room.Players[i].IsAlreadyHu = true
			_room.HuSettlement(i, BeHuPlayer)

			_room.SetCsvData(i)

		} else if _room.Players[i].DoKong {
			//fmt.Println("執行槓")
			_room.Players[i].Kong(tiles[0])
			_room.Players[i].CsvKongTIles = append(_room.Players[i].CsvKongTIles, tiles[0])

			elements := []int{_room.OntableTile}
			_room.Players[_room.PlayerNow].DiscardTiles = deleteElements(_room.Players[_room.PlayerNow].DiscardTiles, elements)

			_room.DoAction.Action = "ExposedKong"
			_room.DoAction.Player = i
			_room.DoAction.Tiles = nil
			_room.DoAction.Tiles = append(_room.DoAction.Tiles, tiles[0])
			_room.SetCsvData(i)

			//槓牌完該玩家要摸進一張牌，index-1，DrawTiles才會是執行動作的玩家
			_room.PlayerNow = (i - 1) % 4
			_room.DrawTiles()
			//drawTile := _room.Players[_room.PlayerNow].OnDrawTile
			//fmt.Println("drawTile: ", drawTile)
			//fmt.Println("槓後抽到的牌: ", _room.Players[_room.PlayerNow].OnDrawTile)

		} else if _room.Players[i].DoConcealedKong {
			fmt.Println("執行暗槓: ", _room.Players[i].DoKongTile)
			_room.Players[i].ConcealedKong(_room.Players[i].DoKongTile)
			_room.Players[i].CsvConcealedKongTiles = append(_room.Players[i].CsvConcealedKongTiles, _room.Players[i].DoKongTile)

			_room.DoAction.Action = "ConcealedKong"
			_room.DoAction.Player = i
			_room.DoAction.Tiles = nil
			_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.Players[i].DoKongTile)
			_room.SetCsvData(i)

			_room.Players[i].CanConcealedKong = false
			_room.Players[i].CanConcealedKongSet = nil
			_room.Players[i].DoKongTile = -1

			// elements := []int{_room.OntableTile}
			// _room.Players[_room.PlayerNow].DiscardTiles = deleteElements(_room.Players[_room.PlayerNow].DiscardTiles, elements)

			_room.PlayerNow = (i - 1) % 4
			_room.DrawTiles()

			// fmt.Println("Player: ", _room.PlayerNow, " | ", _room.RealPlayer)
			// if _room.PlayerNow != _room.RealPlayer {
			// 	ok = _room.DoAIDiscardTile(_room.PlayerNow)
			// 	//fmt.Println("AI槓後丟牌: ", ok)
			// }

		} else if _room.Players[i].DoAddKong {
			//fmt.Println("執行加槓")

			// remove MeldTiles
			temp := make([]TransformGolang.Meld, 0)
			for _, meld := range _room.Players[i].MeldTiles {
				if meld.Tiles[0] != _room.DoAction.Tiles[0] {
					temp = append(temp, meld)
				}
			}

			//if Do AddKong need to remove PongTiles
			for i := 0; i < 3; i++ {
				_room.Players[i].CsvPongTiles = deleteElements(_room.Players[i].CsvPongTiles, _room.DoAction.Tiles)
			}
			_room.Players[i].MeldTiles = temp

			_room.Players[i].OnDrawTile = -1
			_room.PlayerNow = (i - 1) % 4
			_room.DrawTiles()
			_room.Players[i].CsvAddKongTiles = append(_room.Players[i].CsvAddKongTiles, tiles...)

			_room.DoAction.Action = "AddKong"
			_room.DoAction.Player = _room.PlayerNow
			_room.DoAction.Tiles = append(_room.DoAction.Tiles, tiles...)

			_room.SetCsvData(_room.PlayerNow)

		} else if _room.Players[i].DoPon {
			fmt.Println("執行碰")
			_room.Players[i].Pong(tiles[0])
			_room.Players[i].CsvPongTiles = append(_room.Players[i].CsvPongTiles, tiles...)

			//Delete Tile Discard from other Player
			elements := []int{_room.OntableTile}
			_room.Players[_room.PlayerNow].DiscardTiles = deleteElements(_room.Players[_room.PlayerNow].DiscardTiles, elements)
			_room.PlayerNow = i

			_room.DoAction.Action = "Pong"
			_room.DoAction.Player = _room.PlayerNow
			_room.DoAction.Tiles = _room.Players[i].CsvPongTiles
			fmt.Println("_room.DoAction.Tiles: ", _room.DoAction.Tiles)

			// GameLog儲存資訊
			/*csvTile := arrayToString(_room.DoAction.Tiles)
			csvTileSea := arrayToString(_room.TilesSea)
			csvHandTiles := arrayToString(_room.Players[_room.PlayerNow].HandTiles)
			csvDiscardTiles := arrayToString(_room.Players[_room.PlayerNow].DiscardTiles)
			csvAlreadyChow := arrayToString(_room.Players[i].CsvChowTiles)
			csvAlreadyPong := arrayToString(_room.Players[i].CsvPongTiles)
			csvAlreadyKong := arrayToString(_room.Players[i].CsvKongTIles)
			csvConcealedKong := arrayToString(_room.Players[i].CsvConcealedKongTiles)
			csvPoint := " "

			row := []string{strconv.Itoa(_room.RoundCount), strconv.Itoa(_room.PlayerNow), strconv.Itoa(_room.PrevailingWind), _room.DoAction.Action,
				csvTile, csvTileSea, csvHandTiles, csvDiscardTiles, csvPoint, csvAlreadyChow, csvAlreadyPong, csvAlreadyKong, csvConcealedKong}
			writeCsv(row)*/
			_room.SetCsvData(_room.PlayerNow)

		} else if _room.Players[i].DoChow {
			fmt.Println("執行吃")
			_room.Players[i].Chow(tiles, _room.OntableTile)
			_room.Players[i].CsvChowTiles = append(_room.Players[i].CsvChowTiles, tiles...)
			_room.PlayerNow = i

			_room.DoAction.Action = "Chow"
			_room.DoAction.Player = _room.PlayerNow
			//_room.DoAction.Tiles = append(_room.DoAction.Tiles, tiles...)
			_room.DoAction.Tiles = tiles

			/*csvTile := arrayToString(_room.DoAction.Tiles)
			csvTileSea := arrayToString(_room.TilesSea)
			csvHandTiles := arrayToString(_room.Players[_room.PlayerNow].HandTiles)
			csvDiscardTiles := arrayToString(_room.Players[_room.PlayerNow].DiscardTiles)
			csvAlreadyChow := arrayToString(_room.Players[i].CsvChowTiles)
			csvAlreadyPong := arrayToString(_room.Players[i].CsvPongTiles)
			csvAlreadyKong := arrayToString(_room.Players[i].CsvKongTIles)
			csvConcealedKong := arrayToString(_room.Players[i].CsvConcealedKongTiles)
			csvPoint := " "

			row := []string{strconv.Itoa(_room.RoundCount), strconv.Itoa(_room.PlayerNow), strconv.Itoa(_room.PrevailingWind), _room.DoAction.Action,
				csvTile, csvTileSea, csvHandTiles, csvDiscardTiles, csvPoint, csvAlreadyChow, csvAlreadyPong, csvAlreadyKong, csvConcealedKong}
			writeCsv(row)*/
			_room.SetCsvData(_room.PlayerNow)
		} else if _room.Players[i].DoBaoJi {
			_room.Players[i].CanBaoJi = false
			_room.Players[i].IsBaoJi = true

		} else if _room.Players[i].DoMingPai {
			_room.Players[i].CanBaoJi = false
			_room.Players[i].IsMingPai = true

		} else {
			//fmt.Println("玩家", i, "Do_nothing")
		}

		if i != _room.PlayerNow {
			_room.Players[i].UpdateUnOpenPool(tiles)
		}
	}
	if hutabletile {
		// 做完胡牌後，清空ontabletile
		_room.OntableTile = -1
	}
	return ok
}

// 轉成 []byte
func (_room *Room) ToByteArray() []byte {
	jsonBytes, _ := json.Marshal(_room)
	return (jsonBytes)
}

func (_room *Room) GameOver() (ok bool) {
	ok = true
	if _room.Players[_room.PlayerNow].CanHu {
		_room.Gameover = true
		ok = false
	} else if _room.TilesSea == nil {
		_room.Gameover = true
		ok = false
	} else if _room.Players[_room.PlayerNow].CanSelfHu {
		_room.Gameover = true
		ok = false
	} else {
		_room.Gameover = false
	}

	return ok
}

// 這輪結束後做洗牌(但還沒結算)，玩家與牌海洗牌並發給四位玩家牌
func (_room *Room) InitTiles() {

}

// 結算番數or台數，在有人胡牌執行，有可能多人同時胡牌
// 胡牌邏輯: A+1 分, B-1 分，其餘不動
func (_room *Room) HuSettlement(_huPlayer int, _previousPlayer int) (ok bool) {
	ok = true

	//fmt.Println("玩家", _huPlayer, "番型:", _room.Players[_huPlayer].CanHuFaanList)
	faanlist := _room.Players[_huPlayer].CanHuFaanList
	faanScore := int32(1)

	for i := 0; i < len(faanlist); i++ {
		// 番數計算相乘
		faanScore *= _room.GameRule.FanCount[faanlist[i]-1]
		//fmt.Println("玩家", _huPlayer, "番型:", _room.Players[_huPlayer].CanHuFaanListStr[i], "番數:", _room.GameRule.FanCount[faanlist[i]-1])
	}

	_room.Players[_huPlayer].Point += int(faanScore)
	_room.Players[_previousPlayer].Point -= int(faanScore)

	if _room.Players[_huPlayer].IsBaoJi {
		_room.Players[_huPlayer].Point *= 3
		_room.Players[_huPlayer].DoBaoJi = false
		_room.Players[_huPlayer].IsBaoJi = false
	}

	// 空番型的情況，回傳前端傳胡牌字串
	if len(faanlist) == 0 {
		_room.Players[_huPlayer].CanHuFaanListStr = append(_room.Players[_huPlayer].CanHuFaanListStr, "胡牌")
	}

	return ok
}

// 結算番數or台數，在有人自摸執行，不可能多人自摸
func (_room *Room) SelfHuSettlement(_player int) (ok bool) {
	ok = true

	//fmt.Println("玩家", _player, "番型:", _room.Players[_player].CanHuFaanList)
	faanlist := _room.Players[_player].CanHuFaanList
	faanScore := int32(1)

	for i := 0; i < len(faanlist); i++ {
		// 番數計算相乘
		faanScore *= _room.GameRule.FanCount[faanlist[i]-1]
		//fmt.Println("玩家", _player, "番型:", _room.Players[_player].CanHuFaanListStr[i], "番數:", _room.GameRule.FanCount[faanlist[i]-1])
	}

	// 當玩家自摸時，其他玩家要扣掉自摸玩家的番數
	// 自摸邏輯: A自摸的番數*3倍，其他人各扣A自摸的番數
	_room.Players[_player].Point += int(faanScore) * 3
	for i := 0; i < 4; i++ {
		if i != _player {
			_room.Players[i].Point -= int(faanScore)
		}
	}

	if _room.Players[_player].IsBaoJi {
		_room.Players[_player].Point *= 3
		_room.Players[_player].DoBaoJi = false
		_room.Players[_player].IsBaoJi = false
	}

	// 空番型的情況，回傳前端傳自摸字串
	if len(faanlist) == 0 {
		_room.Players[_player].CanHuFaanListStr = append(_room.Players[_player].CanHuFaanListStr, "自摸")
	}

	return ok
}

func (_room *Room) SetCsvData(idx int) {
	csvTile := arrayToString(_room.DoAction.Tiles)
	csvTileSea := arrayToString(_room.TilesSea)
	csvHandTiles := arrayToString(_room.Players[idx].HandTiles)
	csvDiscardTiles := arrayToString(_room.Players[idx].DiscardTiles)
	csvAlreadyChow := arrayToString(_room.Players[idx].CsvChowTiles)
	csvAlreadyPong := arrayToString(_room.Players[idx].CsvPongTiles)
	csvAlreadyKong := arrayToString(_room.Players[idx].CsvKongTIles)
	csvConcealedKong := arrayToString(_room.Players[idx].CsvConcealedKongTiles)
	csvAddKong := arrayToString(_room.Players[idx].CsvAddKongTiles)
	var csvHuFannListStr string
	for id, i := range _room.Players[idx].CanHuFaanListStr {
		if id == len(_room.Players[idx].CanHuFaanListStr) {
			break
		} else {
			csvHuFannListStr += i
			csvHuFannListStr += " "
		}
	}

	var csvPoint string
	for idx, i := range _room.Players {
		if idx == len(_room.Players) {
			break
		} else {
			csvPoint += strconv.Itoa(i.Point)
			csvPoint += " "
		}
	}
	//csvPoint := strconv.Itoa(_room.Players[idx].Point)

	row := []string{strconv.Itoa(int(_room.RoomId)), strconv.Itoa(_room.RoundCount), strconv.Itoa(idx), strconv.Itoa(_room.PrevailingWind), _room.DoAction.Action,
		csvTile, csvHandTiles, csvDiscardTiles, csvPoint, csvAlreadyChow, csvAlreadyPong, csvAlreadyKong, csvConcealedKong, csvAddKong, csvHuFannListStr, csvTileSea}

	_room.GameLog = append(_room.GameLog, row)
	//writeCsv(row)

}

func (_room *Room) WriteCsv() (ok bool) {
	ok = false
	newFileName := "./gamemanager/test.csv" //路徑
	nfs, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
	}
	defer nfs.Close()
	nfs.WriteString("\xEF\xBB\xBF") //寫入UTF-8 BOM防止以Excel開啟導致中文亂碼
	nfs.Seek(0, io.SeekEnd)

	w := csv.NewWriter(nfs)

	w.Comma = ','
	w.UseCRLF = true
	//row = []string{"0", "0", "0", "丟牌[2]", "36", "[0,0,0,1,2,3,11,12,13,15,15,15,25]", "[8,9,6,3]", "99999"}
	for i := 0; i < len(_room.GameLog); i++ {
		row := _room.GameLog[i]
		err = w.Write(row)
		if err != nil {
			log.Fatalf("can not write, err is %+v", err)
			ok = false
			break
		} else {
			ok = true
		}
	}
	w.Flush()
	return ok
}

func (_room *Room) IsFlowerTile(tile int) (isFlower bool) {
	isFlower = false

	// 34~41: 春夏秋冬梅蘭竹菊
	if tile >= 34 && tile <= 41 {
		isFlower = true
	}
	return isFlower
}

// ForAItest
func (_room *Room) SetSeat() {
	for i := 0; i < 4; i++ {
		newIdx := i + rand.Intn(4-i)
		_room.Players[i], _room.Players[newIdx] = _room.Players[newIdx], _room.Players[i]
	}
}

func (_room *Room) AITransAction() {
	//將AI的CanAction轉換為DoAction
	for idx := 0; idx < 4; idx++ {
		if _room.Players[idx].CanHu {
			_room.Players[idx].DoHu = true
		} else if _room.Players[idx].CanSelfHu {
			_room.Players[idx].DoSelfHu = true
		} else if _room.Players[idx].CanKong {
			_room.Players[idx].DoKong = true
		} else if _room.Players[idx].CanPon {
			_room.Players[idx].DoPon = true
		} else if _room.Players[idx].CanConcealedKong {
			_room.Players[idx].DoConcealedKong = true
		} else if _room.Players[idx].CanAddKong {
			_room.Players[idx].DoAddKong = true
		}
	}
}

func (_room *Room) RunGame(GameRule string, round int) { // RunGame()
	ok := _room.InitRoom(GameRule, true, round)

	if ok {
		fmt.Println("StartGame Round:", round)
		//初始發牌

		_room.DealTiles()
	} else {
		fmt.Println("Initial Room Failed")
		return
	}

	// 選定缺
	for idx := 0; idx < 4; idx++ {
		_room.Players[idx].AIChooseWithout(_room.GameRule)
		//fmt.Println("玩家 ", idx, "定缺: ", _room.Players[idx].WithoutTile)

		// 定缺寫進csv
		_room.DoAction.Action = "WithoutTile"
		_room.DoAction.Player = idx
		_room.DoAction.Tiles = []int{_room.Players[idx].WithoutTile}
		_room.SetCsvData(idx)
	}

	// 換三張 (挑牌)
	var ChangeTilesSea []int
	for idx := 0; idx < 4; idx++ {
		startidx := ((idx-1)*3 + 3) % 12
		_room.Players[idx].AIChangeThreeTiles(_room.GameRule, &ChangeTilesSea)
		_room.Players[idx].HandTiles = deleteElements(_room.Players[idx].HandTiles, ChangeTilesSea[startidx:startidx+3])
	}

	// 換三張 (交換)
	for idx := 0; idx < 4; idx++ {
		startidx := (idx*3 + 3) % 12
		_room.Players[idx].ChangeThreeTiles(ChangeTilesSea[startidx : startidx+3])

		// 換三張寫進csv
		_room.DoAction.Action = "ChangeThreeTiles"
		_room.DoAction.Player = idx
		_room.DoAction.Tiles = ChangeTilesSea[startidx : startidx+3]
		_room.SetCsvData(idx)
	}

	//開門
	_room.OpenGame()

	for !_room.Gameover {

		ok = _room.DoAIDiscardTile(_room.PlayerNow)
		// fmt.Println("DoAIDiscardTile: ", ok)
		// fmt.Println("OntableTile: ", _room.OntableTile)

		// 清除玩家所有動作
		for idx := 0; idx < 4; idx++ {
			_room.Players[idx].ResetAction()
		}

		// 確認所有AI想執行的動作
		ok, AItiles := _room.AIDoChowPongKong()

		// 確認動作優先序
		if ok {
			ok = _room.CheckActionPriority()
		}

		if _room.NothingDo {
			// 沒有要執行就回到 下一家摸牌丟牌然後問前端丟回後端的遊戲循環
			// 摸牌需迴圈判定花牌，直到摸到非花牌為止
			isFlower := true
			for ok && isFlower {
				tempPlayerNow := _room.PlayerNow
				ok, isFlower = _room.DrawTiles()

				// 如果抽到花牌要在重抽的情況下，PlayerNow要重整回原先的值
				if isFlower {
					_room.PlayerNow = tempPlayerNow
				}
			}

			if !ok && _room.Gameover {
				// 遊戲結束
				break
			}
			if ok && _room.Players[_room.PlayerNow].DoSelfHu || _room.Players[_room.PlayerNow].DoConcealedKong {
				// AI摸牌後選擇執行自摸
				//fmt.Println("PlayerNow:", _room.PlayerNow, "CanSelfHu", _room.Players[_room.PlayerNow].CanSelfHu, "DoSelfHu", _room.Players[_room.PlayerNow].DoSelfHu)
				_room.AllPlayerDoAction(AItiles)
			}
		} else {
			// 執行動作
			_room.AllPlayerDoAction(AItiles)
		}

		if !_room.NothingDo && _room.DoAction.Action == "Hu" {
			_room.PlayerNow = _room.DoAction.Player

			if _room.GameRule.GameMode != 2 {
				_room.Gameover = true
			}
		}

		// if ok {
		// 	//fmt.Println("_room.NothingDo:", _room.NothingDo)
		// 	if _room.NothingDo {
		// 		// 無玩家動作須執行, 下一家摸進牌後判斷此AI摸牌後動作 (自摸 暗槓 加槓), 確定無動作(不須再摸牌)後跳出迴圈執行丟牌
		// 		for ok && !_room.NothingDo {

		// 			ok, _ = _room.DrawTiles()

		// 			if ok && _room.Players[_room.PlayerNow].DoSelfHu {
		// 				fmt.Println("玩家", _room.PlayerNow, "自摸")

		// 				_room.DoAction.Action = "SelfHu"
		// 				_room.DoAction.Player = _room.PlayerNow
		// 				var huTile []int
		// 				huTile = append(huTile, _room.Players[_room.PlayerNow].CanSelfHuTile)
		// 				_room.DoAction.Tiles = huTile

		// 				_, ok = _room.Players[_room.PlayerNow].CheckHuFaanJudge(int32(_room.Players[_room.PlayerNow].CanSelfHuTile), _room.Players[_room.PlayerNow].IsZimo, _room.GameRule)
		// 				fmt.Println("玩家", _room.PlayerNow, "胡牌番型: ", _room.Players[_room.PlayerNow].CanHuFaanListStr)
		// 				_room.SetCsvData(_room.PlayerNow)
		// 			} else if ok && _room.Players[_room.PlayerNow].DoConcealedKong {
		// 				_room.Players[_room.PlayerNow].ConcealedKong(_room.Players[_room.PlayerNow].OnDrawTile)
		// 				fmt.Println("執行暗槓")

		// 				_room.DoAction.Action = "ConcealedKong"
		// 				_room.DoAction.Player = _room.PlayerNow
		// 				_room.DoAction.Tiles = nil
		// 				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.Players[_room.PlayerNow].OnDrawTile)
		// 				_room.PlayerNow = (_room.PlayerNow - 1) % 4
		// 				_room.SetCsvData(_room.PlayerNow)

		// 				_room.NothingDo = false
		// 			} else if ok && _room.Players[_room.PlayerNow].DoAddKong {
		// 				_room.Players[_room.PlayerNow].AddKong(_room.Players[_room.PlayerNow].OnDrawTile)
		// 				fmt.Println("執行加槓")

		// 				_room.DoAction.Action = "AddKong"
		// 				_room.DoAction.Player = _room.PlayerNow
		// 				_room.DoAction.Tiles = nil
		// 				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.Players[_room.PlayerNow].OnDrawTile)
		// 				_room.PlayerNow = (_room.PlayerNow - 1) % 4
		// 				_room.SetCsvData(_room.PlayerNow)

		// 				_room.NothingDo = false
		// 			}
		// 		}
		// 	} else {
		// 		// 執行AI選定並通過優先序列的動作
		// 		for idx := 0; idx < 4; idx++ {
		// 			if _room.Players[idx].DoHu {
		// 				fmt.Println("胡牌")

		// 				_room.DoAction.Action = "Hu"
		// 				_room.DoAction.Player = idx
		// 				var huTile []int
		// 				huTile = append(huTile, _room.OntableTile)
		// 				_room.DoAction.Tiles = huTile

		// 				_room.PlayerNow = idx
		// 				//_room.DoAIDiscardTile(_room.PlayerNow)
		// 				_, ok = _room.Players[idx].CheckHuFaanJudge(int32(_room.OntableTile), _room.Players[_room.PlayerNow].IsZimo, _room.GameRule)
		// 				fmt.Println("ok:", ok, "玩家", idx, "胡牌番型: ", _room.Players[_room.PlayerNow].CanHuFaanListStr)
		// 				ok = _room.Players[idx].AIHuChowPongKong(_room.GameRule, _room.OntableTile)
		// 				fmt.Println("AIHuChowPongKong: ", ok)
		// 				fmt.Println("ok:", ok, "玩家", idx, "胡牌番型: ", _room.Players[idx].CanHuFaanListStr)
		// 				_room.SetCsvData(idx)
		// 				_room.DrawTiles()

		// 			} else if _room.Players[idx].DoKong {
		// 				fmt.Println("槓牌")

		// 				_room.DoAction.Action = "Kong"
		// 				_room.DoAction.Player = idx
		// 				_room.DoAction.Tiles = append(_room.DoAction.Tiles, _room.OntableTile)

		// 				_room.Players[_room.PlayerNow].Kong(_room.OntableTile)
		// 				_room.SetCsvData(idx)

		// 				// 槓牌後新摸一張進牌
		// 				_room.PlayerNow = (_room.PlayerNow - 1) % 4
		// 				_room.DrawTiles()

		// 			} else if _room.Players[idx].DoPon {
		// 				_room.Players[_room.PlayerNow].Pong(_room.OntableTile)
		// 				_room.SetCsvData(idx)
		// 			} else if _room.Players[idx].DoChow {
		// 				//_room.Players[_room.PlayerNow].Chow()
		// 				_room.SetCsvData(idx)

		// 			}
		// 		}
		// 	}
		// }

	}

}
