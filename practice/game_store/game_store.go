package game_store

import (
	"bufio"
	"fmt"
	"os"
)

func GameStore() {

	games := []Game{
		{
			Item: Item{
				id:    1,
				name:  "god of war",
				price: 50,
			},
			genre: "action adventure",
		}, {
			Item: Item{
				id:    2,
				name:  "X Com 2",
				price: 30,
			},
			genre: "strategy",
		}, {
			Item: Item{
				id:    3,
				name:  "Minecraft",
				price: 20,
			},
			genre: "sandbox",
		},
	}

	printHeader(games)
	consoleOptions(games)

}

func printHeader(games []Game) {
	fmt.Printf("Justo's game store has %d games.\n", len(games))
}

func consoleOptions(games []Game) {
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> list: lists all the games
> quit: quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()
		printResults(games, *in)

		if "quit" == in.Text() {
			return
		}
	}
}

func printResults(games []Game, in bufio.Scanner) {
	switch in.Text() {
	case "quit":
		fmt.Println("bye!")
		return

	case "list":
		for _, g := range games {
			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		}
		return
	}
}
