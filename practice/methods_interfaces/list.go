package methods_interfaces

import "fmt"

// Printer is an abstract type, protocol, or contract that can be interfaced with.
// No implementation
type printer interface {
	print()
}

// List Going to store printer interface values here.
type list []printer

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
	// Setting up this type to check if it exists below
	type discounter interface {
		discount(float64)
	}
	for _, it := range l {
		//We don't have discount methods for all the things, only games. So, extract the game type with a type assertion.
		// isGame will make sure the game is the game we expect.
		discountItem, ok := it.(discounter)
		if !ok {
			continue
		}
		discountItem.discount(ratio)
	}

}
