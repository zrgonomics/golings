// structs1
// Make me compile!
package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

var person = Person{"zrgonomics", 99}

func main() {
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
