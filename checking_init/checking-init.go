package main

import "fmt"
import "abc"

func init() {
	fmt.Println("Main waala init.")
}

func main() {
	abc.Foo()
}
