package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {

	msg = "Hello, world!"

	messages := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	wg.Add(3)

	for _, msg := range messages {
		go updateMessage(msg, &wg)
	}

	wg.Wait()
}

func updateMessage(str string, wg *sync.WaitGroup) {
	defer wg.Done()

	msg = str

	printMessage()
}

func printMessage() {
	fmt.Println(msg)
}
