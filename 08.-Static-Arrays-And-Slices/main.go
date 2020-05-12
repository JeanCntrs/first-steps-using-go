package main

import "fmt"

var arr [10]int      // Presente en todas las funciones por el scope global
var matrix [5][7]int // Conjunto de vectores

func main() {
	arr[1] = 1
	arr[5] = 15
	fmt.Println(arr)

	arr2 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 88}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(i)
	}

	matrix[3][5] = 1
	fmt.Println(matrix)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	sliceVariant()
	sliceVariant2()
	sliceVariant3()
}

func sliceVariant() {
	elements := [5]int{2, 4, 6, 8, 10}
	portion := elements[:3]
	portion2 := elements[3:]
	fmt.Println(portion, portion2)
}

func sliceVariant2() {
	elements := make([]int, 5, 20) // Tipo, largo, capacidad
	fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))
}

func sliceVariant3() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i) // Usar append solo cuando me quede sin espacio, tiene su costo en proceso
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
