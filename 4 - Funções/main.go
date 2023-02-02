package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMat(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 -n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func (txt string) string {
		fmt.Println(txt)
		return txt

	}

	resultado := "teste f1"
	fmt.Println(f(resultado))

	resultadoSoma, resultadoSub := calculosMat(10, 20)
	fmt.Println(resultadoSoma, resultadoSub)

	resultadoSoma2, _ := calculosMat(50, 20)
	fmt.Println(resultadoSoma2)

	

}