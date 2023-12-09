package methods_interfaces

import "fmt"

type book struct {
	title string
	price money
}

// Print To the left is the receiver this attaches the method to the struct
func (b book) print() {
	fmt.Printf("%-15s:  %s\n", b.title, b.price.string())
}
