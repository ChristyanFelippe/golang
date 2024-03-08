package main

import "fmt"

func main() {
	var valor_a_ser_adicionaro uint = 30
	meuArray := [5]int{10, 20, 30, 40, 50}
	fmt.Println(meuArray, len(meuArray))
	var novoArray [len(meuArray) + 1]int
	fmt.Println(len(novoArray))
	for i := 0; i < len(meuArray); i++ {
		// novoArray[i] = meuArray[i]
		copy(novoArray[:], meuArray[:])
	}
	fmt.Println("array atualizado ", novoArray)
	novoArray[len(novoArray)-1] = int(valor_a_ser_adicionaro)
	fmt.Println("final ", novoArray)

}
