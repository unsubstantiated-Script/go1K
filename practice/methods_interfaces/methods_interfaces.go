package methods_interfaces

func MakeDaMethodAndInterface() {
	var (
		//mobydick = book{
		//	title: "moby dick",
		//	price: 10,
		//}

		//derBit = book{
		//	title: "Der Bit",
		//	price: 15,
		//}

		craftymine = game{
			title: "crafty mine",
			price: 7.5,
		}

		tootris = game{
			title: "tootris",
			price: 25,
		}
	)
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

	var items []*game
	items = append(items, &craftymine, &tootris)

	my := list(items)
	//my = nil
	my.print()

}
