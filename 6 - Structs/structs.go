package main

import "fmt"

type usuario1 struct {
    nome string
	idade uint8
	endereco
}

type endereco struct {
	rua string
	numero uint8
}

func main() {
	fmt.Println("Structs")
	var u usuario1
	u.nome = "Bryan"
	u.idade = 26
	fmt.Println(u)

	enderecoEx := endereco{"Rua do braiia", 69}

	usuario2 := usuario1{"Bryan", 26, enderecoEx}
	fmt.Println(usuario2)

	// pass only one value from struct

	usuario3 := usuario1{idade:26}
	fmt.Println(usuario3)

}