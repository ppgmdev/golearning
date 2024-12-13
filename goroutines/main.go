package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // makes sense if you know which operation will take the longest
}

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)
	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] = make(chan bool)
	go greet("how are you", done)
	//dones[2] = make(chan bool)
	go slowGreet("how ... are ... you ... ?", done)
	//dones[3] = make(chan bool)
	go greet("i hope you are doing well", done)

	//for _, done := range dones {
	//<-done
	//}

	for range done {
		// fmt.Println(doneChan)
	}

}
