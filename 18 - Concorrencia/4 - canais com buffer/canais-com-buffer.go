package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Olá Mundo 2 !"

	mensagem := <-canal
	fmt.Println(mensagem)
	mensagem2 := <-canal
	fmt.Println(mensagem2)

}