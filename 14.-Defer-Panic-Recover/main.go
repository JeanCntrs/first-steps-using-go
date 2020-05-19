package main

import "log"

func main() {
	/* file := "test.txt"
	f, err := os.Open(file)

	// Se ejecuta en error o fin de programa
	defer f.Close()

	if err != nil {
		fmt.Println("Error al abrir el archivo")
		os.Exit(1)
	} */
	panicExample()
}

func panicExample() {
	defer func() {
		// Trae el resultado del último panic, de no haber, el resultado será nulo
		reco := recover()
		if reco != nil {
			log.Fatalf("Hubo un error que generó Panic \n %v", reco)
		}
	}()
	number := 1
	if number == 1 {
		// Fuerza un error
		panic("Se encontro el valor de 1")
	}
}
