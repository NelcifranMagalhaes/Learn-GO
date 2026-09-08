package main

import "fmt"

func main() {
	var arrayOne [5]string
	arrayOne[0] = "Position 1"
	fmt.Println(arrayOne)

	arrayTwo := [...]int{1, 2, 3, 4}
	fmt.Println(arrayTwo)

	slice := []int{1, 1, 2, 7, 8, 9, 7}
	slice = append(slice, 58)
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
