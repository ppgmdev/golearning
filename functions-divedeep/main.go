package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 2, 10, 16, 16) // starting value of 1
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int { // slice with intents with as many parameters as i want
	sum := 0
	fmt.Println("Starting value:", startingValue)
	for _, val := range numbers {
		sum += val
	}

	return sum
}
