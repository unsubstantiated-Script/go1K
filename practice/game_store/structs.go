package game_store

type Item struct {
	id    int
	name  string
	price int
}

type Game struct {
	Item
	genre string
}
