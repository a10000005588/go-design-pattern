package visitor

import "testing"

// TestHelper will implement the io.Writer interface (Write)
type TestHelper struct {
	Received string
}

// Put the written bytes on the "Received" field
func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	// create a testHepler instance.
	testHelper := &TestHelper{}
	// create a MessageVistor that will implement the Visitor interface.
	// 該visitor有著MessageVisitor的型態，
	// 有著VisitA 或是 VisitB的功能
	visitor := &MessageVisitor{}
	// 可以想像成 一位visitor 他會幫你把msg帶去給 A參觀，請A做加工
	//                              或是帶給   B參觀，請B做加工
	t.Run("MessageA test", func(t *testing.T) {
		// Create a MessageA struct that contains...
		// a string: "Hello World", and an Output which is a TestHelper strcut.
		// MessageA 實作了兩個方法:
		//    Print()
		//    Accept() : 包含接受一個MessageVistor Struct，並傳給 visitor(幫你帶訊息去參觀的strcut)
		msg := MessageA{
			Msg:    "Hello World",
			Output: testHelper,
		}
		// MessageA struct has implement the Accept.
		// The Accpet function will receive a visitor and
		// put the msg in its(Visitor) VisitA.
		// Visitor has VisitA function which will append extra message to the msg.

		// msg為MessageA struct. 先透過Accpet對msg.Msg和msg.Output做處理
		// Accept接受一個visitor，並且取用visitor.VisitorA()
		// 然後將msg 傳入給該visitor.VisitorA(msg)...做加工
		msg.Accept(visitor)
		// 再透過Print() 印出經過拜訪過後 被處理的msg
		msg.Print()

		expected := "A: Hello World (Visited A)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})

	t.Run("MessageB test", func(t *testing.T) {
		msg := MessageB{
			Msg:    "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "B: Hello World (Visited B)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}
	})
}
