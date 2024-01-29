# adapter-design-pattern
See example of adapter design pattern in go, this is a structural design pattern that allows two incompatible interfaces to work together.

## Documentation

### NewPrinter Interface
NewPrinter is an interface that represents a new printer. It has a method PrintStored that returns a string.

#### PrintStored Method
PrintStored is a method that prints the stored message. It returns the printed message as a string.

### PrinterAdapter Struct
PrinterAdapter is a struct that represents an adapter for a printer. It contains an OldPrinter of type LegacyPrinter and a Msg of type string.

#### OldPrinter
OldPrinter is a LegacyPrinter that represents the old printer.

#### Msg
Msg is a string that represents the message to be printed.

### PrintStored Method
PrintStored is a method of PrinterAdapter that prints the stored message using the adapter. It calls the Print method of the OldPrinter and appends "Adapter" to the returned message. The modified message is then printed to the console. If the OldPrinter is nil, no action is taken. The modified message is returned.