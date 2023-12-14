package methods_interfaces

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `[
 {
  "title": "moby dick",
  "price": 10,
  "released": 118281600
 },
 {
  "title": "odyssey",
  "price": 15,
  "released": 733622400
 },
 {
  "title": "hobbit",
  "price": 25,
  "released": -62135596800
 }
]`

func MakeDaMethodAndInterface() {
	//var (
	//	mobydick = book{
	//		Title:     "moby dick",
	//		Price:     10,
	//		published: 29123812321,
	//	}
	//
	//	derBit = book{
	//		Title:     "Der Bit",
	//		Price:     15,
	//		published: "1702244935",
	//	}
	//
	//	kakenStranz = book{
	//		Title: "Kakenscranz",
	//		Price: 15,
	//	}
	//
	//	craftymine = game{
	//		Title: "crafty mine",
	//		Price: 7.5,
	//	}
	//
	//	tootris = game{
	//		Title: "tootris",
	//		Price: 25,
	//	}
	//)

	//
	////Calling like a method
	//mobydick.print()
	//derBit.print()
	//craftymine.discount(.5)
	//craftymine.print()
	//
	//var h yuge
	//
	//for i := 0; i < 10; i++ {
	//	h.addr()
	//}

	//l := list{
	//	//If a field is not used, it doesn't have to exist here...
	//	{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
	//	{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
	//	{Title: "hobbit", Price: 25},
	//	//{Title: "rubik's cube", Price: 5},
	//	//{Title: "minecraft", Price: 20},
	//	//{Title: "tetris", Price: 5},
	//	//{Title: "yoda", Price: 150},
	//}

	//l.discount(.5)

	//sort.Sort(sort.Reverse(l))
	//
	//sort.Sort(byReleaseDate(l))
	//
	////l.print()
	//fmt.Print(l)
	//
	//var items []*game
	//items = append(items, &craftymine, &tootris)

	//my := List(items)
	//my = nil
	//my.Print()

	//data, err := json.MarshalIndent(l, "", " ")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(data))

	var l list
	if err := json.Unmarshal([]byte(data), &l); err != nil {
		log.Fatal(err)
	}
	fmt.Print(l)

}
