package methods_interfaces

import "fmt"

// List Going to store item interface values here.
type list []*product

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry we're waiting for stuffs.")
		return
	}
	for _, p := range l {
		p.print()
	}
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}

}
