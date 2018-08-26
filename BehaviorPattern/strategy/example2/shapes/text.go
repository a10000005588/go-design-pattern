package shapes

import strategy "Design-Pattern/BehaviorPattern/strategy/example2"

// 包住一個有實作Output {Draw() SetLog() SetWriter()}介面的strcut
// 叫做DrawOutput
// 注意這裡並不是繼承，而是Composite的做法
// 並不會隨著DrawOutput的程式碼變動而改動TextSquare
// 而是因為 DrawOutput“自己主動”實作了Output interface
// 所以若Output interface新增了其他的方法做擴充
// 也不會影響到DrawOutput。。。

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	// 在console中寫入 "Circle"的字串
	t.Writer.Write([]byte("Circle"))
	return nil
}
