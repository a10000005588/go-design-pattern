package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}

	returnMsg := adapter.PrintStore()
	if returnMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s \n", returnMsg)
	}

	adapter2 := PrinterAdapter{
		OldPrinter: nil,
		Msg:        msg,
	}
	returnMsg2 := adapter2.PrintStore()

	if returnMsg2 != "Hello World!" {
		t.Errorf("Message didn't match: %s \n", returnMsg2)
	}
}
