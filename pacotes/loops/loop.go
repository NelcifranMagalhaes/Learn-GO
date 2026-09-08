package main

import (
	"fmt"
	"time"
)

func loopOne() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)
}

func loopTwo() {
	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
}

func loopThree() {
	names := [3]string{"Sasuke", "Drastrea", "DD crow"}

	for indice, name := range names {
		fmt.Println(indice, name)
	}

}

func loopMap() {
	user := map[string]string{
		"name":     "Midoria",
		"lastName": "Izuku",
	}

	for chave, valor := range user {
		fmt.Println(chave, valor)
	}
}

func main() {
	loopMap()
}
