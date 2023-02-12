package main

import (
	"fmt"
	//"time"
)

func main() {

	/* i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando o i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2{
		fmt.Println("Incrementando o j", j)
		time.Sleep(time.Second)
	}
 */

	nomes := [3]string {"Braia","Caio","Bruna"} 	

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome" : "Braia",
		"sobrenome" : "Quadrado",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}

	//range cannot be used in structs
	
}