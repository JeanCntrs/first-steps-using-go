package main

import "fmt"

// Middlewares: Son interceptores que permiten ejecutar instrucciones a varias funciones que reciben y devuelven los mismos tipos de variables

var result int

func main() {
	fmt.Println("Inicio")
	result = operationsMiddleware(sum)(2, 3)
	fmt.Println(result)
	result = operationsMiddleware(subtract)(10, 6)
	fmt.Println(result)
	result = operationsMiddleware(multiply)(2, 4)
	fmt.Println(result)
}

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func operationsMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		// Aquí es donde se realiza la operación que tendrán en comun las funciones con este middleware
		fmt.Println("Inicio de la Operación")
		return f(a, b)
	}
}
