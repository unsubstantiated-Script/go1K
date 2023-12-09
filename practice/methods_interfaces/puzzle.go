package methods_interfaces

import "fmt"

type Book struct {
	Title string
	Price money
}

// Print To the left is the receiver this attaches the method to the struct
func (b *Book) Print() {
	fmt.Printf("%-15s:  %.2s\n", b.Title, b.Price.String())
}
