// anonymous functions2
// Make me compile!

package main

import "fmt"

func main() {
	sayBye := func(n string) {
		fmt.Printf("Bye %s", n)
	}
	sayBye("zrgonomics")
}
