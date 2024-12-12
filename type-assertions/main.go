package main

import "fmt"

func main() {

	var mystery interface{} = "Hello world"

	str, ok := mystery.(string)

	if !ok {
		fmt.Println("err. not a string")
		return
	}

	fmt.Println(str)

	mystery = 6
	integer, ok := mystery.(string)

	if !ok {
		fmt.Println("err. not a string")
		return
	}

	fmt.Println(integer)

}
