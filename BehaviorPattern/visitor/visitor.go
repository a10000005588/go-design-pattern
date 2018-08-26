package visitor

import (
	"fmt"
	"io"
	"os"
)

// Both Message A&B have Msg: message to be printed
// MessageA 擁有Msg和io.Writer型態的Output (用來打印)
// 可以使用 Print 和 Accept
type MessageA struct {
	Msg    string
	Output io.Writer
}

// Print 負責將MessageA內的Output做初始化(m.Output=os.Stdout)
// 若已經初始化完畢，那就直接打印 MessageB的Output和Msg
// (會先透過Accpet，將MessageB的m 送給VisitB做處理，處理完再執行該Print方法)
func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

// Accept 負責接受一位Visitor(interface)
// 並且把MessageB 給予該參觀者A，執行 VisitA
func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

// ------------------------------------------------
// MessageB 擁有Msg和io.Writer型態的Output (用來打印)
// 可以使用 Print 和 Accept
type MessageB struct {
	Msg    string
	Output io.Writer
}

// Print 負責將MessageB內的Output做初始化(m.Output=os.Stdout)
// 若已經初始化完畢，那就直接打印 MessageB的Output和Msg
// (會先透過Accpet，將MessageB的m 送給VisitB做處理，處理完再執行該Print方法)
func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

// Accept 負責接受一位Visitor(interface)
// 並且把MessageB 給予該參觀者A，執行 VisitB
func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

// ------------------------------------------------
// Visitor interface has two method: VisitA & VisitB
// Vistor will encapsulate the logic
// only open to the struct which has implement the : VisitorA and VisitorB
// so that they can use the logic in Visitor.
// 白話一點形容，Vistor interface定義了這位參觀者可以拜訪哪些人(拜訪A&B)
type Visitor interface {
	// Visitor包了一個logic叫做VisitA。
	// VisitA 的工作為，如果有個VistorA過來了，將他帶過來的訊息 MessageA 做加工
	// 由於要對原訊息做加工，所以接收的是指標 *MessageA，而非是MessageA.
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type Visitable interface {
	Accept(Visitor)
}

//------------------------------------------------------------
// MessageVisitor會負責帶訊息 給 A或者B做加工
type MessageVisitor struct{}

// VisitA 對來參觀的A的原本MessageA的訊息m做加工
// 加上(Visited A)的訊息
func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}

// VisitB 對來參觀的A的原本MessageA的訊息m做加工
// 加上(Visited A)的訊息
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

//------------------------------------------------------------
// MsgFieldVisitorPrinter 負責對Visitor帶來的訊息做打印
type MsgFieldVisitorPrinter struct{}

// VisitA 對來參觀的A帶來的訊息做打印
func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}

// VisitB 對來參觀的A帶來的訊息做打印
func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}
