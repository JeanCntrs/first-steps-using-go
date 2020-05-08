package main

import "fmt"

func main() {
	fmt.Println(one(5))
	fmt.Println(variantOfOne(5))
	number, status := two(1)
	fmt.Println(number, status)
	fmt.Println(sum(2, 4))
	fmt.Println(sum(2, 4, 8))
	fmt.Println(sum(2, 4, 8, 16))
	fmt.Println(variantOfsum(5, 10, 15))
}

func one(number int) int {
	return number * 2
}

func variantOfOne(number int) (z int) {
	z = number * 3
	return
}

func two(number int) (int, bool) {
	if number == 1 {
		return 1, true
	}
	return 0, false
}

func sum(number ...int) int {
	total := 0
	for _, num := range number {
		total += num
	}
	return total
}

func variantOfsum(number ...int) int {
	total := 0
	for index, num := range number {
		total += num
		println(index)
	}
	return total
}
