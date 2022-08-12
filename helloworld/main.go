package main

import "fmt"

func main() {
	fmt.Println(bar()(42))
}

func bar() func(x int) int {
	return func(x int) int {
		return x
	}
}
