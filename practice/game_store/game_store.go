package game_store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GameStore Main Program that handles the games.
func GameStore() {
	games := makeGamesSlice(decodeGames())
	gamesMap := makeMap(games)

	printHeader(games)

	consoleOptions(games, gamesMap)
}

// Printing Console Header
func printHeader(games []Game) {
	fmt.Printf("Justo's game store has %d games.\n", len(games))
}

// Decoding the games from JSON
func decodeGames() []JSONGame {
	var decoded []JSONGame
	if err := json.Unmarshal([]byte(Data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
	}
	return decoded
}

// Builds a slice of game Structs
func makeGamesSlice(decoded []JSONGame) []Game {
	var games []Game
	for _, dg := range decoded {
		games = append(games, Game{Item{dg.ID, dg.Name, dg.Price}, dg.Genre})
	}
	return games
}

// Makes a Mapped array of the games.
func makeMap(games []Game) map[int]Game {
	byID := make(map[int]Game)

	for _, g := range games {
		byID[g.id] = g
	}

	return byID
}

// Prints and services the console options.
func consoleOptions(games []Game, byID map[int]Game) {
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> id N: game by id
> list: lists all the games
> save: exports the game to json and quits
> quit: quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		//Trimming out whitespace.
		res := strings.TrimSpace(in.Text())

		//Loading up fields from the CLI scan in.
		cmd := strings.Fields(res)

		if len(cmd) == 0 {
			continue
		}

		processInput(games, cmd, byID)

		if "quit" == res || "save" == res {
			return
		}
	}
}

// Processes the User Input
func processInput(games []Game, cmd []string, byID map[int]Game) {
	switch cmd[0] {
	case "quit":
		fmt.Println("bye!")
		return
	case "save":
		saveFile(games)
		fmt.Println("File Saved Bye!")
		return
	case "list":
		listGames(games)
		return
	case "id":
		retrieveGame(cmd, byID)
		return
	}
}

// Prints out a list of the games
func listGames(games []Game) {
	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}
}

// Retrieves a game by ID
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

// Saves the file to a JSON Object.
func saveFile(games []Game) {
	var encodable []JSONGame

	for _, g := range games {
		encodable = append(encodable,
			JSONGame{
				g.id,
				g.name,
				g.genre,
				g.price,
			})
	}

	out, err := json.MarshalIndent(encodable, "", "\t")
	if err != nil {
		fmt.Println("Sorry: ", err)
		return
	}

	fmt.Println(string(out))
	return
}
