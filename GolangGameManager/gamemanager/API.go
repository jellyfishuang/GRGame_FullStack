package gamemanager

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

type Message struct {
	Command string `json:"Command"`
	Data    string `json:"Data"`
}

type SendMessage struct {
	Command string      `json:"Command"`
	Data    interface{} `json:"Data"`
}

type CreateRoom struct {
	UserId   string `json:"userid"`
	GameRule string `json:"gamerule"`
}

type ChangeThreeTiles struct {
	UserId string `json:"userid"`
	Player int    `json:"player"`
	Tiles  [3]int `json:"tiles"`
}

type WithoutTiles struct {
	Player int `json:"player"`
	Tiles  int `json:"Tiles"`
}

type PlayerDiscardTile struct {
	UserId      string `json:"UserId"`
	Player      int    `json:"Player"`
	DiscardTile int    `json:"DiscardTile"`
}

type PlayerAIDiscardTile struct {
	UserId string `json:"UserId"`
	Player int    `json:"Player"`
}

type PlayerAction struct {
	UserId string `json:"UserId"`
	Player int    `json:"Player"`
	Action string `json:"Action"`
	Tiles  []int  `json:"Tiles"`
}

type AITest struct {
	Rounds   int    `json:"Rounds"`
	GameRule string `json:"gamerule"`
}

func ClientAPI(ws *websocket.Conn) {
	address := ws.Request().RemoteAddr
	fmt.Println(address, "connect.")

	//var game GameManager
	var gamestate Room
	channal := make(chan string, 32)

	go func() {
		for {
			if err := websocket.Message.Send(ws, <-channal); err != nil {
				break
			}
		}
	}()

	for {
		var request string
		if err := websocket.Message.Receive(ws, &request); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println(request, "Receive from", address)

		var msg Message
		//
		// index := strings.Index(request, "Data")
		// if index != -1 {
		// 	tmp, _ := json.Marshal(request[index+6 : len(request)-1])
		// 	request = request[:index+6] + string(tmp) + "}"
		// } else {
		// 	json.Unmarshal([]byte(request), &msg)
		// }

		json.Unmarshal([]byte(request), &msg)
		// fmt.Println("msg.Command:", msg.Command)
		// fmt.Println("msg.Data:", msg.Data)

		switch msg.Command {
		case "Connect":
			// example:
			// {"Command":"Connect"}
			// 新的user連線過來
			gamestate.GameState = StateConnect
			channal <- messageString("game.state", gamestate.ToByteArray())

		case "Createroom":
			// example:
			// {"Command":"Createroom","Data":"{\"UserId\":\"abc\",\"GameRule\":\"8Joker\"}"}

			fmt.Println("Createroom")
			// 新建房間
			var createRoom CreateRoom
			json.Unmarshal([]byte(msg.Data), &createRoom)
			ok := gamestate.InitRoom(createRoom.GameRule, false, 0)
			if ok {
				fmt.Println("Createroom send")
				gamestate.GameState = StateCreateRoom
				channal <- messageString("game.state", gamestate.ToByteArray())
			}

		case "DealTiles":
			// example:
			// {"Command":"DealTiles"}

			// 初始發牌, 並且回傳所有牌給前端(暫時訂立)
			gamestate.DealTiles()

			gamestate.GameState = StateDealTiles
			channal <- messageString("game.state", gamestate.ToByteArray())

		case "ChangeTiles":
			// example:
			// {"Command":"ChangeTiles","Data":"{\"Player\":0,\"Tiles\":[2,17,25]}"}

			// 執行換三張
			var changethreetiles ChangeThreeTiles
			json.Unmarshal([]byte(msg.Data), &changethreetiles)

			ok := gamestate.DoChangeThreeTiles(changethreetiles.Player, changethreetiles.Tiles)

			if ok {
				gamestate.GameState = StateChangeTiles
				channal <- messageString("game.state", gamestate.ToByteArray())
			}

		case "WithoutTiles":
			// example:
			// {"Command":"WithoutTiles","Data":"{\"Player\":0,\"Tiles\":1}"}

			// 執行定缺
			var withouttiles WithoutTiles
			json.Unmarshal([]byte(msg.Data), &withouttiles)
			fmt.Println("withouttiles:", withouttiles)
			ok := gamestate.DoWithoutTile(withouttiles.Player, withouttiles.Tiles)
			if ok {
				gamestate.GameState = StateWithoutTiles
				channal <- messageString("game.state", gamestate.ToByteArray())
			}

		case "Buhua":
			// 執行補花
			//{"Command":"Buhua"}

		case "OpenGame":
			// example:
			// {"Command":"OpenGame"}

			// 開門，發第一張牌給莊家玩家，以進入遊戲循環
			ok := gamestate.OpenGame()
			if ok {
				gamestate.GameState = StateAction
				gamestate.GameState = StateOpenGame
				channal <- messageString("game.state", gamestate.ToByteArray())
			}

		case "DiscardTile":
			// example:
			// {"Command":"DiscardTile","Data":"{\"UserId\":\"ABC\",\"Player\":0,\"DiscardTile\":7}"}

			// 接收前端真人玩家所丟出的牌，檢查是否合法並執行後檢查所有玩家能否做出對應動作
			var discardtile PlayerDiscardTile
			json.Unmarshal([]byte(msg.Data), &discardtile)

			// fmt.Println("discardtile:", discardtile)
			ok := gamestate.DoDiscardTile(discardtile.Player, discardtile.DiscardTile)
			if ok {
				gamestate.GameState = StateDiscard
				channal <- messageString("game.state", gamestate.ToByteArray())
			}

		case "AIDiscardTile":
			// example:
			// {"Command":"AIDiscardTile","Data":"{\"UserId\":\"ABC\",\"Player\":0}"}

			// 在AI進牌後，需要前端通知server，讓server執行AI丟牌
			var aidiscardtile PlayerAIDiscardTile
			json.Unmarshal([]byte(msg.Data), &aidiscardtile)

			ok := gamestate.DoAIDiscardTile(aidiscardtile.Player)
			if ok {
				gamestate.GameState = StateDiscard
				channal <- messageString("game.state", gamestate.ToByteArray())
			}

		case "Action":
			// example:
			// {"Command":"Action","Data":"{\"Player\":0,\"Action\":\"Pong\" ,\"Tiles\":[8,8,8]}"}
			// {"Command":"Action","Data":"{\"Player\":0,\"Action\":\"Exposed Kong\" ,\"Tiles\":[8,8,8,8]}"}
			// {"Command":"Action","Data":"{\"Player\":0,\"Action\":\"Nothing\"}"}

			// 	Action:
			// 	"Nothing", "Chow", "Pong", "Exposed Kong", "Concealed Kong", "Hu", "SelfDrawn" ...

			var action PlayerAction
			json.Unmarshal([]byte(msg.Data), &action)

			// 先重置所有玩家動作
			for idx := 0; idx < 4; idx++ {
				gamestate.Players[idx].ResetAction()
			}

			// 1.把真人傳到後端的 他選擇做什麼這件事塞進gamestate
			ok := gamestate.DoChowPongKong(action.Action, action.Tiles)
			//fmt.Println("doAction:", gamestate.DoAction)

			if ok {
				// 2.把 AI選擇做什麼塞進 gamestate
				ok, AItiles := gamestate.AIDoChowPongKong()
				action.Tiles = AItiles

				if !ok {
					fmt.Println("AIChowPongKong Fail!!!")
				}
				// 接著作動作優先序檢查
				//gamestate.CheckAllplayerAction(gamestate.PlayerNow)
				ok = gamestate.CheckActionPriority()
				fmt.Println("CheckActionPriority: ", ok)
			}

			if ok {
				// 如果有人可以做動作就執行
				fmt.Println("gamestate.NothingDo:", gamestate.NothingDo)
				if gamestate.NothingDo {
					// 沒有要執行就回到 下一家摸牌丟牌然後問前端丟回後端的遊戲循環
					// 摸牌需迴圈判定花牌，直到摸到非花牌為止
					isFlower := true
					for ok && isFlower {
						tempPlayerNow := gamestate.PlayerNow
						ok, isFlower = gamestate.DrawTiles()

						// 如果抽到花牌要在重抽的情況下，PlayerNow要重整回原先的值
						if isFlower {
							gamestate.PlayerNow = tempPlayerNow
						}
					}
					//gamestate.CheckAllplayerAction(gamestate.PlayerNow)

					if !ok && gamestate.Gameover {
						// 遊戲結束
						gamestate.GameState = StateGameOver
						channal <- messageString("game.state", gamestate.ToByteArray())
						break
					}
					if ok && gamestate.Players[gamestate.PlayerNow].CanSelfHu {
						// AI摸牌後選擇執行自摸
						fmt.Println("PlayerNow:", gamestate.PlayerNow, "CanSelfHu", gamestate.Players[gamestate.PlayerNow].CanSelfHu, "DoSelfHu", gamestate.Players[gamestate.PlayerNow].DoSelfHu)
						gamestate.AllPlayerDoAction(action.Tiles)
					}
				} else {
					// 執行動作

					//進到action.Tiles時是空的導致Call CheckFannList會出現Index out of range
					//fmt.Println("action.Tiles: ", action.Tiles)
					gamestate.AllPlayerDoAction(action.Tiles)
					//fmt.Println("AllPlayerDoAction")
				}
				if ok {
					fmt.Println("doAction_ok:", ok, "NothingDo:", gamestate.NothingDo, "PlayerNow:", gamestate.PlayerNow)
					if !gamestate.NothingDo && gamestate.PlayerNow != gamestate.RealPlayer {
						// 如果是AI執行的就可以直接丟牌到桌上，並檢查所有玩家能否做出對應動作並扔給前端

						if !(gamestate.DoAction.Action == "Hu" || gamestate.DoAction.Action == "SelfHu") {
							// AI玩家做吃碰槓後, AI要選擇丟哪張牌
							//gamestate.OntableTile = _playernow.AIThrow(gamestate.GameRule, _playernow.OnDrawTile)
							fmt.Println("玩家", gamestate.PlayerNow, "吃碰槓後丟牌", gamestate.OntableTile)
							//ok = gamestate.DoAIDiscardTile(gamestate.PlayerNow)

						} else if gamestate.DoAction.Action == "Hu" {
							// AI 玩家胡別人牌 不須丟牌 過水換下一家摸牌
							//afterhu = true
							gamestate.PlayerNow = gamestate.DoAction.Player
							fmt.Println("AI玩家胡別人牌 不須丟牌 過水換下一家摸牌")

						} else if gamestate.DoAction.Action == "SelfHu" {
							// AI玩家自摸, 自摸後要丟這巡摸進而判定自摸的牌, 接著輪下家摸牌
							//afterhu = true
							//gamestate.OntableTile = _playernow.AIThrow(gamestate.GameRule, _playernow.OnDrawTile)
							fmt.Println("玩家", gamestate.PlayerNow, "自摸丟牌", gamestate.Players[gamestate.PlayerNow].OnDrawTile)
							//ok = gamestate.DoDiscardTile(gamestate.PlayerNow, gamestate.Players[gamestate.PlayerNow].OnDrawTile)
						}

						if !ok {
							fmt.Println("AI動作後丟牌失敗")
						}

						//fmt.Println("桌上牌: ", gamestate.OntableTile)
						//fmt.Println("玩家", gamestate.PlayerNow, "手牌: ", _playernow.HandTiles)
					}

					gamestate.GameState = StateAction
					channal <- messageString("game.state", gamestate.ToByteArray())
				}
			}
		case "AITest":
			// to do AI test function in many rounds // +mode json  data
			var AItest AITest
			json.Unmarshal([]byte(msg.Data), &AItest)
			fmt.Println("執行回合數: ", AItest.Rounds)

			for i := 0; i < AItest.Rounds; i++ {
				gamestate.RunGame(AItest.GameRule, i+1)

			}
		}
	}
	fmt.Println(address, "disconnect.")
}

// func messageString(command string, data string, ok bool) string {
// 	var output gameManager

// 	switch command {
// 	case "roomRequest":
// 		if ok {
// 			output.roomRequest = 1
// 			output.roomId = 9999
// 		} else {
// 			output.roomRequest = 0
// 			output.roomId = -1
// 		}
// 	default:
// 	}

// 	returnValue, _ := json.Marshal(output)
// 	fmt.Println(returnValue)
// 	return string(returnValue)
// }

func messageString(command string, data []byte) string {
	jsonBytes, _ := json.Marshal(SendMessage{Command: command, Data: (json.RawMessage(data))})
	fmt.Println("send")
	return string(jsonBytes)
}
