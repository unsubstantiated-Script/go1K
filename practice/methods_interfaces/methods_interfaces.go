package methods_interfaces

import "fmt"

func MakeDaMethodAndInterface() {
	//var (
	//	mobydick = book{
	//		title:     "moby dick",
	//		price:     10,
	//		published: 29123812321,
	//	}
	//
	//	derBit = book{
	//		title:     "Der Bit",
	//		price:     15,
	//		published: "1702244935",
	//	}
	//
	//	kakenStranz = book{
	//		title: "Kakenscranz",
	//		price: 15,
	//	}
	//
	//	craftymine = game{
	//		title: "crafty mine",
	//		price: 7.5,
	//	}
	//
	//	tootris = game{
	//		title: "tootris",
	//		price: 25,
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

	l := list{
		//If a field is not used, it doesn't have to exist here...
		{title: "moby dick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
		//{title: "rubik's cube", price: 5},
		//{title: "minecraft", price: 20},
		//{title: "tetris", price: 5},
		//{title: "yoda", price: 150},
	}

	l.discount(.5)
	//l.print()
	fmt.Print(l)
	//
	//var items []*game
	//items = append(items, &craftymine, &tootris)

	//my := List(items)
	//my = nil
	//my.Print()

}
