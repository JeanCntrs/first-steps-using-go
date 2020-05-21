package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan time.Duration)
	go loop(channel)
	fmt.Println("Legué hasta aquí")

	// Espera a que el buble devuelva un valor en el canal para terminar, similar al await en node
	// Espera a que canal tenga un valor y lo grabe en la variable msg antes de terminar la ejecución del programa
	// Hasta que esto no ocurra, el programa no continua
	msg := <-channel
	fmt.Println(msg)
}

func loop(channel chan time.Duration) {
	start := time.Now()
	for i := 0; i < 10000000000; i++ {

	}
	end := time.Now()

	// Asigna el valor resultante de la función Sub
	channel <- end.Sub(start)
}
