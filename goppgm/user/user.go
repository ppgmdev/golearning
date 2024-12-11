package user

import (
	"encoding/json"
	"fmt"
)

type favSongsType map[string]string

type user struct {
	Name     string       `json:"NAME"`
	Age      int          `json:"AGE"`
	Country  string       `json:"COUNTRY"`
	Hobbies  *[]string    `json:"HOBBIES"`
	FavSongs favSongsType `json:"SONGS"`
}

func (u user) PrintRequest() {
	fmt.Println(u)
}

func (u user) GetFavSongs() {
	fmt.Println("Favorite Songs:")
	for artist, song := range u.FavSongs {
		fmt.Println(artist, song)
	}
}

func New(name, country string, age int, hobbies *[]string, favSongs favSongsType) (user, []byte) {
	user := user{
		Name:     name,
		Country:  country,
		Age:      age,
		Hobbies:  hobbies,
		FavSongs: favSongs,
	}

	json, _ := json.Marshal(user)

	return user, json
}
