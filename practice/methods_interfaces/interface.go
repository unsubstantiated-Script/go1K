package practice

import "go1K/practice/methods_interfaces"

func InterfacePractice() {
	var (
		craftymine = methods_interfaces.Game{
			Title: "crafty mine",
			Price: 7.5,
		}

		tootris = methods_interfaces.Game{
			Title: "tootris",
			Price: 25,
		}
		deezmoby = methods_interfaces.Book{
			Title: "Deez Moby",
			Price: 15,
		}
		loxx = methods_interfaces.Puzzle{
			Title: "Loxx",
			Price: 25,
		}
	)

	var store methods_interfaces.List

	store = append(store, &craftymine, &tootris, &deezmoby, &loxx)
	store.Print()
}
