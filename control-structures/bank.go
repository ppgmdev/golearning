package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFlowFromFile(accountBalanceFile)
	sayHi()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("Can't continue. Sorry")
	}
	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())
	for {

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return
				continue //continue with the next for loop iteration
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated. New amount: ", accountBalance)
		case 3:
			fmt.Print("Your amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdrawl amount. Must be greater than 0")
				return
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Can't be greater than balance amount")
				return
			}
			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated. New amount: ", accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}

		if choice == 1 {
		} else if choice == 2 {
		} else if choice == 3 {
		}
	}
	//fmt.Println("Thanks for choosing our bank")
}
