package game_store

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	consoleOptions(games, makeMap(games))
}

func printHeader(games []Game) {
	fmt.Printf("Justo's game store has %d games.\n", len(games))
}

func makeMap(games []Game) map[int]Game {
	byID := make(map[int]Game)

	for _, g := range games {
		byID[g.id] = g
	}

	return byID
}

func consoleOptions(games []Game, byID map[int]Game) {
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> id: game by id
> list: lists all the games
> quit: quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {
			continue
		}

		printResults(games, cmd, byID)

		if "quit" == in.Text() {
			return
		}
	}
}

func printResults(games []Game, cmd []string, byID map[int]Game) {
	switch cmd[0] {
	case "quit":
		fmt.Println("bye!")
		return
	case "list":
		listGames(games)
		return
	case "id":
		retrieveGame(cmd, byID)
		return
	}
}

func listGames(games []Game) {
	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}
}

func retrieveGame(cmd []string, byID map[int]Game) {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return
	}
	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return
	}

	g, ok := byID[id]

	if !ok {
		fmt.Println("sorry. This game doesn't exist")
		return
	}

	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
