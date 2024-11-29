package main

import "fmt"

func main() {
	//with structs you have predefined data structures
	//maps are dynamic, are great for collections of values with custom values, that you can identify with a key
	//any value can be used in maps as a key, even structs

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon Web Services"])

	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	fmt.Println(websites["Linkedin"])

	delete(websites, "Google")
	fmt.Println(websites)
}
