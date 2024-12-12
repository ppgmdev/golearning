package main

import "fmt"

func main() {
	myStr := "hello bedrock"
	myStrPointer := &myStr

	var mystery interface{} = myStrPointer

	str, ok := mystery.(*string)
	fmt.Println(ok)

	if !ok {
		fmt.Println("err. not a string")
	}

	fmt.Println("Str:", str)

	mystery = 6
	integer, ok := mystery.(string)

	if !ok {
		fmt.Println("err. not a string")
	}

	fmt.Println(integer)

}
