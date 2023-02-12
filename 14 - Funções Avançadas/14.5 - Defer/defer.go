package main

import "fmt"

func execucao1() {
	fmt.Println("Executando a funcao 1")
}

func execucao2() {
	fmt.Println("Executando a funcao 2")
}

func alunaEstaAprovado(n1, n2 float32) bool {

	defer fmt.Println("Sua situação é")

	fmt.Println("Verificando se o aluno esta aprovado")
	media := (n1 + n2) / 2

	if media >= 6{
		return true
	}

	return false
}

func main() {

	defer execucao1()
	execucao2()
	fmt.Println(alunaEstaAprovado(8, 7))

}