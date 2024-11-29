package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1)
	hobbies := [3]string{"Triathlon", "Reading", "Movies"}
	fmt.Println(hobbies)

	//2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5)
	courseGoals := []string{"Learn go", "Learn the basics"} //dynamic array
	fmt.Println(courseGoals)

	//6)
	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Learn all the basics")
	fmt.Println(courseGoals)

	//7)
	products := []Product{
		{
			"first",
			"socks",
			12.99,
		},
		{
			"second",
			"Jacket",
			50.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third",
		"shoes",
		15.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
