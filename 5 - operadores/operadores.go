package main

import "fmt"

func main() {

	soma := 2 + 2
	sub := 2 - 2
	div := 2 / 2
	mult := 2 * 2
	resto := 2 % 2
	fmt.Println(soma, sub, div, mult, resto)

	verdadeiro := true
	falso := false
	fmt.Println(verdadeiro, falso)
	fmt.Println("-----------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println("-----------")
	num := 10
	num2 := num + 1
	fmt.Println(num2)
	num++
	fmt.Println(num)
	num += 15
	fmt.Println(num)

	num--
	fmt.Println(num)
	num -= 15
	fmt.Println(num)

	// operador ternario
	// não existe em golang.
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"

	var texto string
	if num > 5 {
		texto = "num maior que 5"
	} else {
		texto = "num menor que 5"
	}
	fmt.Println(texto)
}
