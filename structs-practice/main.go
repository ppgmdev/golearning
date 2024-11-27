package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
)

type saver interface {
	Save() error // struct with a Save method and returns an error
}

type outputable interface {
	saver
	Display()
}

//type displayer interface {
//Display()
//}

//type outputable interface {
//saver
//displayer
//}

//type outputable interface {
//Save() error
//Display()
//}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println("saving the note file")
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		panic("can't continue")
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println("saving the note file")
		return
	}

	printSomething(todo) //not supported
}

func printSomething(value interface{}) { // accept any value
	typedVal, ok := value.(int)
	if ok {
		fmt.Println("Integer + 1: ", typedVal+1)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float + 1: ", floatVal+1)
		return
	}
	//switch value.(type) {
	//case int:
	//fmt.Println("Integer: ", value)
	//case float64:
	//fmt.Println("Float: ", value)
	//case string:
	//fmt.Println("String: ", value)
	//}
}

func outputData(data outputable) error {
	data.Display()
	err := saveData(data)
	return err
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note file")
		return err
	}

	fmt.Println("save succeeded")

	return nil
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') //use ' '

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
