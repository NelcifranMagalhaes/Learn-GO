package main

import "fmt"

func somar(numberOne int8, numberTwo int8) int8 {
	return numberOne + numberTwo
}

func somaAndSubtra(numberOne, numberTwo int8) (int8, int8) {
	soma := numberOne + numberTwo
	subtraçao := numberOne - numberTwo
	return soma, subtraçao
}

func main() {
	resultado := somar(10, 20)
	println(resultado)

	var funcao = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fmt.Printf(funcao("Midoria\n"))

	resultSoma, ResultSubtra := somaAndSubtra(15, 5)
	fmt.Println(resultSoma, ResultSubtra)
}
