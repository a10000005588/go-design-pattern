package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// GameState 負責根據 GameContext的狀態來做執行
type GameState interface {
	executeState(*GameContext) bool
}

// GameContext 會紀錄Game的重要資訊 ：答案、嘗試次數、贏了沒、下一場的狀態
type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

// StartState 負責實作executeState Interface
// 以下每個state都會實作 executeState
type StartState struct{}

func (s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)
	fmt.Println("Introduce a number a number of retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}

// FinishState 會執行贏的狀態該做的方法
type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	// 檢查GameContext 的Won 是true 還 false
	if c.Won {
		println("Congrats, you won")
	} else {
		fmt.Printf("You loose. The correct number was: %d\n", c.SecretNumber)
	}
	// 結束for loop 用false來停止
	return false
}

// AskState 會在用戶尚未挑戰成功繼續詢問的state
type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)

	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	// 若猜過一次，將嘗試次數 - 1
	c.Retries = c.Retries - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	// 若用完嘗試次數，將Next塞入 &FinishState{}
	if c.Retries == 0 {
		c.Next = &FinishState{}
	}

	return true
}

func main() {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}
	// 這是什麼奇怪的for寫法？！
	for game.Next.executeState(&game) {
	}

}
