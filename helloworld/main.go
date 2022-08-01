package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8, 12}

	for i, v := range x {
		fmt.Println(i, v)
	}

}
