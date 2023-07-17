package Api_Tai

type FaanType int

const (
	// 一般番型
	NoneFaan     = iota
	SkyHu        // 天胡 24台
	GroundHu     // 地胡 16台
	BigFourHappy // 大四喜 16台
	AllWord      // 字一色 16台

	ThirteenOrphans // 七搶一 8台
	EightImmortals  // 八仙過海 8台

	BigThreeDollor   // 大三元 8台
	SmallFourHappy   // 小四喜 8台
	OneColor         // 清一色 8台
	FiveConcealedPon // 五暗刻 8台

	FourConcealedPon // 四暗刻 5台

	SmallThreeDollor // 小三元 4台
	MixOneColor      // 混一色 4台
	AllPongs         // 碰碰胡 4台

	ThreeConcealedPon // 三暗刻 2台
	PinHu             // 平胡 2台
	AllPairs          // 全求人 2台
	FlowerGon         // 花槓 2台

	DoorClean // 門清 1台
	// 自摸 1台
	// 不求人 1台 (門清自摸一摸三) 1+1+1=3台

	// 三元台 1台 (中發白刻子)
	// 圈風台 1台 (東南西北刻子)
	// 門風台 1台 (東南西北刻子)
	// 正花 1台 (花牌)

	// 單聽 1台

)
