package main

import "fmt"

func main() {
	fmt.Println("test")
	foo()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("testagain")
	}
}
