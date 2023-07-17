package gamemanager

// 單純確認此玩家的手牌與桌上的牌可不可以讓玩家吃碰槓胡(cando)
// AI或真人玩家選擇要不要做則是另外引用AI模組

// 檢查吃之後，要把可以吃的組合回傳，組合可能不只一組(暗槓同理)
func CheckChow(_handTiles []int, _ontableTile int) (canChow bool, chowTiles [][]int) {
	canChow = false

	// 可以吃的組合只有在 0~26(萬筒條)
	if _ontableTile < 27 {
		// 檢查前後兩張牌是否存在即可
		front2Tile := false
		front1Tile := false
		back1Tile := false
		back2Tile := false

		for _, tile := range _handTiles {
			if tile == _ontableTile-2 {
				front2Tile = true
			} else if tile == _ontableTile-1 {
				front1Tile = true
			} else if tile == _ontableTile+1 {
				back1Tile = true
			} else if tile == _ontableTile+2 {
				back2Tile = true
			}
		}

		if _ontableTile == 0 || _ontableTile == 9 || _ontableTile == 18 {
			front1Tile = false
			front2Tile = false
		} else if _ontableTile == 1 || _ontableTile == 10 || _ontableTile == 19 {
			front2Tile = false
		} else if _ontableTile == 9 || _ontableTile == 17 || _ontableTile == 26 {
			back1Tile = false
			back2Tile = false
		} else if _ontableTile == 8 || _ontableTile == 16 || _ontableTile == 25 {
			back2Tile = false
		}

		if front2Tile && front1Tile {
			canChow = true
			chowTiles = append(chowTiles, []int{_ontableTile - 2, _ontableTile - 1, _ontableTile})
		}
		if front1Tile && back1Tile {
			canChow = true
			chowTiles = append(chowTiles, []int{_ontableTile - 1, _ontableTile, _ontableTile + 1})
		}
		if back1Tile && back2Tile {
			canChow = true
			chowTiles = append(chowTiles, []int{_ontableTile, _ontableTile + 1, _ontableTile + 2})
		}
	}

	return canChow, chowTiles
}

func CheckPong(_handTiles []int, _ontableTile int) (canPong bool, pongTiles []int) {
	canPong = false

	sameTileCount := 0
	for _, tile := range _handTiles {
		if tile == _ontableTile {
			sameTileCount++

		}
	}

	if sameTileCount >= 2 {
		canPong = true
		pongTiles = append(pongTiles, _ontableTile, _ontableTile, _ontableTile)
	}

	return canPong, pongTiles
}

func CheckKong(_handTiles []int, _ontableTile int) (canKong bool, kongTiles [][]int) {
	canKong = false

	sameTileCount := 0
	for _, tile := range _handTiles {
		if tile == _ontableTile {
			sameTileCount++
		}
	}

	if sameTileCount >= 3 {
		canKong = true
		kongTiles = append(kongTiles, []int{_ontableTile, _ontableTile, _ontableTile, _ontableTile})
	}

	return canKong, kongTiles
}

// 摸牌後對手牌與摸進的牌做暗槓檢查
func CheckConcealedKong(_handTiles []int, _drawTile int) (canConcealedKong bool, concealedKongTiles [][]int) {
	canConcealedKong = false

	sameTileCount := 0
	for _, tile := range _handTiles {
		if tile == _drawTile {
			sameTileCount++
		}
	}

	if sameTileCount >= 3 {
		canConcealedKong = true
		concealedKongTiles = append(concealedKongTiles, []int{_drawTile, _drawTile, _drawTile, _drawTile})
	}

	// 檢查原先手牌內是否有可暗槓的牌(有4張)
	for _, tile := range _handTiles {
		sameTileCount = 0
		for _, tile2 := range _handTiles {
			if tile == tile2 {
				sameTileCount++
			}
		}

		if sameTileCount >= 4 {
			canConcealedKong = true
			concealedKongTiles = append(concealedKongTiles, []int{tile, tile, tile, tile})
		}
	}

	return canConcealedKong, concealedKongTiles
}

// 加槓要對摸進的牌與刻子做檢查
func CheckAddKong(_meldTiles []int, _drawTile int) (canAddKong bool, canAddKongSet []int) {
	canAddKong = false
	sameTileCount := 0
	for _, tile := range _meldTiles {
		if tile == _drawTile {
			sameTileCount++
		}
	}

	if sameTileCount == 3 {
		canAddKong = true
		canAddKongSet = append(canAddKongSet, _drawTile, _drawTile, _drawTile, _drawTile)
	}

	return canAddKong, canAddKongSet
}

/*
func CheckHu(_handTiles []int, _meldTiles []TransformGolang.Meld, _ontableTile int, _without int, _gamerule TransformGolang.CppGameRule, _isZimo ...bool) (canHu bool) {
	canHu = false

	var parm ApiStruct.Input_VerifyHu

	parm.Hand = TransformGolang.CppTiles(_handTiles)
	parm.Tile = int64(_ontableTile)
	_, parm.Pong, parm.Kong, parm.ConcealedKong = TransformGolang.CppMeldTiles(_meldTiles)
	parm.Without = int64(_without)

	if len(_isZimo) > 0 {
		parm.IsZimo = _isZimo[0]
	} else {
		parm.IsZimo = true
	}

	parm.GameRule = _gamerule

	canHu = (ApiStruct.VerifyHu(parm) == 1)

	return canHu
}


func CheckSelfDraw(_handTiles []int, _meldTiles [][]int, _drawTile int) (canSelfDraw bool) {
	canSelfDraw = false

	return canSelfDraw
}
*/

func CheckBaoJi(CanHu bool, isBaoJi bool) (CanBaoJi bool) {
	CanBaoJi = false
	if CanHu && isBaoJi == false {
		CanBaoJi = true
	} else {
		return
	}
	return CanBaoJi
}

func CheckMingPai(IsAlreadyHu bool, idx int, playerNow int) (CanMingPai bool) {
	CanMingPai = false
	if IsAlreadyHu == false && idx == playerNow {
		CanMingPai = true
	} else {
		return
	}
	return CanMingPai
}
