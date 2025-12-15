// concurrent3
// Make the tests pass!

package main_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSendAndReceive(t *testing.T) {
	var buf bytes.Buffer

	messages := make(chan string)
	sendAndReceive(&buf, messages)

	got := buf.String()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func sendAndReceive(buf *bytes.Buffer, messages chan string) {
	go func() {
		messages <- "Hello"
		messages <- "World"
		close(messages)
	}()

	first := true
	for msg := range messages {
		// If we're not on the first msg, prepend a space to each msg in the range
		if !first {
			fmt.Fprint(buf, " ")
		}
		fmt.Fprint(buf, msg)
		first = false
	}
}
