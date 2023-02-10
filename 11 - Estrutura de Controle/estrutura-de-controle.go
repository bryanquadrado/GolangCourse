package main

import "fmt"

func main() {
	x := -1

	if x > 10 {

		fmt.Println("Maior que 10")

	} else {
		fmt.Println("Menor que 10")
	}


	if y := x; y > 0 {
		fmt.Println("Maior que 0")
	} else if y < -10{
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Está em -10 e 0")
	}

	//fmt.Println(y)


}