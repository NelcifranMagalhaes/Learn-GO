package main

import "fmt"

type game struct {
	name             string
	price            float64
	developerCompany developer
}

type developer struct {
	name       string
	created_at string
}

type animal struct {
	height float32
	weight float32
}

type dog struct {
	animal
	quote string
	color string
}

func main() {

	dev1 := developer{"Capcom", "18/11/1991"}
	gameOne := game{"ultimate Ninja Storm", 258.75, dev1}

	fmt.Println(gameOne)

	animalOne := animal{25, 12}
	dogOne := dog{animalOne, "Au AU", "yellow"}
	fmt.Println(dogOne)
}
