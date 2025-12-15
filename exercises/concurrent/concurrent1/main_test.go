// concurrent1
// Make the tests pass!

package main_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

func TestPrinter(t *testing.T) {
	var buf bytes.Buffer
	print(&buf)

	out := buf.String()

	for i := 0; i < 3; i++ {
		want := fmt.Sprintf("Hello from goroutine %d!", i)
		if !bytes.Contains([]byte(out), []byte(want)) {
			t.Errorf("Output did not contain expected string. Wanted: %q, Got: %q", want, out)
		}
	}
}

func print(buf *bytes.Buffer) {
	var wg sync.WaitGroup // A counter of goroutines
	var mu sync.Mutex     // Mutual Exclusion lock

	goroutines := 3

	// For each goroutine...
	for i := 0; i < goroutines; i++ {
		// Add 1 to the WaitGroup counter
		wg.Add(1)
		// Execute goroutine
		go func(i int) {
			// Decrement WaitGroup by one (equivalent to wg.Add(-1))
			defer wg.Done()
			// Lock the mutex
			mu.Lock()
			// Do thing
			fmt.Fprintf(buf, "Hello from goroutine %d!\n", i)
			// Unlock mutex
			mu.Unlock()
		}(i)
	}
	// Wait until WaitGroup counter is 0
	wg.Wait()
}
