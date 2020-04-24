package main

import (
	"fmt"
	"strconv"
)

var number int
var text string
var status bool
var status2 bool = true

func main() {
	var number3, number4, number5 int
	number6, number7, number8, text2, status := 2, 3, 4, "text2", true
	text3 := "text3"
	text4 := "text4"

	var coin float32 = 0
	number6 = int(coin)
	text3 = fmt.Sprintf("%f", coin) // convierte todo a texto, pero es necesario especificar el tipo de conversion
	text4 = strconv.Itoa(int(coin))

	fmt.Println(number3, number4, number5)
	fmt.Println(number6, number7, number8, text2)
	fmt.Println(number, text, status, status2)
	fmt.Println(text3, text4)

	showStatus()
}

func showStatus() {
	fmt.Println(status)
}
