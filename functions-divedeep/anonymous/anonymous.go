package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2 // anonymous function, has to match the type in transformNumbers
	})

	fmt.Println(transformed)

	doubled := transformNumbers(&numbers, double)
	doubled2 := transformNumbers(&numbers, createTransformer(2))
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(doubled2)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
