package main

import "fmt"

func main() {
	var temperaturaemC float64

	//Nessa primeira etapa, vamos solicitar a temperatura original em C, valendo que o programa ira executar de C pra F
	fmt.Println("Digite a temperatura")
	fmt.Scanln(temperaturaemC)

	//Algoritmo de Conversao
	tempEmF := (temperaturaemC - 32) * 5 / 9

	fmt.Println("A temperatura em F e = ", tempEmF)

}
