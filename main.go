package main

import "fmt"

type Printer interface {
	Print() string
}

type Scanner interface {
	Scan() string
}

type Faxer interface {
	Fax() string
}

type PrintScanner interface {
	Printer
	Scanner
}

type FaxerPrinter interface {
	Faxer
	Printer
}

type FaxerPrinterScanner interface {
	Faxer
	Printer
	Scanner
}

type firstPrinter struct{}

func (fp firstPrinter) Print() string {
	return "First printer printing..."
}

func (fp firstPrinter) Scan() string {
	return "First printer scanning..."
}

func (fp firstPrinter) Fax() string {
	return "First printer scanning...."
}

type secondPrinter struct{}

func (sp secondPrinter) Print() string {
	return "Second printer printing..."
}

func (sp secondPrinter) Scan() string {
	return "Second printer scanning..."
}

func (sp secondPrinter) Fax() string {
	return "Second printer scanning...."
}

func process(equipment FaxerPrinterScanner) {
	fmt.Println("Running Print.... ", equipment.Print())
	fmt.Println("Running Scan..... ", equipment.Scan())
	fmt.Println("Running Fax...... ", equipment.Fax())
}

func main() {
	fPrinter := firstPrinter{}
	sPrinter := secondPrinter{}

	process(fPrinter)
	process(sPrinter)
}
