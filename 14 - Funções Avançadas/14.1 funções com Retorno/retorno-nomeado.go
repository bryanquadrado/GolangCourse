package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(5, 6)
	fmt.Println(soma, subtracao)

}