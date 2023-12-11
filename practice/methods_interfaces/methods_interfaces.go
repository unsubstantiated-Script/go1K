package methods_interfaces

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

	store := list{
		&book{product{"moby dick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 150}},
	}

	store.discount(.5)
	store.print()
	//
	//var items []*game
	//items = append(items, &craftymine, &tootris)

	//my := List(items)
	//my = nil
	//my.Print()

}
