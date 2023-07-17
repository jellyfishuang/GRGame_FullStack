package Api_8joker

type FaanType int

const (
	NoneFaan                  = iota
	BigThreeDollor            // 大三元
	SmallThreeDollor          // 小三元
	BigFourHappy              // 大四喜
	SmallFourHappy            // 小四喜
	OnlyOneNine               // 清么九
	OneNineWord               // 混么九
	Only13579                 // 寒鴉恓復
	MixFive                   // 全帶五
	MixOneNine                // 全帶么
	NoOneNine                 // 斷么九
	ContinueAllPair           // 連七對
	AllPair                   // 七對
	DragonAllPair             // 龍七對
	DoubleDragonAllPair       // 雙龍七對
	TripleDragonAllPair       // 參龍七對
	AllGreen                  // 綠一色
	AllWord                   // 字一色
	OneColor                  // 清一色
	TwoColor                  // 缺一門
	MixOneColor               // 混一色
	HaveThreeColorWindWord    // 五門齊
	NoWord                    // 無字
	AllBig                    // 全大
	AllMedium                 // 全中
	AllSmall                  // 全小
	FourEatSameOneColor       // 一色四同
	ThreeEatSameOneColor      // 一色三同
	TwoEatSameOneColorTwoTime // 一般高*2
	TwoEatSameOneColor        // 一般高
	TwoEatSameTwoColorTwoTime // 喜相逢*2
	TwoEatSameTwoColor        // 喜相逢
	ThreeEatSameThreeColor    // 三色三同
	FourEatUpOneColor         // 一色四步
	ThreeEatUpOneColor        // 一色三步
	TwoEatUpOneColor          // 一色二步
	ThreeEatUpThreeColor      // 三色三步
	ContinueSixOneColor       // 連六
	FiveConcealedPon          // 五暗刻
	FourConcealedPon          // 四暗刻
	ThreeConcealedPon         // 三暗刻
	TwoConcealedPon           // 雙暗刻
	OneConcealedPon           // 暗刻
	ThreeWindPon              // 三風刻
	TwoWindPon                // 雙風刻
	OneWindPon                // 風刻
	TwoWordPon                // 雙箭刻
	OneWordPon                // 箭刻
	ThreeSamePon              // 三同刻
	TwoSamePon                // 雙同刻
	TwoSamePonDouble          // 雙同刻*2
	FivePonUpOneColor         // 五節高
	FourPonUpOneColor         // 四節高
	ThreePonUpOneColor        // 三節高
	ThreePonUpThreeColor      // 三色三節
	AllEvenPon                // 全雙刻
	PonPon                    // 對對胡
	OneNinePon                // 幺九刻
	CircleWindPon             // 圈風刻
	DoorWindPon               // 門風刻
	PonPon258                 // 將對
	GonPlusAnGonFour          // 十八羅漢
	GonPlusAnGonThree         // 十二金釵
	GonFive                   // 明槓*5
	GonFour                   // 明槓*4
	GonThree                  // 明槓*3
	GonTwo                    // 明槓*2
	GonOne                    // 明槓*1
	AnGonFive                 // 暗槓*5
	AnGonFour                 // 暗槓*4
	AnGonThree                // 暗槓*3
	AnGonTwo                  // 暗槓*2
	AnGonOne                  // 暗槓*1
	JokerSingleGonOne         // #紅中單槓*1
	JokerSingleGonTwo         // #紅中單槓*2
	JokerSingleGonThree       // #紅中單槓*3
	JokerSingleGonFour        // #紅中單槓*4
	JokerSingleGonFive        // #紅中單槓*5
	JokerSingleGonSix         // #紅中單槓*6
	JokerSingleGonSeven       // #紅中單槓*7
	JokerSingleGonEight       // #紅中單槓*8
	JokerFourGonOne           // 紅中槓
	JokerFourleGonTwo         // 紅中槓*2
	SmallerFive               // 小於五
	BiggerFive                // 大於五
	DragonOneColor            // 清龍
	DragonFlowerColor         // 花龍
	DoubleDragonOneColor      // 一色雙龍
	DoubleDragonThreeColor    // 三色雙龍
	HoleHu                    // 坎張
	SideHu                    // 邊張
	OldYoung                  // 老少副*1
	OldYoungTwoTime           // 老少副*2
	RootOne                   // 根*1
	RootTwo                   // 根*2
	RootThree                 // 根*3
	EightJoker                // 八仙過海
	SixJoker                  // 六六大順
	FourJoker                 // 四方來財
	NoJoker                   // 素胡
	MergeDragon               // 組合龍
	AllNotDepend              // 全不靠
	DoorClean                 // 門清
	SixStarNoDepend           // 七星不靠
	MillionStone              // 百萬石
	KnowAllPonGon             // 金勾勾
	NineGate                  // 九蓮寶燈
	HundredForestBirds        // 百林鳥
	OneInWan                  // 萬里挑一
	HundredForest             // 百節林
	HundredDoller             // 百貫銅
	CantPushDown              // 推不倒
	NoFan                     // 無番和
	ThirteenOrphans           // 十三幺
	TinSingle                 // 單釣
	PushHu                    // 推倒胡
	NoCallFor                 // 不求人
	Zimo                      // 自摸
	SkyHu                     // 天胡
	GroundHu                  // 地胡
	GonUpFlower               // 槓上開花
	Rejuvenation              // 妙手回春
	HaidilaoMoon              // 海底撈月
	AfterGonHu                // 槓上炮
	LastCard                  // 絕張
	GrabGon                   // 搶槓胡
	TinBegin                  // 起手叫
	ShowHands                 // 明牌
	PinHu                     // 平胡
	OnlyTwoGon                // 五行八卦
	OnlyTinJoker              // 守中抱一
	NoOneEightTiao            // 井井有條
	OnlyHuFiveWan             // 捉五魁
	TwoColorConnect           // 雙節高
	TinSingleJoker            // 紅中釣
	AllEven                   // 全雙
	AllOdd                    // 全單
	BiggerThanSixty           // 六十劃
	SixContinue               // 六連順
	PeopleHu                  // 人胡
	CallChange                // 呼叫轉移
	Ting                      // 聽牌
	MingAnGon                 // 明暗槓
	FlowerCard                // 花牌
	TwoFiveEight              // 二五八將
	BankerWin                 // 莊贏
	FlowerGon                 // 花槓
	EightFlower               // 春暖花開
	FourInOne                 // 四歸一
	QiChiZhuLian              // 七尺珠簾
	KongQueDongNanFei         // 孔雀東南
	QiXianLanYue              // 七星攬月
	SiHuTenScreen             // 西湖十景
	XiBeiWuYin                // 西北無垠
	YinYangLiangYi            // 陰陽兩儀
	FuJiaYiFang               // 富甲一方
	XuRiDongSheng             // 旭日東升
	ZhuQueXuanWu              // 朱雀玄武
	QingLongBaiHu             // 青龍白虎
	BaiFaBaiZhong             // 百發百中
	HunSanJie                 // 混三節
	HongQueBaoXi              // 紅雀報喜
	FengHuaXueTue             // 風花雪月
	JiGuMaCao                 // 擊鼓罵曹
	DanFengChaoYang           // 丹鳳朝陽
	XiBeiYiPianBai            // 西北一片
	YiQiHuaSanQing            // 一氣化三
	SanSeSiJieGao             // 三色四節
	ErSeSiJieGao              // 二色四節
	ErSeSanJieGao             // 二色三節
	HanJiangDuDiao            // 寒江獨釣
	QuanDanKe                 // 全單刻
	HunDan                    // 混單
	HunShuang                 // 混雙
	BiYueXiuHua               // 閉月羞花
	SanHuaJuDing              // 三花聚鼎
	WuSeXiangXuan             // 五色相宣
	HuaQianYueXia             // 花前月下
	ChunSeManYuan             // 春色滿園
	EnumSIZE                  // just use getting enum size, not a EnumTai
)

func GetFaanName(input int) string {
	Faan := FaanType(input)

	switch Faan {
	case BigThreeDollor:
		return "大三元"
	case SmallThreeDollor:
		return "小三元"
	case BigFourHappy:
		return "大四喜"
	case SmallFourHappy:
		return "小四喜"
	case OnlyOneNine:
		return "清么九"
	case OneNineWord:
		return "混么九"
	case Only13579:
		return "寒鴉恓復"
	case MixFive:
		return "全帶五"
	case MixOneNine:
		return "全帶么"
	case NoOneNine:
		return "斷么九"
	case ContinueAllPair:
		return "連七對"
	case AllPair:
		return "七對"
	case DragonAllPair:
		return "龍七對"
	case DoubleDragonAllPair:
		return "雙龍七對"
	case TripleDragonAllPair:
		return "三龍七對"
	case AllGreen:
		return "綠一色"
	case AllWord:
		return "字一色"
	case OneColor:
		return "清一色"
	case TwoColor:
		return "缺一門"
	case MixOneColor:
		return "混一色"
	case HaveThreeColorWindWord:
		return "五門齊"
	case NoWord:
		return "無字"
	case AllBig:
		return "全大"
	case AllMedium:
		return "全中"
	case AllSmall:
		return "全小"
	case FourEatSameOneColor:
		return "一色四同"
	case ThreeEatSameOneColor:
		return "一色三同"
	case TwoEatSameOneColorTwoTime:
		return "一般高*2"
	case TwoEatSameOneColor:
		return "一般高"
	case TwoEatSameTwoColorTwoTime:
		return "喜相逢*2"
	case TwoEatSameTwoColor:
		return "喜相逢"
	case ThreeEatSameThreeColor:
		return "三色三同"
	case FourEatUpOneColor:
		return "一色四步"
	case ThreeEatUpOneColor:
		return "一色三步"
	case TwoEatUpOneColor:
		return "一色二步"
	case ThreeEatUpThreeColor:
		return "三色三步"
	case ContinueSixOneColor:
		return "連六"
	case FiveConcealedPon:
		return "五暗刻"
	case FourConcealedPon:
		return "四暗刻"
	case ThreeConcealedPon:
		return "三暗刻"
	case TwoConcealedPon:
		return "雙暗刻"
	case OneConcealedPon:
		return "暗刻"
	case ThreeWindPon:
		return "三風刻"
	case TwoWindPon:
		return "雙風刻"
	case OneWindPon:
		return "風刻"
	case TwoWordPon:
		return "雙箭刻"
	case OneWordPon:
		return "箭刻"
	case ThreeSamePon:
		return "三同刻"
	case TwoSamePon:
		return "雙同刻"
	case TwoSamePonDouble:
		return "雙同刻*2"
	case FivePonUpOneColor:
		return "五節高"
	case FourPonUpOneColor:
		return "四節高"
	case ThreePonUpOneColor:
		return "三節高"
	case ThreePonUpThreeColor:
		return "三色三高"
	case AllEvenPon:
		return "全雙刻"
	case PonPon:
		return "對對胡"
	case OneNinePon:
		return "么九刻"
	case CircleWindPon:
		return "圈風刻"
	case DoorWindPon:
		return "門風刻"
	case PonPon258:
		return "將對"
	case GonPlusAnGonFour:
		return "十八羅漢"
	case GonPlusAnGonThree:
		return "十二金釵"
	case GonFive:
		return "明槓*5"
	case GonFour:
		return "明槓*4"
	case GonThree:
		return "明槓*3"
	case GonTwo:
		return "明槓*2"
	case GonOne:
		return "明槓"
	case AnGonFive:
		return "暗槓*5"
	case AnGonFour:
		return "暗槓*4"
	case AnGonThree:
		return "暗槓*3"
	case AnGonTwo:
		return "暗槓*2"
	case AnGonOne:
		return "暗槓"
	case JokerSingleGonOne:
		return "紅中單槓*1"
	case JokerSingleGonTwo:
		return "紅中單槓*2"
	case JokerSingleGonThree:
		return "紅中單槓*3"
	case JokerSingleGonFour:
		return "紅中單槓*4"
	case JokerSingleGonFive:
		return "紅中單槓*5"
	case JokerSingleGonSix:
		return "紅中單槓*6"
	case JokerSingleGonSeven:
		return "紅中單槓*7"
	case JokerSingleGonEight:
		return "紅中單槓*8"
	case JokerFourGonOne:
		return "紅中槓"
	case JokerFourleGonTwo:
		return "紅中槓*2"
	case SmallerFive:
		return "小於五"
	case BiggerFive:
		return "大於五"
	case DragonOneColor:
		return "清龍"
	case DragonFlowerColor:
		return "花龍"
	case DoubleDragonOneColor:
		return "一色雙龍"
	case DoubleDragonThreeColor:
		return "三色雙龍"
	case HoleHu:
		return "坎張"
	case SideHu:
		return "邊張"
	case OldYoung:
		return "老少副*1"
	case OldYoungTwoTime:
		return "老少副*2"
	case RootOne:
		return "根*1"
	case RootTwo:
		return "根*2"
	case RootThree:
		return "根*3"
	case EightJoker:
		return "八仙過海"
	case SixJoker:
		return "六六大順"
	case FourJoker:
		return "四方來財"
	case NoJoker:
		return "素胡"
	case MergeDragon:
		return "組合龍"
	case AllNotDepend:
		return "全不靠"
	case DoorClean:
		return "門清"
	case SixStarNoDepend:
		return "七星不靠"
	case MillionStone:
		return "百萬石"
	case KnowAllPonGon:
		return "金勾勾"
	case NineGate:
		return "九蓮寶燈"
	case HundredForestBirds:
		return "百林鳥"
	case OneInWan:
		return "萬里挑一"
	case HundredForest:
		return "百節林"
	case HundredDoller:
		return "百貫銅"
	case CantPushDown:
		return "推不倒"
	case NoFan:
		return "無番和"
	case ThirteenOrphans:
		return "十三幺"
	case TinSingle:
		return "單釣"
	case PushHu:
		return "推倒胡"
	case NoCallFor:
		return "不求人"
	case Zimo:
		return "自摸"
	case SkyHu:
		return "天胡"
	case GroundHu:
		return "地胡"
	case GonUpFlower:
		return "槓上開花"
	case Rejuvenation:
		return "妙手回春"
	case HaidilaoMoon:
		return "海底撈月"
	case AfterGonHu:
		return "槓上炮"
	case LastCard:
		return "絕張"
	case GrabGon:
		return "搶槓胡"
	case TinBegin:
		return "起手叫"
	case ShowHands:
		return "明牌"
	case PinHu:
		return "平胡"
	case OnlyTwoGon:
		return "五行八卦"
	case OnlyTinJoker:
		return "守中抱一"
	case NoOneEightTiao:
		return "井井有條"
	case OnlyHuFiveWan:
		return "捉五魁"
	case TwoColorConnect:
		return "雙節高"
	case TinSingleJoker:
		return "紅中釣"
	case AllEven:
		return "全雙"
	case AllOdd:
		return "全單"
	case BiggerThanSixty:
		return "六十劃"
	case SixContinue:
		return "六連順"
	case PeopleHu:
		return "人胡"
	case CallChange:
		return "呼叫轉移"
	case Ting:
		return "聽牌"
	case MingAnGon:
		return "明暗槓"
	case FlowerCard:
		return "花牌"
	case TwoFiveEight:
		return "二五八將"
	case BankerWin:
		return "莊贏"
	case FlowerGon:
		return "花槓"
	case EightFlower:
		return "春暖花開"
	case FourInOne:
		return "四歸一"
	case QiChiZhuLian:
		return "七尺珠簾"
	case KongQueDongNanFei:
		return "孔雀東南飛"
	case QiXianLanYue:
		return "七星攬月"
	case SiHuTenScreen:
		return "西湖十景"
	case XiBeiWuYin:
		return "西北無垠"
	case YinYangLiangYi:
		return "陰陽兩儀"
	case FuJiaYiFang:
		return "富甲一方"
	case XuRiDongSheng:
		return "旭日東升"
	case ZhuQueXuanWu:
		return "朱雀玄武"
	case QingLongBaiHu:
		return "青龍白虎"
	case BaiFaBaiZhong:
		return "百發百中"
	case HunSanJie:
		return "混三節"
	case HongQueBaoXi:
		return "紅雀報喜"
	case FengHuaXueTue:
		return "風花雪月"
	case JiGuMaCao:
		return "擊鼓罵曹"
	case DanFengChaoYang:
		return "丹鳳朝陽"
	case XiBeiYiPianBai:
		return "西北一片白"
	case YiQiHuaSanQing:
		return "一氣化三清"
	case SanSeSiJieGao:
		return "三色四節高"
	case ErSeSiJieGao:
		return "二色四節高"
	case ErSeSanJieGao:
		return "二色三節高"
	case HanJiangDuDiao:
		return "寒江孤釣"
	case QuanDanKe:
		return "全單刻"
	case HunDan:
		return "混單"
	case HunShuang:
		return "混雙"
	case BiYueXiuHua:
		return "閉月羞花"
	case SanHuaJuDing:
		return "三花聚頂"
	case WuSeXiangXuan:
		return "五色相宣"
	case HuaQianYueXia:
		return "花前月下"
	case ChunSeManYuan:
		return "春色滿園"
	default:
		return "Unknown"
	}
}
