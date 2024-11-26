package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//testUser := user.User{
	//FirstName: "Pepe",
	//}

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		panic("cant continue")
	}

	admin := user.NewAdmin("test@example.com", "password123")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	fmt.Println("-------------------")

	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
