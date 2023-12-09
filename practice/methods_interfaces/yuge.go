package methods_interfaces

import "fmt"

type yuge struct {
	games [10000000]game
}

func (h *yuge) addr() {
	fmt.Printf("%p\n", h)
}
