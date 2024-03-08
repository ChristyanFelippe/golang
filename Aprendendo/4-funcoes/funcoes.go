package main

import (
	"fmt"
)

// func <nome da função> (<parametros passados para funçaõ>) (<o que a função vai retornar>)
//
//	{
//		dados da função
//	}
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	return n1 + n2, n1 * n2
}

func main() {
	fmt.Println(somar(2, 3))

	var f = func(txt string) string {
		fmt.Println(txt)
		return "resultado da função f"
	}

	resultado := f("texto repassado a função f")
	fmt.Println(resultado)
	resultado1, _ := calculosMatematicos(2, 5)
	fmt.Println(resultado1)
	_, resultado2 := calculosMatematicos(2, 5)
	fmt.Println(resultado2)
	resultado3, resultado4 := calculosMatematicos(2, 5)
	fmt.Println(resultado3, resultado4)

}
