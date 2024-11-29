package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// go will create an array with length 2, max 5  items, will create an array with 2 empty slots
	//you can then assign elements to this slot, useful to assign values to specific indexes
	// useful if you know in advance the approx list of items, could be a good optimization
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Jose"

	userNames = append(userNames, "Pepe")
	userNames = append(userNames, "Antonio")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 3.7
	courseRatings["angular"] = 4.3

	courseRatings.output()

	//fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
