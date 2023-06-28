package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Point struct {
	X int
	Y int
}

type Fish struct {
	ID               int
	Position         Point
	Size             int
	DeathProbability float64
}

type Cannon struct {
	Position  Point
	Direction int
}

type FishPool struct {
	Width  int
	Height int
	Fishes map[int]Fish
	Cannon Cannon
}

func (fp *FishPool) updateFishPositions() {
	for id, fish := range fp.Fishes {
		fish.Position.X += 5
		fish.Position.Y += 5
		fp.Fishes[id] = fish
	}
}

func (fp *FishPool) startUpdatingFishPositions() {
	go func() {
		for {
			fp.updateFishPositions()
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

type Bullet struct {
	ID                       int
	Position                 Point
	Size                     int
	IncreaseDeathProbability float64
}

func (fp *FishPool) updateBulletPosition(bullet Bullet) {
	bullet.Position.X += 10
	bullet.Position.Y += 10
}

func (fp *FishPool) startUpdatingBulletPosition(bullet Bullet) {
	go func() {
		for {
			fp.updateBulletPosition(bullet)
			time.Sleep(50 * time.Millisecond)
		}
	}()
}

func startWebSocketServer(fp *FishPool) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		go func(conn *websocket.Conn) {
			for {
				// 等待接收訊息
				_, msg, err := conn.ReadMessage()
				if err != nil {
					break
				}

				// 處理訊息並回傳結果
				result := handleMessage(fp, string(msg))
				err = conn.WriteMessage(websocket.TextMessage, []byte(result))
				if err != nil {
					break
				}
			}
			conn.Close()
		}(conn)
	})

	log.Println("WebSocket server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessage(fp *FishPool, msg string) string {
	// 處理訊息並回傳結果
	return "ok"
}

func NewFishPool(width int, height int) *FishPool {
	fp := &FishPool{
		Width:  width,
		Height: height,
		Fishes: make(map[int]Fish),
		Cannon: Cannon{
			Position: Point{
				X: 0,
				Y: 0,
			},
			Direction: 0,
		},
	}
	return fp
}

func main() {
	fp := NewFishPool(1024, 768)
	fp.updateFishPositions()
	startWebSocketServer(fp) // 監聽魚池

}
