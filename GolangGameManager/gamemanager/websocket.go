package gamemanager

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func WebsocketServerStart(ip string) {

	game_port := ":1234"
	http_port := ":8964"
	go func() {
		//開啟Cocos Web Frontend
		log.Println("Cocos Web Frontend Start at " + ip + http_port)
		if err := http.ListenAndServe(http_port, http.FileServer(http.Dir("./web"))); err != nil {
			log.Fatal("Cocos Web Frontend ERROR:", err)
		}
	}()

	http.Handle("/", websocket.Handler(ClientAPI))
	log.Println("Game Server port Start at " + ip + game_port)
	if err := http.ListenAndServe(game_port, nil); err != nil {
		fmt.Println("Game Server ERROR:", err)
	}
}
