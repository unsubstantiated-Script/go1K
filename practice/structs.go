package practice

//type person struct {
//	name, lastname string
//	age            int
//}

//type playlist struct {
//	genre string
//	songs []song
//}
//
//type song struct {
//	title, artist string
//}

type text struct {
	title string
	words int
}

type book struct {
	text
	isbn string
}

func StructsStuff() {
	//var picasso person
	//
	//picasso.name, picasso.lastname, picasso.age = "Pablo", "Picasso", 91
	//
	//fmt.Printf("\n Picasso: %v\n", picasso)

	//song1 := song{title: "wonderwall", artist: "oasis"}
	//song2 := song{title: "super sonic", artist: "oasis"}

	//Adding one struct to another via slices

	//songs := []song{
	//	{title: "wonderwall", artist: "oasis"},
	//	{title: "super sonic", artist: "oasis"},
	//}
	//
	//rock := playlist{
	//	genre: "indie rock",
	//	songs: songs,
	//}
	//
	//song := rock.songs[0]
	//song.title = "What I want..."
	//
	//fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	//for _, s := range rock.songs {
	//	fmt.Printf("%-20s %20s\n", s.title, s.artist)
	//}

	//moby := book{
	//	text: text{
	//		title: "Moby Dick",
	//		words: 206502,
	//	},
	//	isbn: "102030",
	//}
	//
	//fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

}
