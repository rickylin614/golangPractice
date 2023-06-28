package main

import (
	"fmt"
	"math"
	"math/rand"
)

const screenWidth = 1024
const screenHeight = 768

// FishType表示不同種類的魚
type FishType int

const (
	FishType1 FishType = iota
	FishType2
	FishType3
	FishType4
	FishType5
)

var fishDeathProbabilities = map[FishType]float64{
	FishType1: 0.01,
	FishType2: 0.05,
	FishType3: 0.1,
	FishType4: 0.5,
	FishType5: 1,
}

// Fish表示一條魚
type Fish struct {
	X        int
	Y        int
	Type     FishType
	Alive    bool
	Velocity int
	Height   int
	Width    int
}

// BulletType表示不同種類的子彈
type BulletType int

const (
	BulletType1 BulletType = iota
	BulletType2
	BulletType3
)

var bulletDeathProbabilityIncrements = map[BulletType]float64{
	BulletType1: 0.001,
	BulletType2: 0.005,
	BulletType3: 0.01,
}

// Bullet表示一顆子彈
type Bullet struct {
	X        int
	Y        int
	Type     BulletType
	Velocity int
}

// Launcher表示一個發射台
type Launcher struct {
	X          int
	Y          int
	Direction  float64
	FiringRate int
	Cooldown   int
	Velocity   int
}

// GameState表示遊戲狀態
type GameState struct {
	FishPool  []*Fish
	Bullets   []*Bullet
	Launchers []*Launcher
	Velocity  int
	Score     int
	Time      int
	GameOver  bool
}

func main() {
	// 初始化遊戲狀態
	state := &GameState{
		FishPool:  make([]*Fish, 0),
		Bullets:   make([]*Bullet, 0),
		Launchers: make([]*Launcher, 0),
		Score:     0,
		Time:      0,
		GameOver:  false,
	}

	// 初始化魚池
	for i := 0; i < 100; i++ {
		fishType := FishType(rand.Intn(5))
		fish := &Fish{
			X:        rand.Intn(screenWidth),
			Y:        rand.Intn(screenHeight),
			Type:     fishType,
			Alive:    true,
			Velocity: rand.Intn(5) + 1,
		}
		state.FishPool = append(state.FishPool, fish)
	}

	// 初始化發射台
	for i := 0; i < 4; i++ {
		launcher := &Launcher{
			X:          screenWidth/2 + (i%2)*(screenWidth/2),
			Y:          screenHeight/2 + (i/2)*(screenHeight/2),
			Direction:  0,
			FiringRate: 2,
			Cooldown:   0,
		}
		state.Launchers = append(state.Launchers, launcher)
	}

	// 遊戲主循環
	for !state.GameOver {
		// 更新遊戲狀態
		state.Time++

		// 更新魚的位置
		for _, fish := range state.FishPool {
			fish.X += fish.Velocity
			fish.Y += fish.Velocity
			if fish.X < 0 || fish.X > screenWidth || fish.Y < 0 || fish.Y > screenHeight {
				fish.Alive = false
			}
		}

		// 更新子彈的位置
		for _, bullet := range state.Bullets {
			bullet.X += bullet.Velocity
			bullet.Y += bullet.Velocity
			if bullet.X < 0 || bullet.X > screenWidth || bullet.Y < 0 || bullet.Y > screenHeight {
				bullet.Velocity = 0
			}
		}

		// 檢查子彈是否擊中魚
		for _, bullet := range state.Bullets {
			for _, fish := range state.FishPool {
				if fish.Alive && isCollision(bullet, fish) {
					// 增加魚的死亡機率
					fishDeathProbability := fishDeathProbabilities[fish.Type]
					fishDeathProbability += bulletDeathProbabilityIncrements[bullet.Type]
					if rand.Float64() < fishDeathProbability {
						fish.Alive = false
						state.Score++
					}
					bullet.Velocity = 0
				}
			}
		}

		// 更新發射台冷卻時間
		for _, launcher := range state.Launchers {
			if launcher.Cooldown > 0 {
				launcher.Cooldown--
			}
		}

		// 發射子彈
		for _, launcher := range state.Launchers {
			// 根據發射台的方向和速度計算子彈的位置
			bulletX := launcher.X + int(float64(launcher.Velocity)*math.Cos(launcher.Direction))
			bulletY := launcher.Y + int(float64(launcher.Velocity)*math.Sin(launcher.Direction))
			bullet := &Bullet{
				X:        bulletX,
				Y:        bulletY,
				Type:     BulletType1, // 假設我們只有一種子彈
				Velocity: 5,
			}
			state.Bullets = append(state.Bullets, bullet)
			launcher.Cooldown = launcher.FiringRate
		}
		// 繪製遊戲畫面

		// 檢查遊戲是否結束
		if state.Time > 100 { // 假設遊戲持續時間為100秒
			state.GameOver = true
		}
	}

	fmt.Println("遊戲結束！得分：", state.Score)
}

// isCollision檢查子彈和魚是否發生碰撞
func isCollision(bullet *Bullet, fish *Fish) bool {
	// 模擬子彈和魚的面積碰撞
	if bullet.X > fish.X && bullet.X < fish.X+fish.Width && bullet.Y > fish.Y && bullet.Y < fish.Y+fish.Height {
		return true
	}
	return false
}
