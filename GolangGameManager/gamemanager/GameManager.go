package gamemanager

import "time"

type GameManager struct {
	Rooms        map[uint8]*Room
	RoomCount    uint8
	Disconnected bool
	TimeStamp    time.Time //The disconnected time
}

// NewGameManager creates a new gameManager
func NewGameManager() (gm GameManager) {

	gm.Rooms = make(map[uint8]*Room)
	gm.RoomCount = 0
	gm.Disconnected = false
	gm.TimeStamp = time.Now()

	return gm
}
