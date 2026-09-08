package auxiliar

import (
	"errors"
	"fmt"
)

func Escrever() {
	mensagem := "Escrevendo Mensagem!!"
	fmt.Println(mensagem)

	var numberx = 54
	fmt.Println(numberx)

	var stringNova string = "String Nova"
	fmt.Println(stringNova)

	var booleano bool
	fmt.Println(booleano)

	var variable_erro error = errors.New("Erro encontrado!!")
	fmt.Println(variable_erro)

}
