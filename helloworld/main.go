package main

import "fmt"

func main() {
	fmt.Println(bar()(42))
	fmt.Println("we added something to test")
}

func bar() func(x int) int {
	return func(x int) int {
		return x
	}
}
