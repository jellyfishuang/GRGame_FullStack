package gamemanager

import (
	AI_8JOKER "GolangGameManager/api8joker"
	AI_PUBLIC "GolangGameManager/apiPublic"

	GameRule "GolangGameManager/gamerule"
	TransformGolang "GolangGameManager/transform"

	"fmt"
)

// AI 玩家換三張(選擇階段)
func (_player *Player) AIChangeThreeTiles(_gamerule GameRule.GameRule, tiles *[]int) {
	if _gamerule.GameMode == GameRule.Mode8Joker {
		var parm AI_8JOKER.Input_AISelectChangeTile

		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.GameRule.SetGameRule32(_gamerule)
		parm.Strategies = _player.strategies

		// CheatTiles 作弊用牌 用不到 塞 -1
		arr := [1]int32{-1}
		parm.CheatTiles = []int32(arr[:])
		//fmt.Println("Before_ChangeTiles: ", _player.HandTiles)
		apiResult := AI_8JOKER.AISelectChangeTile(parm)
		//fmt.Println("ChangeTiles: ", apiResult)
		*tiles = append(*tiles, apiResult[:3]...)

	} /* else if _gamerule.GameMode == GameRule.Mode46 {

	} else if _gamerule.GameMode == GameRule.ModeBloodBattle {

	} else if _gamerule.GameMode == GameRule.ModeTai {

	} else {

	}*/
}

// AI 玩家定缺
func (_player *Player) AIChooseWithout(_gamerule GameRule.GameRule) {
	if _gamerule.GameMode == GameRule.Mode8Joker {
		var parm AI_8JOKER.Input_AISelectWithout

		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}

		// CheatTiles 作弊用牌 用不到 塞 -1
		parm.CheatTiles = []int32{-1}

		_player.ChooseWithout(AI_8JOKER.AISelectWithOut(parm))
	} /* else if _gamerule.GameMode == GameRule.Mode46 {

	} else if _gamerule.GameMode == GameRule.ModeBloodBattle {

	} else if _gamerule.GameMode == GameRule.ModeTai {

	} else {

	}*/
}

// func (_player *Player) AICheckHuPongKong(_tile int) {
// 	_player.CanChow = _player.CheckChow(int64(_tile))[0] != -1
// 	_player.CanPon = _player.CheckPong(int64(_tile))
// 	_player.CanKong = _player.CheckKong(int64(_tile))
// }

// AI 玩家的丟牌動作
func (_player *Player) AIThrow(_gamerule GameRule.GameRule, _tile int) (result int) {
	if _gamerule.GameMode == GameRule.Mode8Joker {

		fmt.Println("AI Before Throw: ", _player.HandTiles, _tile)
		result = _tile
		fmt.Println("AI Throw: ", result)

		var parm AI_8JOKER.Input_AIThrow
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = int32(_tile)
		parm.UnOpenPool = _player.UnOpenPool
		parm.Without = int32(_player.WithoutTile)
		_, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.GameRule.SetGameRule32(_gamerule)
		parm.Strategies = _player.strategies

		// CheatTiles 作弊用牌 用不到 塞 -1
		arr := [1]int32{-1}
		parm.CheatTiles = []int32(arr[:])

		fmt.Println("AIThrow_Hand: ", parm.Hand)
		fmt.Println("AIThrow_Pong: ", parm.Pong)
		fmt.Println("AIThrow_Kong: ", parm.Kong)
		fmt.Println("AIThrow_ConcealedKong: ", parm.ConcealedKong)

		fmt.Println("AIThrow_Tile: ", parm.Tile)
		fmt.Println("AIThrow_UnOpenPool:")
		for idx := range parm.UnOpenPool {
			fmt.Print(parm.UnOpenPool[idx], ",")
		}

		result = AI_8JOKER.AIThrow(parm)
		fmt.Println("AI Throw: ", result)

		// } else if _gamerule.GameMode == GameRule.ModeBloodBattle {
		// /result = _tile
		// } else if _gamerule.GameMode == GameRule.ModeTai {
		// result = _tile

		// //} else if _gamerule.GameMode == GameRule.ModePublic {
		// if _gamerule.GameMode == GameRule.Mode8Joker {
		// 	var parm AI_PUBLIC.Input_AIThrow
		// 	parm.Hand = make([]int32, len(_player.HandTiles)+1)
		// 	parm.Hand[len(_player.HandTiles)] = -1
		// 	for idx, hand := range _player.HandTiles {
		// 		parm.Hand[idx] = int32(hand)
		// 	}
		// 	parm.Tile = int32(_tile)
		// 	parm.UnOpenPool = _player.UnOpenPool

		// 	parm.Chow, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		// 	parm.GameRule.SetGameRule32(_gamerule)
		// 	parm.Strategies = _player.strategies

		// 	result = AI_PUBLIC.AIThrow(parm)
		// 	fmt.Println("AI Throw: ", result)
	} else if _gamerule.GameMode == GameRule.ModePublic {

		var parm AI_PUBLIC.Input_AIThrow
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = int32(_tile)

		parm.Chow, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.UnOpenPool = _player.UnOpenPool

		parm.Uid = int32(_player.Uid)
		parm.Without = int32(_player.Without)
		parm.SingleJokerAnGonCount = int32(_player.SingleJokerAnGonCount)
		parm.DoorSize = make([]int32, len(_player.DoorSize)+1)
		parm.DoorSize[len(_player.DoorSize)] = -1
		for idx, door := range _player.DoorSize {
			parm.DoorSize[idx] = int32(door)
		}
		parm.ThrowSeq = make([]int32, len(_player.ThrowSeq)+1)
		parm.ThrowSeq[len(_player.ThrowSeq)] = -1
		for idx, throw := range _player.ThrowSeq {
			parm.ThrowSeq[idx] = int32(throw)
		}
		parm.CheatTiles = make([]int32, len(_player.CheatTiles)+1)
		parm.CheatTiles[len(_player.CheatTiles)] = -1
		for idx, cheat := range _player.CheatTiles {
			parm.CheatTiles[idx] = int32(cheat)
		}
		parm.AlreadyHuTiles = make([]int32, len(_player.AlreadyHuTiles)+1)
		parm.AlreadyHuTiles[len(_player.AlreadyHuTiles)] = -1
		for idx, hu := range _player.AlreadyHuTiles {
			parm.AlreadyHuTiles[idx] = int32(hu)
		}
		parm.AlreadyHu = _player.AlreadyHu
		parm.AlreadyMing = _player.AlreadyMing
		parm.AlreadyCrit = _player.AlreadyCrit
		parm.IsTingCard = _player.IsTingCard

		//parm.Strategies = _player.strategies
		//parm.Strategies.PhaseParameter = make([]int32, len(_player.strategies.PhaseParameter)+1)

		//parm.GameRule.SetGameRule32(_gamerule)
		// fmt.Println("Strategies: ", parm.Strategies)
		// fmt.Println("AIThrow_Hand: ", parm.Hand)
		// fmt.Println("AIThrow_Chows: ", parm.Chow)
		// fmt.Println("AIThrow_Pong: ", parm.Pong)
		// fmt.Println("AIThrow_Kong: ", parm.Kong)
		// fmt.Println("AIThrow_ConcealedKong: ", parm.ConcealedKong)
		// fmt.Println("Uid: ", parm.Uid)
		// fmt.Println("Without: ", parm.Without)
		// fmt.Println("SingleJokerAnGonCount: ", parm.SingleJokerAnGonCount)
		// fmt.Println("DoorSize: ", parm.DoorSize)
		// fmt.Println("ThrowSeq: ", parm.ThrowSeq)
		// fmt.Println("CheatTiles: ", parm.CheatTiles)
		// fmt.Println("AlreadyHuTiles: ", parm.AlreadyHuTiles)
		// fmt.Println("AlreadyHu: ", parm.AlreadyHu)
		// fmt.Println("AlreadyMing: ", parm.AlreadyMing)
		// fmt.Println("AlreadyCrit: ", parm.AlreadyCrit)
		// fmt.Println("IsTingCard: ", parm.IsTingCard)

		result = AI_PUBLIC.AIThrow(parm)
		fmt.Println("AI Throw: ", result)

	} else {
		result = _tile
	}

	return result
}

// 其他玩家的回合時，AI 玩家的動作
func (_player *Player) AIHuChowPongKong(_gamerule GameRule.GameRule, _tile int) (ok bool) {
	ok = true
	if _gamerule.GameMode == GameRule.Mode8Joker {
		var parm AI_8JOKER.Input_AIHuPongKong

		// 把手牌轉成int32
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = int32(_tile)
		parm.UnOpenPool = _player.UnOpenPool
		parm.Without = int32(_player.WithoutTile)
		_, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.SignleJokerAnGonCount = _player.SignleJokerAnGonCount
		parm.AlreadyHu = _player.IsAlreadyHu

		// 明牌跟暴擊(待server之後補進遊戲後再開啟)
		// parm.AlreadyMing = _player.IsAlreadyMing
		// parm.AlreadyCrit = _player.IsAlreadyCrit
		parm.AlreadyMing = false
		parm.AlreadyCrit = false

		// CheatTiles 作弊用牌 用不到 塞 -1
		parm.CheatTiles = []int32{-1}

		parm.GameRule.SetGameRule32(_gamerule)
		parm.GameRule.FanCount = nil
		parm.GameRule.XorTable = nil
		parm.Strategies = _player.strategies
		// fmt.Println("AIHuPongKong_Hand: ", parm.Hand)
		// fmt.Println("AIHuPongKong_Tile: ", parm.Tile)
		// fmt.Println("AIHuPongKong_UnOpenPool:")
		// for idx := range parm.UnOpenPool {
		// 	fmt.Print(parm.UnOpenPool[idx], ",")
		// }
		//fmt.Println("AIHuPongKong_parm: ", parm)

		apiResult := AI_8JOKER.AIHuPongKong(parm)
		fmt.Println("AIHuPongKong: ", apiResult)

		// 0表示不做事、1表示胡、2表示碰、3表示槓、6表示暴擊(server後續補暴擊後再做)。
		if apiResult == 0 {

		} else if apiResult == 1 {
			_player.CanHu = true
			_player.DoHu = true
		} else if apiResult == 2 {
			_player.DoPon = true
		} else if apiResult == 3 {
			_player.DoKong = true
		} else if apiResult == 6 {
			_player.DoHu = true
		} else {
			fmt.Println("AIHuPongKong error")
			ok = false
		}

	} else if _gamerule.GameMode == GameRule.ModePublic {
		var parm AI_PUBLIC.Input_AIHuPongKong
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = int32(_tile)

		parm.Chow, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.UnOpenPool = _player.UnOpenPool
		//
		parm.Flower = make([]int32, 8)
		parm.Flower[len(_player.FlowerTiles)] = -1
		for idx, flower := range _player.FlowerTiles {
			parm.Flower[idx] = int32(flower)
		}
		parm.IsZimo = _player.IsZimo
		parm.Uid = int32(_player.Uid)
		parm.Without = int32(_player.Without)
		parm.SingleJokerAnGonCount = int32(_player.SingleJokerAnGonCount)
		parm.DoorSize = make([]int32, len(_player.DoorSize)+1)
		parm.DoorSize[len(_player.DoorSize)] = -1
		for idx, door := range _player.DoorSize {
			parm.DoorSize[idx] = int32(door)
		}
		parm.ThrowSeq = make([]int32, len(_player.ThrowSeq)+1)
		parm.ThrowSeq[len(_player.ThrowSeq)] = -1
		for idx, throw := range _player.ThrowSeq {
			parm.ThrowSeq[idx] = int32(throw)
		}
		parm.CheatTiles = make([]int32, len(_player.CheatTiles)+1)
		parm.CheatTiles[len(_player.CheatTiles)] = -1
		for idx, cheat := range _player.CheatTiles {
			parm.CheatTiles[idx] = int32(cheat)
		}
		parm.AlreadyHuTiles = make([]int32, len(_player.AlreadyHuTiles)+1)
		parm.AlreadyHuTiles[len(_player.AlreadyHuTiles)] = -1
		for idx, hu := range _player.AlreadyHuTiles {
			parm.AlreadyHuTiles[idx] = int32(hu)
		}
		parm.AlreadyHu = _player.AlreadyHu
		parm.AlreadyMing = _player.AlreadyMing
		parm.AlreadyCrit = _player.AlreadyCrit
		parm.IsTingCard = _player.IsTingCard

		parm.Strategies = _player.strategies
		fmt.Println("strategies: ", parm.Strategies)
		apiResult := AI_PUBLIC.AIHuPongKong(parm)
		fmt.Println("AIHuPongKong: ", apiResult)

		// 0表示不做事、1表示胡、2表示碰、3表示槓、6表示暴擊(server後續補暴擊後再做)。
		if apiResult == 0 {

		} else if apiResult == 1 {
			_player.CanHu = true
			_player.DoHu = true
		} else if apiResult == 2 {
			_player.DoPon = true
		} else if apiResult == 3 {
			_player.DoKong = true
		} else if apiResult == 6 {
			_player.DoHu = true
		} else {
			fmt.Println("AIHuPongKong error")
			ok = false
		}
	} /*else if _gamerule.GameMode == GameRule.ModeTai {

	  } else if _gamerule.GameMode == GameRule.Mode46 {

	  } else if _gamerule.GameMode == GameRule.ModeBloodBattle {

	  } else {

	  }*/
	return ok
}

// AI玩家在自己的回合時，AI 玩家的動作
func (_player *Player) AISelfHuKong(_gamerule GameRule.GameRule, _tile int) (ok bool) {
	ok = true
	if _gamerule.GameMode == GameRule.Mode8Joker {
		var parm AI_8JOKER.Input_AISelfHuKong

		// 把手牌轉成int32
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = int32(_tile)
		parm.UnOpenPool = _player.UnOpenPool
		parm.Without = int32(_player.WithoutTile)
		_, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.SignleJokerAnGonCount = _player.SignleJokerAnGonCount
		parm.AlreadyHu = _player.IsAlreadyHu

		// 明牌跟暴擊(待server之後補進遊戲後再開啟)
		// parm.AlreadyMing = _player.IsAlreadyMing
		// parm.AlreadyCrit = _player.IsAlreadyCrit
		parm.AlreadyMing = false
		parm.AlreadyCrit = false

		// CheatTiles 作弊用牌 用不到 塞 -1
		parm.CheatTiles = []int32{-1}

		parm.GameRule.SetGameRule32(_gamerule)
		parm.Strategies = _player.strategies

		apiResult := AI_8JOKER.AISelfHuKong(parm)
		fmt.Println("AISelfHuKong: ", apiResult)

		// 0: 不做事, 1: 胡, 2: 碰(用不到), 3: (暗)槓, 4: 紅中單槓(機制?), 5: 明牌, 6: 暴擊(server後續補明牌暴擊後再做)。
		// 5,6 有可能塞在 apiResult.OtherAction 裡面 (向 AI 確認)
		if apiResult.Action == 0 {

		} else if apiResult.Action == 1 {
			_player.CanSelfHu = true
			_player.DoSelfHu = true
		} else if apiResult.Action == 2 {
			//_player.DoPon = true
		} else if apiResult.Action == 3 {
			// 不確定是加槓還是暗槓, 已在server端確認過可執行何種
			fmt.Println("AISelfHuKong Kong")
			if _player.CanConcealedKong {
				_player.DoConcealedKong = true
				_player.DoKongTile = int(apiResult.Tile)
			} else if _player.CanAddKong {
				_player.DoAddKong = true
				_player.DoKongTile = int(apiResult.Tile)
			} else {
				fmt.Println("AISelfHuKong Kong Error")
				fmt.Println("Hand:", _player.HandTiles, "DrawTile:", _tile, "ActionTile: ", apiResult.Tile)
				ok = false
			}
		} else if apiResult.Action == 4 {
			fmt.Println("AISelfHuKong Red Kong")
			fmt.Println("Hand:", _player.HandTiles, "Tile:", _tile, "ActionTile: ", apiResult.Tile)

		} else if apiResult.Action == 5 {
			_player.CanSelfHu = true
			_player.DoSelfHu = true
		} else if apiResult.Action == 6 {
			_player.CanSelfHu = true
			_player.DoSelfHu = true
		} else {
			fmt.Println("AISelfHuKong error")
			ok = false
		}

	} else if _gamerule.GameMode == GameRule.ModePublic {
		var parm AI_PUBLIC.Input_AISelfHuKong

		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = int32(_tile)
		parm.Chow, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.UnOpenPool = _player.UnOpenPool
		parm.Flower = make([]int32, 8)
		parm.Flower[len(_player.FlowerTiles)] = -1
		for idx, flower := range _player.FlowerTiles {
			parm.Flower[idx] = int32(flower)
		}
		parm.IsZimo = _player.IsZimo
		parm.Uid = int32(_player.Uid)
		parm.Without = int32(_player.Without)
		parm.SingleJokerAnGonCount = int32(_player.SingleJokerAnGonCount)
		parm.DoorSize = make([]int32, len(_player.DoorSize)+1)
		parm.DoorSize[len(_player.DoorSize)] = -1
		for idx, door := range _player.DoorSize {
			parm.DoorSize[idx] = int32(door)
		}
		parm.ThrowSeq = make([]int32, len(_player.ThrowSeq)+1)
		parm.ThrowSeq[len(_player.ThrowSeq)] = -1
		for idx, throw := range _player.ThrowSeq {
			parm.ThrowSeq[idx] = int32(throw)
		}
		parm.CheatTiles = make([]int32, len(_player.CheatTiles)+1)
		parm.CheatTiles[len(_player.CheatTiles)] = -1
		for idx, cheat := range _player.CheatTiles {
			parm.CheatTiles[idx] = int32(cheat)
		}
		parm.AlreadyHuTiles = make([]int32, len(_player.AlreadyHuTiles)+1)
		parm.AlreadyHuTiles[len(_player.AlreadyHuTiles)] = -1
		for idx, hu := range _player.AlreadyHuTiles {
			parm.AlreadyHuTiles[idx] = int32(hu)
		}
		parm.AlreadyHu = _player.AlreadyHu
		parm.AlreadyMing = _player.AlreadyMing
		parm.AlreadyCrit = _player.AlreadyCrit
		parm.IsTingCard = _player.IsTingCard

		parm.Strategies = _player.strategies

		apiResult := AI_PUBLIC.AISelfHuKong(parm)
		fmt.Println("AISelfHuKong: ", apiResult)

		// 0: 不做事, 1: 胡, 2: 碰(用不到), 3: (暗)槓, 4: 紅中單槓(機制?), 5: 明牌, 6: 暴擊(server後續補明牌暴擊後再做)。
		// 5,6 有可能塞在 apiResult.OtherAction 裡面 (向 AI 確認)
		if apiResult.Action == 0 {

		} else if apiResult.Action == 1 {
			_player.CanSelfHu = true
			_player.DoSelfHu = true
		} else if apiResult.Action == 2 {
			//_player.DoPon = true
		} else if apiResult.Action == 3 {
			// 不確定是加槓還是暗槓, 已在server端確認過可執行何種
			fmt.Println("AISelfHuKong Kong")
			if _player.CanConcealedKong {
				_player.DoConcealedKong = true
				_player.DoKongTile = int(apiResult.Tile)
			} else if _player.CanAddKong {
				_player.DoAddKong = true
				_player.DoKongTile = int(apiResult.Tile)
			} else {
				fmt.Println("AISelfHuKong Kong Error")
				fmt.Println("Hand:", _player.HandTiles, "DrawTile:", _tile, "ActionTile: ", apiResult.Tile)
				ok = false
			}
		} else if apiResult.Action == 4 {
			//fmt.Println("AISelfHuKong Red Kong")
			//fmt.Println("Hand:", _player.HandTiles, "Tile:", _tile, "ActionTile: ", apiResult.Tile)

		} else if apiResult.Action == 5 {
			_player.CanSelfHu = true
			_player.DoSelfHu = true
		} else if apiResult.Action == 6 {
			_player.CanSelfHu = true
			_player.DoSelfHu = true
		} else {
			fmt.Println("AISelfHuKong error")
			ok = false
		}

	} /*else if _gamerule.GameMode == GameRule.Mode46 {

	} else if _gamerule.GameMode == GameRule.ModeBloodBattle {

	} else if _gamerule.GameMode == GameRule.ModeTai {

	} else {

	}*/
	return ok
}

// 檢查玩家是否可胡牌，以及胡牌之後的番型與番數
func (_player *Player) CheckHuFaanJudge(_tile int32, _isZimo int32, _gamerule GameRule.GameRule) (canhu bool, ok bool) {
	// Verify if the player can Hu the tile with different gamerules
	ok = true
	canhu = false
	if _gamerule.GameMode == GameRule.Mode8Joker {
		// println("8Joker_CheckHu")
		var parm AI_8JOKER.Input_CheckHu

		// 把手牌轉成int32
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.IsZimo = _isZimo
		parm.Tile = _tile
		parm.Without = int32(_player.WithoutTile)
		_, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.SignleJokerAnGonCount = int32(0)
		parm.GameRule.SetGameRule32(_gamerule)

		Output_CheckHu := AI_8JOKER.DoCheckHu(parm)

		var faanidx int

		for faanidx = 0; faanidx < len(Output_CheckHu.Faans); faanidx++ {
			if Output_CheckHu.Faans[faanidx] == -1 {
				break
			}
			_player.CanHuFaanList = append(_player.CanHuFaanList, Output_CheckHu.Faans[faanidx])
		}
		// fmt.Println("Output_CheckHu_Faans:", Output_CheckHu.Faans)
		// fmt.Println("Output_CheckHu_Counts:", Output_CheckHu.Counts)
		// fmt.Println("Output_CheckHu_Error_Code:", Output_CheckHu.Error_Code)

		if faanidx == 0 {
			// 代表沒有胡任何番型 = 沒有胡
			canhu = false
			//_player.CanHu = false
		} else {
			// 代表有胡
			for i := 0; i < faanidx; i++ {
				// 存取胡牌的番型名稱
				_player.CanHuFaanListStr = append(_player.CanHuFaanListStr, AI_8JOKER.GetFaanName(int(_player.CanHuFaanList[i])))
				// fmt.Println(AI_8JOKER.GetFaanName(int(_player.CanHuFaanList[i])))
			}
			canhu = true
			//_player.CanHu = true
			//_player.CanHuTile = int(_tile)
		}
	} else if _gamerule.GameMode == GameRule.ModePublic {
		var parm AI_PUBLIC.Input_CheckHu

		// 把手牌轉成int32
		parm.Hand = make([]int32, len(_player.HandTiles)+1)
		parm.Hand[len(_player.HandTiles)] = -1
		for idx, hand := range _player.HandTiles {
			parm.Hand[idx] = int32(hand)
		}
		parm.Tile = _tile
		parm.Chow, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles32(_player.MeldTiles)
		parm.Flower = make([]int32, 8)
		parm.Flower[len(_player.FlowerTiles)] = -1
		for idx, flower := range _player.FlowerTiles {
			parm.Flower[idx] = int32(flower)
		}
		parm.IsZimo = _player.IsZimo
		parm.Uid = int32(_player.Uid)
		parm.Without = int32(_player.Without)
		parm.SingleJokerAnGonCount = int32(_player.SingleJokerAnGonCount)
		parm.DoorSize = make([]int32, len(_player.DoorSize)+1)
		parm.DoorSize[len(_player.DoorSize)] = -1
		for idx, door := range _player.DoorSize {
			parm.DoorSize[idx] = int32(door)
		}
		parm.ThrowSeq = make([]int32, len(_player.ThrowSeq)+1)
		parm.ThrowSeq[len(_player.ThrowSeq)] = -1
		for idx, throw := range _player.ThrowSeq {
			parm.ThrowSeq[idx] = int32(throw)
		}
		parm.CheatTiles = make([]int32, len(_player.CheatTiles)+1)
		parm.CheatTiles[len(_player.CheatTiles)] = -1
		for idx, cheat := range _player.CheatTiles {
			parm.CheatTiles[idx] = int32(cheat)
		}
		parm.AlreadyHuTiles = make([]int32, len(_player.AlreadyHuTiles)+1)
		parm.AlreadyHuTiles[len(_player.AlreadyHuTiles)] = -1
		for idx, hu := range _player.AlreadyHuTiles {
			parm.AlreadyHuTiles[idx] = int32(hu)
		}
		parm.AlreadyHu = _player.AlreadyHu
		parm.AlreadyMing = _player.AlreadyMing
		parm.AlreadyCrit = _player.AlreadyCrit
		parm.IsTingCard = _player.IsTingCard

		Output_CheckHu := AI_PUBLIC.DoCheckHu(parm)

		//var faanidx int

		// for faanidx = 0; faanidx < len(Output_CheckHu.Faans); faanidx++ {
		// 	if Output_CheckHu.Faans[faanidx] == -1 {
		// 		fmt.Println("Output_CheckHu_Faans:", Output_CheckHu.Faans)
		// 		break
		// 	}
		// 	_player.CanHuFaanList = append(_player.CanHuFaanList, Output_CheckHu.Faans[faanidx])
		// }
		// fmt.Println("Output_CheckHu_Faans:", Output_CheckHu.Faans)
		// fmt.Println("Output_CheckHu_Counts:", Output_CheckHu.Counts)
		// fmt.Println("Output_CheckHu_Error_Code:", Output_CheckHu.Error_Code)

		// if faanidx == 0 {
		// 	// 代表沒有胡任何番型 = 沒有胡
		// 	canhu = false
		// 	//_player.CanHu = false
		// } else {
		// 	// 代表有胡
		// 	for i := 0; i < faanidx; i++ {
		// 		// 存取胡牌的番型名稱
		// 		//_player.CanHuFaanListStr = append(_player.CanHuFaanListStr, AI_8JOKER.GetFaanName(int(_player.CanHuFaanList[i])))
		// 		// fmt.Println(AI_8JOKER.GetFaanName(int(_player.CanHuFaanList[i])))
		// 	}
		// 	canhu = true
		// 	//_player.CanHu = true
		// 	//_player.CanHuTile = int(_tile)
		// }
		fmt.Println("Output_CheckHu:", Output_CheckHu)
	}

	return canhu, ok
}

func (_player *Player) InitWithout() {
	total := []int{0, 0, 0}
	for _, tile := range _player.HandTiles {
		if tile < 9 {
			total[0]++
		} else if tile < 18 {
			total[1]++
		} else if tile < 27 {
			total[2]++
		}
	}

	result := 0
	if total[result] > total[1] {
		result = 1
	}
	if total[result] > total[2] {
		result = 2
	}

	//fmt.Println("0: ", total[0], ", 1: ", total[1], ", 2: ", total[2])

	_player.WithoutTile = result
}
