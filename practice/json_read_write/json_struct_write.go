package json_read_write

import (
	"encoding/json"
	"fmt"
)

func JSONStructWrite() {
	users := []UserIn{
		{"Derp", "1234", nil},
		{"Bilbo", "42", Permissions{"admin": true}},
		{"Cattor", "73", Permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(string(out))
}
