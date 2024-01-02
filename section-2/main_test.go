package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	os.Stdout = stdOut

	result, _ := io.ReadAll(r)
	output := string(result)

	if !strings.Contains(output, "13920") {
		t.Error("Unexpected result returned")
	}

}
