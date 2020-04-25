package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1, number2, result int
var legend string

func main() {
	fmt.Println("Ingrese número 1:")
	//fmt.Scanf("%d", &number1) // Problemas en windows, esta pensado para linux
	fmt.Scanln(&number1) // puntero

	fmt.Println("Ingrese número 2:")
	//fmt.Scanf("%d", &number2)
	fmt.Scanln(&number2)
	fmt.Println("La suma es:", number1+number2)

	fmt.Println("Descripción:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		legend = scanner.Text()
	}

	result = number1 + number2
	fmt.Println(legend, result)
}
