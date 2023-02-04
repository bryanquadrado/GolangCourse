package main

import "fmt"

type usuario struct {
	nome, sobrenome  string
	idade, altura    uint8
}

type estudante struct {
	usuario
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Inheritance")

	p1 := usuario{"Bryan", "Quadrado", 26, 171}
	fmt.Println(p1)

	u1 := estudante{p1, "Computer Enginner", "UNP"}
	fmt.Println(u1)
	fmt.Println(u1.sobrenome)

}