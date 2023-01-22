package main

import (
	"module/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("bryan@gmail.com")
	fmt.Println(erro)

}