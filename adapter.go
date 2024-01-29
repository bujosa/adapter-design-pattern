package main

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = s + "MyLegacyPrinter"
	println(newMsg)
	return
}

// NewPrinter is an interface that represents a new printer.
// It has a method PrintStored that returns a string.
type NewPrinter interface {
    // PrintStored is a method that prints the stored message.
    // It returns the printed message as a string.
    PrintStored() string
}


// PrinterAdapter is a struct that represents an adapter for a printer.
// It contains an OldPrinter of type LegacyPrinter and a Msg of type string.
type PrinterAdapter struct {
    // OldPrinter is a LegacyPrinter that represents the old printer.
    OldPrinter LegacyPrinter
    // Msg is a string that represents the message to be printed.
    Msg        string
}

// PrintStored prints the stored message using the adapter.
// It calls the Print method of the OldPrinter and appends "Adapter" to the returned message.
// The modified message is then printed to the console.
// If the OldPrinter is nil, no action is taken.
// The modified message is returned.
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = p.OldPrinter.Print(p.Msg)
		newMsg = newMsg + "Adapter"
		println(newMsg)
	}
	return
}