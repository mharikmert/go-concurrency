package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	message := "Updated message"

	var wg sync.WaitGroup

	wg.Add(1)

	go updateMessage(message, &wg)

	wg.Wait()

	if msg != message {
		t.Errorf("Expected message: %s, received message: %s", message, msg)
	}

}

func Test_printMessage(t *testing.T) {

}

func TestMain(t *testing.T) {
	// save the standard output to stdOut
	stdOut := os.Stdout

	// create standard output to test the actual stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(3)

	str := []string{
		"Hello, world!",
		"Hello, cosmos!",
		"Hello, universe!",
	}

	for _, msg := range str {
		go updateMessage(msg, &wg)
	}

	wg.Wait()

	// close the pipe
	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !(strings.Contains(output, str[0]) && strings.Contains(output, str[1]) && strings.Contains(output, str[2])) {
		t.Errorf("Expected to find all of messages, but found: %s", str)
	}
}
