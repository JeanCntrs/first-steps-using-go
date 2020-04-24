package main

import "fmt"

var status bool

func main() {
	status = true
	if status = false; status == true {
		fmt.Println(status)
	} else {
		fmt.Println("Es falso")
	}

	switch number := 5; number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")

	}
}
