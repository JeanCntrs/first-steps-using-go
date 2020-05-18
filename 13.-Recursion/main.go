package main

import "fmt"

func main() {
	recursion(2)
}

func recursion(number int) {
	if number > 100000000 {
		return
	}
	fmt.Println(number)
	recursion(number * 2)
}
