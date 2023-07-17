package gamerule

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mode int64

const (
	// 遊戲模式
	Mode8Joker = 2
	//Mode46          = 3
	ModeBloodBattle = 4
	ModeTai         = 5
	ModePublic      = 3
)

type GameRule struct {
	JokerTurnIntos           []int32 // Length 34
	LimitTai                 int32
	CanEat                   bool
	CanPong                  bool
	CanKong                  bool
	ChangeTileSameColor      bool
	ChangeTileCount          int32
	KongAfterHuUseSimpleRule bool
	Debug                    bool
	CanUseJokerAsEye         bool
	CanJokersingleGon        bool
	CanJokerGon              bool
	CanMing                  bool
	LogAICsv                 bool
	JokerNumbers             int32
	GameMode                 int32 //暫定!!! 8Joker:2, 46:3, BloodBattle:4, Tai:5
	CanCrit                  bool
	LogAIParameter           bool
	GuoShouHu                bool
	CanTing                  bool
	SpecialMode              bool
	FanCount                 []int32 // Length 128
	XorTable                 []int32 // Length 16384
	TileUpperLimit           int32   // 手牌上限
	AvailableTiles           []int32 // 可用牌
	AvailableTileCount       int32   // 可用牌數量

	// 範例
	// Dragon: [8, 4, 4] 代表 8個中, 4個發, 4個白
	// Dot: [2,4,4,4,4,4,4,4,8] 代表 2張一萬, 4張二萬, 4張三萬, 4張四萬, 4張五萬, 4張六萬, 4張七萬, 4張八萬, 8張九萬
	Dot       []int32 // 萬 (1-9萬)
	Bamboo    []int32 // 筒 (1-9筒)
	Character []int32 // 條 (1-9條)
	Wind      []int32 // 風 (東 南 西 北)
	Dragon    []int32 // 三元 (中 發 白)
	Flower    []int32 // 花 (春 夏 秋 冬 梅 蘭 竹 菊)

	IsbuHua bool
}

// 設定 GameRule
func (gr *GameRule) SetGameRule(_gamerule string) (ok bool) {
	path, _ := os.Getwd()
	switch _gamerule {
	case "8Joker":
		filepath := path + "\\gamerule\\gamerule8joker.csv"
		gr.GameMode = Mode8Joker
		ok = gr.LoadGameRule(filepath)
		gr.IsbuHua = false
		//fmt.Println("LoadGameRule: " + filepath)
	//case "46":
	// filepath := path + "\\gamerule\\gamerule46.csv"
	// gr.GameMode = Mode46
	// ok = gr.LoadGameRule(filepath)
	case "BloodBattle":
		filepath := path + "\\gamerule\\gameruleBloodBattle.csv"
		gr.GameMode = ModeBloodBattle
		ok = gr.LoadGameRule(filepath)
	case "TaiwanMJ":
		filepath := path + "\\gamerule\\gameruleTai.csv"
		gr.GameMode = ModeTai
		ok = gr.LoadGameRule(filepath)
		gr.IsbuHua = true
	case "Public":
		filepath := path + "\\gamerule\\gamerulePublic.csv"
		gr.GameMode = ModePublic
		ok = gr.LoadGameRule(filepath)
		gr.IsbuHua = true
		fmt.Println("Set GameRule Public")
	default:
		ok = false
		fmt.Println("default")
	}
	fmt.Println("gameMode: ", gr.GameMode)
	return ok
}

func (gr *GameRule) LoadGameRule(filepath string) (ok bool) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer file.Close()

	// read file
	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return false
	}
	gr.SetGameRuleByArray(data)
	return true
}

func (gr *GameRule) SetGameRuleByArray(data [][]string) {
	// data[i][0] = field name
	// data[i][1] = field value

	_fancount := 0
	_xortable := 0

	for i := 0; i < len(data); i++ {
		if data[i][0] == "JokerTurnIntos" {
			strs := strings.Split(data[i][1], ",")
			ary := make([]int32, len(strs))
			for i, str := range strs {
				num, _ := strconv.ParseInt(str, 10, 32)
				ary[i] = int32(num)
				//ary[i], _ = strconv.ParseInt(str, 10, 64)
			}
			gr.JokerTurnIntos = ary
			continue
		}

		if data[i][0] == "LimitTai" {
			num, _ := strconv.ParseInt(data[i][1], 10, 32)
			gr.LimitTai = int32(num)
			//gr.LimitTai, _ = strconv.ParseInt(data[i][1], 10, 64)
			continue
		}

		if data[i][0] == "CanEat" {
			gr.CanEat, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanPong" {
			gr.CanPong, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanKong" {
			gr.CanKong, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "ChangeTileSameColor" {
			gr.ChangeTileSameColor, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "ChangeTileCount" {
			num, _ := strconv.ParseInt(data[i][1], 10, 32)
			gr.ChangeTileCount = int32(num)
			//gr.ChangeTileCount, _ = strconv.ParseInt(data[i][1], 10, 64)
			continue
		}

		if data[i][0] == "KongAfterHuUseSimpleRule" {
			gr.KongAfterHuUseSimpleRule, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "Debug" {
			gr.Debug, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanUseJokerAsEye" {
			gr.CanUseJokerAsEye, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanJokersingleGon" {
			gr.CanJokersingleGon, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanJokerGon" {
			gr.CanJokerGon, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanMing" {
			gr.CanMing, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "LogAICsv" {
			gr.LogAICsv, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "JokerNumbers" {
			num, _ := strconv.ParseInt(data[i][1], 10, 32)
			gr.JokerNumbers = int32(num)
			//gr.JokerNumbers, _ = strconv.ParseInt(data[i][1], 10, 64)
			continue
		}

		if data[i][0] == "GameMode" {
			num, _ := strconv.ParseInt(data[i][1], 10, 32)
			gr.GameMode = int32(num)
			//gr.GameMode, _ = strconv.ParseInt(data[i][1], 10, 64)
			continue
		}

		if data[i][0] == "CanCrit" {
			gr.CanCrit, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "LogAIParameter" {
			gr.LogAIParameter, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "GuoShouHu" {
			gr.GuoShouHu, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "CanTing" {
			gr.CanTing, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "SpecialMode" {
			gr.SpecialMode, _ = strconv.ParseBool(data[i][1])
			continue
		}

		if data[i][0] == "TileUpperLimit" {
			num, _ := strconv.ParseInt(data[i][1], 10, 32)
			gr.TileUpperLimit = int32(num)
			//gr.TileUpperLimit, _ = strconv.ParseInt(data[i][1], 10, 64)
			continue
		}

		if data[i][0] == "AvailableTiles" {
			strs := strings.Split(data[i][1], ",")
			ary := make([]int32, len(strs))
			for i, str := range strs {
				num, _ := strconv.ParseInt(str, 10, 32)
				ary[i] = int32(num)
				//ary[i], _ = strconv.ParseInt(str, 10, 64)
			}
			gr.AvailableTiles = ary
			gr.AvailableTileCount = 0
			for _, count := range gr.AvailableTiles {
				gr.AvailableTileCount += count
			}

			// 萬牌 從 AvailableTiles 中取出(0~8)
			gr.Dot = gr.AvailableTiles[0:9]
			// 筒牌 從 AvailableTiles 中取出(9~17)
			gr.Bamboo = gr.AvailableTiles[9:18]
			// 條牌 從 AvailableTiles 中取出(18~26)
			gr.Character = gr.AvailableTiles[18:27]
			// 風牌 從 AvailableTiles 中取出(27~30)
			gr.Wind = gr.AvailableTiles[27:31]
			// 三元牌 從 AvailableTiles 中取出(31~33)
			gr.Dragon = gr.AvailableTiles[31:34]
			// 花牌 從 AvailableTiles 中取出(34~42)
			gr.Flower = gr.AvailableTiles[34:42]

			// fmt.Println("Dot:", gr.Dot)
			// fmt.Println("Bamboo:", gr.Bamboo)
			// fmt.Println("Character:", gr.Character)
			// fmt.Println("Wind:", gr.Wind)
			// fmt.Println("Dragon:", gr.Dragon)
			// fmt.Println("Flower:", gr.Flower)
			// fmt.Println("AvailableTilesCount:", gr.AvailableTileCount)
			continue
		}

		if data[i][0] == "FanCount" {
			strs := strings.Split(data[i][1], ",")
			ary := make([]int32, len(strs))
			for i, str := range strs {
				num, _ := strconv.ParseInt(str, 10, 32)
				ary[i] = int32(num)
				//ary[i], _ = strconv.ParseInt(str, 10, 64)

				if ary[i] != 0 {
					_fancount++
				}
			}
			gr.FanCount = ary
			continue
		}

		if data[i][0] == "XorTable" {
			for j := 1; j < 4; j++ {
				strs := strings.Split(data[i][j], ",")
				ary := make([]int32, len(strs))
				for k, str := range strs {
					num, _ := strconv.ParseInt(str, 10, 32)
					ary[k] = int32(num)
					//ary[k], _ = strconv.ParseInt(str, 10, 64)

					if ary[k] != 0 {
						_xortable++
					}
				}
				// 把ary 串接到 gr.XorTable
				gr.XorTable = append(gr.XorTable, ary...)
			}
			// strs := strings.Split(data[i][1], ",")
			// ary := make([]int32, len(strs))
			// for i, str := range strs {
			// 	num, _ := strconv.ParseInt(str, 10, 32)
			// 	ary[i] = int32(num)
			// }
			// gr.XorTable = ary
			continue
		}
	}
	//fmt.Println(gr)
	// fmt.Println("GameRule FanCount:", _fancount)
	// fmt.Println("GameRule XorTable:", _xortable)
}
