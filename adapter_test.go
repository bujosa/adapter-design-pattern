package main

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	
	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Hello World!MyLegacyPrinterAdapter" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}