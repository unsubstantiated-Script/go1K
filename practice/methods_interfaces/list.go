package methods_interfaces

import "fmt"

// Printer is an abstract type, protocol, or contract that can be interfaced with.
// No implementation
type item interface {
	print()
	discount(ratio float64)
}

// List Going to store item interface values here.
type list []item

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry we're waiting for stuffs.")
		return
	}
	for _, it := range l {
		it.print()
	}
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		it.discount(ratio)
	}

}
