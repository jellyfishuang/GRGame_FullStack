package drawtool

// 目前皆以隨機抽取為主，後續要開發手牌工具就可以用這個接口進行修改(指定玩家發什麼牌)

import (
	GameRule "GolangGameManager/gamerule"
	"math/rand"
	"sort"
	"time"
)

// 初始化牌海，用於 _room.TilesSea 與 _player.UnOpenPool
func InitTileSea[T int | int32](_tilesSea *[]T, _gamerule GameRule.GameRule, _shuffle bool) {

	*_tilesSea = make([]T, 0, _gamerule.AvailableTileCount)

	for tile := 0; tile < len(_gamerule.AvailableTiles); tile++ {
		for cardcount := int32(0); cardcount < _gamerule.AvailableTiles[tile]; cardcount++ {
			*_tilesSea = append(*_tilesSea, T(tile))
		}
	}

	// 洗牌海
	if _shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(*_tilesSea), func(i, j int) { (*_tilesSea)[i], (*_tilesSea)[j] = (*_tilesSea)[j], (*_tilesSea)[i] })
	}

}

// 開局(隨機)發牌, 根據遊戲規則決定要發給四個玩家 13 or (16張牌)
func DealTile(_player int, _tilesSea *[]int, _handTilesCounts int) (_player1hand []int, _player2hand []int, _player3hand []int, _player4hand []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < _handTilesCounts; i++ {
		_, tile := DrawTile(0, _tilesSea)
		_player1hand = append(_player1hand, tile)
		_, tile = DrawTile(1, _tilesSea)
		_player2hand = append(_player2hand, tile)
		_, tile = DrawTile(2, _tilesSea)
		_player3hand = append(_player3hand, tile)
		_, tile = DrawTile(3, _tilesSea)
		_player4hand = append(_player4hand, tile)
	}
	sort.Slice(_player1hand, func(i, j int) bool {
		return _player1hand[i] < _player1hand[j]
	})
	sort.Slice(_player2hand, func(i, j int) bool {
		return _player2hand[i] < _player2hand[j]
	})
	sort.Slice(_player3hand, func(i, j int) bool {
		return _player3hand[i] < _player3hand[j]
	})
	sort.Slice(_player4hand, func(i, j int) bool {
		return _player4hand[i] < _player4hand[j]
	})

	// _player1hand = []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 4, 4, 4, 6}
	// _player2hand = []int{5, 6, 7, 24, 7, 7, 8, 8, 8, 9, 9, 9, 10}
	// _player3hand = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	// _player4hand = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	return _player1hand, _player2hand, _player3hand, _player4hand
}

// // 抽出來以方便上面模組直接套用
// func drawhandtile(_player int, _tilesSea []int, _handTilesCounts int) (returnTile []int) {
// 	rand.Seed(time.Now().UnixNano())

// 	return returnTile
// }

func RemoveIndex(int_arr []int, index int) (return_arr []int) {
	return append(int_arr[:index], int_arr[index+1:]...)
}

// 指定玩家回傳此回合從牌海所抽取的牌
func DrawTile(_player int, _tilesSea *[]int) (ok bool, returnTile int) {
	//從目前牌海長度隨機出一個index
	ok = true
	rand.Seed(time.Now().UnixNano())

	// 檢查牌海是否還有牌
	if len(*_tilesSea) == 0 {
		ok = false
		returnTile = -1
		return ok, returnTile
	}

	// 抽取牌海中最後一張牌
	//n := rand.Intn(len(*_tilesSea))
	n := len(*_tilesSea) - 1

	returnTile = (*_tilesSea)[n]

	//從牌海中刪除摸到的那張牌
	(*_tilesSea)[n] = (*_tilesSea)[len((*_tilesSea))-1]
	*_tilesSea = (*_tilesSea)[:(len((*_tilesSea)) - 1)]

	return ok, returnTile
}
