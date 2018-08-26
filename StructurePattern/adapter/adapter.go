package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (m *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStore() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter // interface. won't do zero-initialization.
	// should use &funName{} which has implement LegacyPrinter interface.
	Msg string
}

func (p PrinterAdapter) PrintStore() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		// if no oldPrinter , adapter just print the original msg...
		newMsg = p.Msg
	}
	return
}
