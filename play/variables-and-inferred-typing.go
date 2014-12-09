package main

import "fmt"

var (
	name     string
	age      int
	location string
)

func main() {
	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)
}
