package main

import "fmt"

var calculate func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	calculate2 := func(num1 int, num2 int, num3 int) int {
		return num1 + num2 + num3
	}
	fmt.Printf("Suma: %d \n", calculate(4, 5))
	fmt.Printf("Suma: %d \n", calculate2(5, 5, 5))
	calculate = func(num1 int, num2 int) int {
		return num1 - num2
	}
	calculate2 = func(num1 int, num2 int, num3 int) int {
		return num1 - num2 - num3
	}
	fmt.Println("Resta:", calculate(10, 9))
	fmt.Println("Resta:", calculate2(10, 5, 3))
	calculate = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Println("División:", calculate(50, 2))
	operations()
	/* Closures */
	table2 := 2
	myTable := table(table2)
	for i := 1; i < 11; i++ {
		fmt.Println(myTable())
	}
}

func operations() {
	result := func() int {
		var a int = 1
		var b int = 2
		return a + b
	}
	fmt.Println("Operations:", result())
}

func table(value int) func() int {
	number := value
	sequence := 0
	return func() int { // La primera ejecucion de table lee todo el código de la función, luego solo toma en cuenta el segundo return. Sequence queda grabada la primera vez con 0, luego se incrementa en uno por cada iteración
		sequence++
		return number * sequence
	}
}
