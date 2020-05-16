package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readTextFile()
	readTextFile2()
	writeFileText()
	writeFileText2()
}

func readTextFile() {
	file, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error:", err)
	} else {
		fmt.Println(string(file))
	}
}

// Brinda más libertad al leer línea por línea y realizar operaciones en base a eso
func readTextFile2() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error:", err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			registry := scanner.Text()
			fmt.Printf("Línea > " + registry + "\n")
		}
	}
	file.Close()
}

func writeFileText() {
	file, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Println("Hubo un error:", err)
		return
	}
	fmt.Fprintln(file, "This is a new line")
	file.Close()
}

func writeFileText2() {
	filename := "./newFile.txt"
	if append(filename, "This is a second line") == false {
		fmt.Println("Error en línea dos")
	}
}

func append(filename string, text string) bool {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error:", err)
		return false
	}
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Hubo un error:", err)
		return false
	}
	return true
}
