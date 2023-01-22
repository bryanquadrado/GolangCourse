package main

import "fmt"

func main() {

	var variavel1 string = "variavel1"
	variavel2 := "variavel2"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5","variavel6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "conatante1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}