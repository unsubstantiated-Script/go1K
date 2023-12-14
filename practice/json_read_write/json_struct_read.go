package json_read_write

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func JSONStructRead() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	//fmt.Println(string(input))

	var users []UserOut

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write stuffs.")
		}
		fmt.Println()
	}

}
