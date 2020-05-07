package main

import "fmt"

func main() {
	i := 1
	for i < 4 {
		fmt.Println(i)
		i++
	}

	for x := 1; x < 4; x++ {
		fmt.Println(x)
	}

	number := 0
	for {
		fmt.Println("Continuo", number)
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&number) // Con este puntero ingreso números que se alojan en la variable número
		if number == 100 {
			break // Rompe el ciclo
		}
	}

	var y = 0
	for y < 10 {
		fmt.Printf("\n valor de y: %d", y)
		if y == 5 {
			fmt.Printf("   Multimplicamos por 2 \n")
			y = y * 2
			continue // Manda el control de la ejecución nuevamente al inicio del ciclo for, no continua ejecutando el resto del código
		}
		fmt.Printf("   Pasó por aquí \n")
		y++
	}

	var p int = 0
RUTINA:
	for p < 10 {
		if p == 4 {
			p = p + 2
			fmt.Println("Voy a RUTIN sumando 2 a p")
			goto RUTINA // La ejecución va hacia donde nosotros le digamos, en este caso es RUTINA
		}
		fmt.Printf("Valor de p: %d \n", p)
		p++
	}
}
