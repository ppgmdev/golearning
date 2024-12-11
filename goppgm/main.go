package main

import (
	"fmt"

	"example.com/me/user"
)

func main() {
	myHobbies := []string{"triathlon", "movies", "friends", "piano"}
	//myHobbies := make([]string, 3, 4)
	//myHobbies = append(myHobbies, "triathlon")
	//myHobbies = append(myHobbies, "movies")
	//myHobbies = append(myHobbies, "friends")
	//myHobbies = append(myHobbies, "piano")

	favSongs := map[string]string{
		"ColdPlay":      "Magic",
		"Phil Collins":  "Another day in paradise",
		"Justin Bieber": "Ghost",
	}

	user, json := user.New("Pepe", "Mex", 26, &myHobbies, favSongs)

	fmt.Println("JSON: ", string(json))

	user.PrintRequest()
	fmt.Println(user)

	user.GetFavSongs()

}
