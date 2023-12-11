package methods_interfaces

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	product
	published interface{}
}

// Print To the left is the receiver this attaches the method to the struct
func (b *book) print() {
	b.product.print()

	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}

	const layout = "2006/01"

	//Getting Unix time value
	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
