package entity

type Printer interface {
	Print()
}

func DisplayInfo(p Printer) {
	p.Print()
}
