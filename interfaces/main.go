package main

import "fmt"

type Talker interface {
	SayHello()
}

type Person struct {
}

func (person Person) SayHello() {
	fmt.Println("Hello! I am a person")
}

type Robot struct {
}

func (robot Robot) SayHello() {
	fmt.Println("Hello! I am a robot")
}

func main() {
	person := Person{}
	robot := Robot{}

	var talkerPerson Talker = person
	talkerPerson.SayHello()

	var talkerRobot Talker = robot
	talkerRobot.SayHello()
}
