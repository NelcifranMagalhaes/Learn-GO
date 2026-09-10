package main

import "fmt"

type user struct {
	name string
	age  int8
}

func (u user) salvar() {
	fmt.Printf("salvando o user %s \n", u.name)
}

func (u user) ageVerify() bool {
	return u.age >= 18
}

func main() {

	user := user{"primeiro cara", 15}
	fmt.Println(user)
	user.salvar()
	fmt.Println(user.ageVerify())
}
