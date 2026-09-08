package main

import "fmt"

func soma(numbers ...int) int {

	total := 0

	for _, numero := range numbers {
		total += numero
	}
	return total
}

func main() {

	sumTotal := soma(1, 5, 12)
	fmt.Println(sumTotal)
}
