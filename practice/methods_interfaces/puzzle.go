package methods_interfaces

import "fmt"

type puzzle struct {
	title string
	price money
}

// Print To the left is the receiver this attaches the method to the struct
func (p puzzle) print() {
	fmt.Printf("%-15s:  %s\n", p.title, p.price.string())
}
