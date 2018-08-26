package shapes

import (
	strategy "Design-Pattern/BehaviorPattern/strategy/example2"
	"fmt"
	"os"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

// 會根據s是text or image產生不同策略的Writer
// 並且回傳type為 Output的interface
// 只有實作Output的type strcut才能被回傳
// 這裡的範例為 DrawOutput 有實作Output的strcut.
func Factory(s string) (strategy.Output, error) {
	switch s {
	case TEXT_STRATEGY:
		// 回傳一個strcut都必須要用& 來做回傳
		return &TextSquare{
			// DrawOutput這裡已經將原本 只能直接寫入console和log
			// 抽離成text可以隨意寫入console或是log
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
