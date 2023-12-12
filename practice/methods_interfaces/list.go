package methods_interfaces

import (
	"strings"
)

// List Going to store item interface values here.
type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry we're waiting for stuffs.\n"

	}

	var str strings.Builder

	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}

}
