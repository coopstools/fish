package layout

import "fmt"

type printer struct {
	output []rune
}

func (p *printer) print(c rune) {
	p.output = append(p.output, c)
	fmt.Print("\n" + string(p.output))
}

func (p *printer) printN(v int) {
	cs := []rune(fmt.Sprintf("%x", v))
	p.output = append(p.output, cs...)
	fmt.Print("\n" + string(p.output))
}

func (p *printer) terminate() {
	fmt.Println("\nProgram exited properly")
}
