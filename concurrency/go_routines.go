package concurrency

import (
	"fmt"
	"time"
)

var start time.Time

func getTimeSinceStart() string {
	return fmt.Sprintf("- time(%s)", time.Since(start).String())
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func StartGoRoutines() {
	start = time.Now()
	done := make(chan bool)

	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	// fmt.Println("Waiting for the slow greeting...", <-done)

	for range done {
		fmt.Println("Done!", getTimeSinceStart())
	}
}
