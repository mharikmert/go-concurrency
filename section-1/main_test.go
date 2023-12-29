package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printV2(t *testing.T) {

	// save the standard output to stdOut
	stdOut := os.Stdout

	// create standard output to test the actual stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	str := "string to be tested"
	go printV2(str, &wg)

	wg.Wait()

	// close the pipe
	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, str) {
		t.Errorf("Expected to find '%s', but not found", str)
	}

}
