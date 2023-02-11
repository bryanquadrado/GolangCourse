package main

import "fmt"

func diasDasemana(numero int) string {
	switch numero {
	case 1: 
	    return "Domingo"
	case 2: 
		return "Segunda-Feira"
	case 3: 
		return "Terça-Feira"
	case 4: 
		return "Quarta-Feira"
	case 5: 
		return "Quinta-Feira"
	case 6: 
		return "Sexta-Feira"
	case 7: 
		return "Sabádo"
	default:
	    return "numero invalido"
	} 

	//return "Numero invalido"

}

func diasDasemana2(numero int) string {
	var diasDasemana string
	switch  {
	case numero == 1: 
	     diasDasemana = "Domingo"
		 fallthrough
	case numero == 2: 
		 diasDasemana = "Segunda-Feira"
	case numero == 3: 
		 diasDasemana = "Terça-Feira"
	case numero == 4: 
		 diasDasemana = "Quarta-Feira"
	case numero == 5: 
		 diasDasemana = "Quinta-Feira"
	case numero == 6: 
		 diasDasemana = "Sexta-Feira"
	case numero == 7: 
		 diasDasemana = "Sabádo"
	default:
	     diasDasemana = "numero invalido"
	} 

	return diasDasemana

}

func main() {

	fmt.Println("Switch")

	x := diasDasemana(12)
	fmt.Println(x)

	y := diasDasemana2(1)
	fmt.Println(y)

}