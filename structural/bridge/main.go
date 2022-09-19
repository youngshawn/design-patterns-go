package main

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
}
