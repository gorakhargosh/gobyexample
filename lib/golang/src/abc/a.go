package abc

import "fmt"

func init() {
	fmt.Println("a wala init.")
}

func Foo() {
	fmt.Println("abc.Foo a walla")
}

func main() {
	fmt.Println("Main cannot be here.")
}
