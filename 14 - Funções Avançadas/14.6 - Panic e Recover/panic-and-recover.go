package main

import "fmt"

func tentandoRecuperar() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunaEstaAprovado(n1, n2 float32) bool {

	defer tentandoRecuperar()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media é exatemente 6")
}

func main() {

	fmt.Println(alunaEstaAprovado(6, 6))
	fmt.Println("Após execução")

}