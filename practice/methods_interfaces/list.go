package methods_interfaces

import "fmt"

type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry we're waiting for stuffs.")
		return
	}
	for _, it := range l {
		it.print()
	}
}
