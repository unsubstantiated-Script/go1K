package methods_interfaces

import "fmt"

type book struct {
	title string
	price money
}

// To the left is the reciever this attaches the method to the struct
func (b *book) print() {
	fmt.Printf("%-15s:  %.2s\n", b.title, b.price.string())
}
