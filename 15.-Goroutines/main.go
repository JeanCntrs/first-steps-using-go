package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go slowName("Jean C")
	fmt.Println("I'm here")
	var x string
	fmt.Scanln(&x)
	fmt.Println("Ejecución terminada:", x)
}

func slowName(name string) {
	letters := strings.Split(name, "")
	for index, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(index, letter)
	}
}
