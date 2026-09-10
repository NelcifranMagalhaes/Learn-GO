package main

import (
	"comand-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inicio.....!!")
	aplication := app.Gerar()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
