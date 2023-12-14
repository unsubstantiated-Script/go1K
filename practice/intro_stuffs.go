package practice

import (
	"fmt"
	"os"
)

// FilePath testing out Split library method.
func IntroStuffs() {
	//Getting the File Path Jazz...
	//var dir, file string

	//Method has two returns.
	//dir, file = path.Split("css/main.css")

	//Underscore is an opt out of returns.
	//_, file = path.Split("css/main.css")
	//dir, _ = path.Split("css/main.css")

	//dir, file := path.Split("css/main.css")
	//
	//fmt.Println("dir:", dir)
	//fmt.Println("file:", file)

	//Type Conversion
	//speed := 100
	//force := 2.5
	//
	////Doing it this way ensures rounding doesn't jack up the results.
	//speed = int(float64(speed) * force)
	//
	//fmt.Println(speed)

	//Reading from the terminal.

	//Go compile and run setup
	// go build -o greeter.
	// Run './greeter' in the terminal. And you pass arguments to it.
	fmt.Printf("%#v\n", os.Args)
	//Grabbing the first argument of the Arg Slice.
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st Argument", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("3rd Argument:", os.Args[3])

}
