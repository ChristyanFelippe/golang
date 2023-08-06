package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(str string) string {
		fmt.Println(str)
		return str
	}

	fmt.Println(f("Mandei esta mensagem para a função"))

	resultadoSoma, resultadoSubstracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSubstracao)
	resultadoSoma1, _ := calculosMatematicos(30, 30) // recebe só o primeiro retorno, não precisa chamar o segundo parametro de retorno
	fmt.Println(resultadoSoma1)
	_, resultadoSubstracao1 := calculosMatematicos(10, 20)
	fmt.Println(resultadoSubstracao1)

}

func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
