package strategy

import "io"

// 一個叫做Outpput的interface
type Output interface {
	Draw() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

// DrawOutput 提供別人若要設定Writer和LogWriter的值
// 透過實作好Output的方法 SetLog和SetWriter來做設定
type DrawOutput struct {
	// Writer方法 會回傳io.Writer 型態
	Writer io.Writer
	// LogWriter方法 會回傳io.Writer 型態
	LogWriter io.Writer
}

// type為DrawOutput的strcut d 實作interface Output的SetLog
func (d *DrawOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}

// type為DrawOutput的strcut d 實作interface Output的SetWriter
func (d *DrawOutput) SetWriter(w io.Writer) {
	d.Writer = w
}
