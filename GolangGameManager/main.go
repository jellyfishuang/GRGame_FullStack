package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"syscall"
	"time"

	DrawTool "GolangGameManager/drawtool"
	"GolangGameManager/gamemanager"
	GameRule "GolangGameManager/gamerule"
	TransformGolang "GolangGameManager/transform"

	AI_8JOKER "GolangGameManager/api8joker"
)

// GetOutboundIP get ip
// *****Code Download From https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go*****//
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.To4().String()
}

// func ApiStructUnitTest() {
// 	// cd GolangGameManager
// 	// go run main.go
// 	var data ApiStruct.Input_CheckReadHandFor14

// 	var arr []int64 = []int64{9, 9, 9, 12, 12, 12, 14, 14, 14, 15, 15, 15, 10, 10, -1}
// 	var gr8Joker GameRule.GameRule

// 	gr8Joker.SetGameRule("8Joker")

// 	data.Hand = arr
// 	data.Without = 2
// 	data.GameRule.SetGameRule(gr8Joker)

// 	fmt.Println("ApiStructUnitTest.")

// 	ans := ApiStruct.CheckReadHandFor14(data)
// 	fmt.Println(ans)

/*
	var d ApiStruct.Input_CheckChow
	d.AlreadyHu = false
	var t []int64 = []int64{9, 10, 11, 12, -1}
	d.Hand = t
	d.Tile = 13
	d.Without = 0

	r := ApiStruct.CheckChow(d)
	fmt.Println(r)

	cont := true
	if cont {
		return
	}
*/
//}

func drawtest() {
	_tilesSea := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	rand.Seed(time.Now().UnixNano())
	n1 := rand.Intn(len(_tilesSea))
	n2 := rand.Intn(len(_tilesSea))
	n3 := rand.Intn(len(_tilesSea))
	n4 := rand.Intn(len(_tilesSea))

	fmt.Println("first ", n1, n2, n3, n4, len(_tilesSea))

	fmt.Println("_tilesSea 1", _tilesSea)

	_tilesSea[2] = _tilesSea[len(_tilesSea)-1]
	fmt.Println("_tilesSea 2", _tilesSea)

	_tilesSea = _tilesSea[:(len(_tilesSea) - 1)]
	fmt.Println("_tilesSea 3", _tilesSea)

	Players0HandTiles, Players1HandTiles, Players2HandTiles,
		Players3HandTiles := DrawTool.DealTile(0, &_tilesSea, 5)

	fmt.Println("Players 0", Players0HandTiles)
	fmt.Println("Players 1", Players1HandTiles)
	fmt.Println("Players 2", Players2HandTiles)
	fmt.Println("Players 3", Players3HandTiles)
	fmt.Println("tilesSea", _tilesSea)
}

/*func checktest() {
	handtiles := []int{4, 4, 4, 4, 7, 9, 13, 16, 16, 16, 28, 29, 31}
	tabletile := 4
	//meldtiles := [][]int{{4, 4, 4}, {5, 6, 7}}

	//ok := gamemanager.CheckKong(handtiles, tabletile)

	// ok, chowTiles := gamemanager.CheckChow(handtiles, tabletile)
	// fmt.Println("ok:", ok)
	// fmt.Println("chowTile_Len:", len(chowTiles))
	// fmt.Println("chowTiles:", chowTiles)

	// ok, concealedkongtiles := gamemanager.CheckConcealedKong(handtiles, tabletile)
	// fmt.Println("ok:", ok)
	// fmt.Println("concealedkongtiles_Len:", len(concealedkongtiles))
	// fmt.Println("concealedkongtiles:", concealedkongtiles)

	//ok := gamemanager.CheckAddKong(meldtiles, tabletile)
	//fmt.Println("ok:", ok)
	//fmt.Println("meldtiles:", len(meldtiles))
	//fmt.Println("meldtiles:", meldtiles)

	// CheckHu
	gr := GameRule.GameRule{}
	cpp_gr := TransformGolang.CppGameRule{}
	gr.SetGameRule("8Joker")
	cpp_gr.SetGameRule(gr)

	handtiles = []int{1, 1, 3, 3, 3, 4, 4}
	_meldtiles := []TransformGolang.Meld{
		{Action: 1, Tiles: []int{2}},
		{Action: 2, Tiles: []int{5}},
	}
	tabletile = 4
	without := 2

	checkhu := gamemanager.CheckHu(handtiles, _meldtiles, tabletile, without, cpp_gr)
	fmt.Println("CheckHu", checkhu)
	fmt.Println("handtiles:", handtiles)
	fmt.Println("meldtiles:", _meldtiles)
	fmt.Println("tabletile:", tabletile, "without:", without)
}*/

func CheckAllplayerActionTest() {
	var room gamemanager.Room
	room.OntableTile = 3
	room.Players[0].HandTiles = []int{1, 1, 1, 1, 2, 2, 2, 2, 4, 4, 4, 4, 5}
	room.Players[1].HandTiles = []int{3, 3, 7, 8, 9, 10, 10, 10, 11, 12, 13}
	room.Players[2].HandTiles = []int{1, 1, 1, 1, 2, 2, 2, 2, 4, 4, 4, 4, 5}
	room.Players[3].HandTiles = []int{3, 3, 3, 1, 1, 1, 1, 2, 2, 2, 2, 4, 4}

	// room.Players[0].HandTiles = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// room.Players[1].HandTiles = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// room.Players[2].HandTiles = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// room.Players[3].HandTiles = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	ok := room.CheckAllplayerAction(2)
	fmt.Println(ok)
	room.Players[0].DoChow = true
	ChowSet := []int{3, 4, 5}
	room.DoChowPongKong("Chow", ChowSet)
	fmt.Println("Action:", room.DoAction.Action)
	fmt.Println("Player:", room.DoAction.Player)
	fmt.Println("Tiles:", room.DoAction.Tiles)

	// fmt.Println("player1.CanPon: ", room.Players[1].CanPon)
	// fmt.Println("player3.CanKong: ", room.Players[3].CanKong)
	// fmt.Println("player2.DoNothing:", room.Players[2].Donothing)
	// fmt.Println("NothingDo: ", room.NothingDo)

}

func ChowPongKongTest(mode int) {
	var player gamemanager.Player

	if mode == 1 {
		player.HandTiles = []int{1, 1, 1, 1, 2, 2, 2, 2, 4, 4, 4, 4, 5}
		player.MeldTiles = []TransformGolang.Meld{}
		discardTile := 3
		tiles := []int{4, 5}
		tile := 3

		player.CanChow, _ = gamemanager.CheckChow(player.HandTiles, tile)
		fmt.Println("CanChow:", player.CanChow)
		fmt.Println("Chow")
		fmt.Println("handtiles:", player.HandTiles)
		fmt.Println("meldtiles:", player.MeldTiles)

		player.Chow(tiles, discardTile)

		fmt.Println("handtiles:", player.HandTiles)
		fmt.Println("meldtiles:", player.MeldTiles)
	} else if mode == 2 {
		player.HandTiles = []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 4}
		player.MeldTiles = []TransformGolang.Meld{}
		tile := 4

		player.CanPon, player.CanPonSet = gamemanager.CheckPong(player.HandTiles, tile)
		fmt.Println("Pong")
		fmt.Println("CanPon:", player.CanPon)
		fmt.Println("CanPonSet", player.CanPonSet)
		fmt.Println("handtiles:", player.HandTiles)
		fmt.Println("meldtiles:", player.MeldTiles)

		player.Pong(tile)

		fmt.Println("handtiles:", player.HandTiles)
		fmt.Println("meldtiles:", player.MeldTiles)
	} else if mode == 3 {
		player.HandTiles = []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 4}
		player.MeldTiles = []TransformGolang.Meld{}
		tile := 3
		//isConcealedKong := false
		player.CanKong, player.CanKongSet = gamemanager.CheckKong(player.HandTiles, tile)
		fmt.Println("CanKong:", player.CanKong)
		fmt.Println("CanKongSet", player.CanKongSet)
		fmt.Println("Kong")
		fmt.Println("handtiles:", player.HandTiles)
		fmt.Println("meldtiles:", player.MeldTiles)

		//player.Kong(tile, isConcealedKong)

		fmt.Println("handtiles:", player.HandTiles)
		fmt.Println("meldtiles:", player.MeldTiles)
	}
}

func UpdateUnOpenPoolTest() {
	var player gamemanager.Player

	gr := GameRule.GameRule{}
	gr.SetGameRule("8Joker")

	player.InitPlayer(0, gr)
	player.HandTiles = []int{1, 1, 1, 1, 2, 2, 2, 2, 4, 4, 4, 4, 5}

	fmt.Println("UnOpenPool", player.UnOpenPool)

	player.UpdateUnOpenPool(player.HandTiles)
	fmt.Println("UnOpenPool", player.UnOpenPool)

	player.UpdateUnOpenPool([]int{17})
	fmt.Println("UnOpenPool", player.UnOpenPool)
}

func CheckAIHuPonKongTest() {

	var gr8Joker GameRule.GameRule
	var player gamemanager.Player

	gr8Joker.SetGameRule("8Joker")
	player.InitPlayer(0, gr8Joker)

	player.HandTiles = []int{3, 5, 20, 20, 20, 21, 22, 31, 31, 31}
	player.MeldTiles = []TransformGolang.Meld{
		{Action: 3, Tiles: []int{26}},
	}
	tile := 7
	player.UpdateUnOpenPool([]int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4})
	player.WithoutTile = 1

	player.AIHuChowPongKong(gr8Joker, tile)
}

func CheckAISelfHuKongTest() {

	var gr8Joker GameRule.GameRule
	var player gamemanager.Player

	gr8Joker.SetGameRule("8Joker")
	player.InitPlayer(0, gr8Joker)

	player.HandTiles = []int{2, 3, 5, 5, 13, 13, 15, 16, 17, 31}
	player.MeldTiles = []TransformGolang.Meld{
		{Action: 1, Tiles: []int{1}},
	}
	tile := 1
	player.UpdateUnOpenPool([]int{12, 12, 12, 17, 17, 17, 19, 19, 21, 21, 21, 31, 31})
	player.WithoutTile = 2

	player.AISelfHuKong(gr8Joker, tile)
}

func CheckAIThrowTest() {
	var gr8Joker GameRule.GameRule
	var player gamemanager.Player

	gr8Joker.SetGameRule("8Joker")
	player.InitPlayer(0, gr8Joker)

	player.HandTiles = []int{1, 1, 2, 4, 5, 8, 20, 20, 23, 23, 18, 19, 22}
	player.UpdateUnOpenPool([]int{1, 1, 2, 4, 5, 8, 20, 20, 23, 23, 18, 19, 22})
	player.WithoutTile = 1

	player.AIThrow(gr8Joker, 14)
}

func CheckAIThrowTestPublic() {
	var grPublic GameRule.GameRule
	var player gamemanager.Player

	grPublic.SetGameRule("Public")
	fmt.Println("game rule:", grPublic.GameMode)
	fmt.Println(grPublic.CanEat)
	player.InitPlayer(0, grPublic)

	player.HandTiles = []int{15, 17, 2, 4, 5, 8, 20, 20, 23, 23, 18, 19, 22}
	player.UpdateUnOpenPool([]int{1, 1, 2, 4, 5, 8, 20, 20, 23, 23, 18, 19, 22, -1})
	player.Uid = 1
	player.Without = 1
	player.SingleJokerAnGonCount = 0
	player.DoorSize = []int32{0, 0, 0, 0}
	player.ThrowSeq = []int32{9, 9, -1, -1, 10, 11, -1, -1}
	player.CheatTiles = []int32{-1}
	player.AlreadyHuTiles = []int32{-1}
	player.AlreadyHu = false
	player.AlreadyMing = false
	player.AlreadyCrit = false
	player.IsTingCard = false

	result := player.AIThrow(grPublic, 20)
	fmt.Println("result:", result)
}

func CheckAIHuPonKongTestPublic() {

	var grPublic GameRule.GameRule
	var player gamemanager.Player

	grPublic.SetGameRule("Public")
	player.InitPlayer(0, grPublic)

	//player.HandTiles = []int{2, 3, 4, 4, 5, 6, 22, 22, 23, 24}
	player.HandTiles = []int{5, 7, 8, 8, 8, 9, 10, 13, 17, 31}
	player.MeldTiles = []TransformGolang.Meld{
		{Action: 0, Tiles: []int{11, 12, 13}},
	}
	//player.MeldTiles = []TransformGolang.Meld{}
	tile := 6
	player.UpdateUnOpenPool([]int{5})
	//player.WithoutTile = 1

	player.FlowerTiles = []int{}
	player.Uid = 1
	player.Without = 2
	player.SingleJokerAnGonCount = 0
	player.DoorSize = []int32{0, 0, 0, 0}
	player.ThrowSeq = []int32{9, 9, -1, -1, 10, 11, -1, -1}
	player.CheatTiles = []int32{-1}
	player.AlreadyHuTiles = []int32{-1}
	player.AlreadyHu = false
	player.AlreadyMing = false
	player.AlreadyCrit = false
	player.IsTingCard = false

	player.AIHuChowPongKong(grPublic, tile)
}

func CheckAISelfHuKongTestPublic() {

	var grPublic GameRule.GameRule
	var player gamemanager.Player

	grPublic.SetGameRule("Public")
	player.InitPlayer(0, grPublic)

	player.HandTiles = []int{2, 3, 5, 5, 13, 13, 15, 16, 17, 13}
	player.MeldTiles = []TransformGolang.Meld{
		{Action: 1, Tiles: []int{1}},
	}
	tile := 1
	player.UpdateUnOpenPool([]int{12, 12, 12, 17, 17, 17, 19, 19, 21, 21, 21, 31, 31})
	player.FlowerTiles = []int{}
	player.Uid = 1
	player.Without = 2
	player.SingleJokerAnGonCount = 0
	player.DoorSize = []int32{0, 0, 0, 0}
	player.ThrowSeq = []int32{9, 9, -1, -1, 10, 11, -1, -1}
	player.CheatTiles = []int32{-1}
	player.AlreadyHuTiles = []int32{-1}
	player.AlreadyHu = false
	player.AlreadyMing = false
	player.AlreadyCrit = false
	player.IsTingCard = false

	result := player.AISelfHuKong(grPublic, tile)
	fmt.Println("result:", result)
}

func CheckHuTestPublic() {
	var grPublic GameRule.GameRule
	var player gamemanager.Player

	player.HandTiles = []int{3, 5, 20, 20, 20, 21, 22, 31, 31, 31, -1}
	player.MeldTiles = []TransformGolang.Meld{
		{Action: 3, Tiles: []int{7}},
	}
	tile := 1
	player.UpdateUnOpenPool([]int{2, 3})
	player.FlowerTiles = []int{}
	player.Uid = 1
	player.Without = 2
	player.SingleJokerAnGonCount = 0
	player.DoorSize = []int32{0, 0, 0, 0}
	player.ThrowSeq = []int32{9, 9, -1, -1, 10, 11, -1, -1}
	player.CheatTiles = []int32{-1}
	player.AlreadyHuTiles = []int32{-1}
	player.AlreadyHu = false
	player.AlreadyMing = false
	player.AlreadyCrit = false
	player.IsTingCard = false
	player.IsZimo = 0

	grPublic.SetGameRule("Public")
	canhu, ok := player.CheckHuFaanJudge(int32(tile), int32(player.IsZimo), grPublic)
	fmt.Println("CheckHu : ", canhu, ok)
}

func Lib_add(a, b int) {
	lib := syscall.NewLazyDLL("MyM16Headless.dll")
	fmt.Println("dll:", lib.Name)
	add := lib.NewProc("add")
	fmt.Println("+++++++NewProc:", add, "+++++++")

	ret, _, _ := add.Call(
		uintptr(a),
		uintptr(b),
	)
	fmt.Println("ret:", ret)
}

func DoAIDiscardTileTest() {
	var room gamemanager.Room

	room.InitRoom("8Joker", false, 0)
	room.DealTiles()
	room.OpenGame()

	fmt.Println("DoDiscardTileTest room.Players[0].OnDrawTile: ", room.Players[0].OnDrawTile)

	room.DoAIDiscardTile(0)

	//for idx := range room.Players {
	//fmt.Println(idx, "UnOpenPool ", room.Players[idx].UnOpenPool)
	//}
}

func TestTaiMJ() {
	//var room gamemanager.Room
	//room.InitRoom("TaiwanMJ")
	// room.DealTiles()
	// room.OpenGame()
}

func SetGameRuleTest() {
	var room gamemanager.Room

	fmt.Println("room.GameRule.GameMode", room.GameRule.GameMode)
	room.GameRule.SetGameRule("8Joker")
	fmt.Println("room.GameRule.GameMode", room.GameRule.GameMode)
}

func CheckHuTest() {
	var gr8Joker GameRule.GameRule
	var player gamemanager.Player

	player.WithoutTile = 1
	player.HandTiles = []int{3, 5, 20, 20, 20, 21, 22, 31, 31, 31, -1}
	player.MeldTiles = []TransformGolang.Meld{
		{Action: 3, Tiles: []int{26}},
	}
	tile := 4
	isZimo := 0

	gr8Joker.SetGameRule("8Joker")
	canhu, ok := player.CheckHuFaanJudge(int32(tile), int32(isZimo), gr8Joker)
	fmt.Println("CheckHu : ", canhu, ok)
}

func CheckWithoutTest() {
	var gr8Joker GameRule.GameRule
	var player gamemanager.Player

	player.HandTiles = []int{0, 3, 4, 5, 6, 6, 8, 8, 8, 10, 10, 14, 21}

	gr8Joker.SetGameRule("8Joker")

	player.AIChooseWithout(gr8Joker)

	fmt.Println("CheckWithout : ", player.WithoutTile)
}

func DLLTestInput() {
	AI_8JOKER.TestDLL()
}

func InitCsv() {
	//GameLog columns
	columnsName := [][]string{{"round", "hand", "playerIndex", "wind", "actionId", "actionCards", "handTiles", "discardTiles", "point", "AlreadyChow", "AlreadyPong",
		"AlreadyKong", "AlreadyConcealedKong", "AlreadyAddKong", "HuFannList", "cardPool"}}

	csvFile, err := os.Create("./gamemanager/test.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)

	//避免excel開啟亂碼
	csvFile.WriteString("\xEF\xBB\xBF")

	for _, column := range columnsName {
		if err := w.Write(column); err != nil {
			fmt.Println(err)

		}
	}
	w.Flush()
}

// func BaoJiaoTest() {

// 	var player gamemanager.Player
// 	player.CanHu = true
// 	player.IsBaoJi = true
// 	player.CanBaoJi = gamemanager.CheckBaoJi(player.CanHu, player.IsBaoJi)
// 	fmt.Println("CanBaoJi : ", player.CanBaoJi)
// }

// func MingPaiTest() {
// 	var player gamemanager.Player
// 	var room gamemanager.Room
// 	idx := 0
// 	room.PlayerNow = 0
// 	player.IsAlreadyHu = true
// 	player.CanMingPai = gamemanager.CheckMingPai(player.IsAlreadyHu, idx, room.PlayerNow)
// 	fmt.Println("CanMingPai : ", player.CanMingPai)
// }

func main() {
	fmt.Println("GameManager Server.")
	rand.Seed(time.Now().UnixNano())

	//DLLTestInput()
	//CheckWithoutTest()
	//CheckAIHuPonKongTest()a
	//CheckHuTest()
	//return
	//CheckAISelfHuKongTest()
	//CheckHuFaanJudgeTest()
	//CheckAIThrowTest()
	//CheckHuTestPublic()
	//InitCsv()

	// ApiStructUnitTest()

	// checktest()
	//ChowPongKongTest(1)
	//CheckAllplayerActionTest()
	//UpdateUnOpenPoolTest()
	//DoAIDiscardTileTest()
	//SetGameRuleTest()
	// 測試網路用
	//CheckHuTest()
	//gamemanager.WebsocketServerStart("localhost")
	gamemanager.WebsocketServerStart(GetOutboundIP())

	// 測試房間
	//room := new(gamemanager.Room)
	//room.InitRoom("8Joker")
	//room.DoAction()
	//BaoJiaoTest()
	//MingPaiTest()
	//CheckAIThrowTestPublic()
	//CheckAIHuPonKongTestPublic()
	//CheckAISelfHuKongTestPublic()
	//CheckAIHuPonKongTest()
	//CheckAIThrowTest()
	//Lib_add(3, 4)
}
