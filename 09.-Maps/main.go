package main

import "fmt"

func main() {
	countries := make(map[string]string, 5) // Tipo y capacidad
	fmt.Println(countries)

	countries["Mexico"] = "D.F."
	countries["Argentina"] = "Buenos Aires"
	println(countries["Mexico"])

	fmt.Println(countries)

	championship := map[string]int{"Barcelona": 5, "Real Madrid": 10, "Boca Juniors": 15}

	fmt.Println(championship) // Ordena alfabéticamente por clave

	championship["River Plate"] = 20 // Agregar nuevo valor al mapa
	championship["Real Madrid"] = 25 // Modificar valor por clave

	delete(championship, "Boca Juniors") // Eliminar valor por clave

	fmt.Println(championship)

	for team, score := range championship { // El range no ordena alfabéticamente
		fmt.Printf("\n El equipo %s tiene una puntuación de: %d", team, score)
	}

	score, exist := championship["Japan"]
	fmt.Printf("\nEl puntaje es: %d, y el equipo existe %t", score, exist)

	score2, exist := championship["River Plate"]
	fmt.Printf("\nEl puntaje es: %d, y el equipo existe %t", score2, exist)
}
