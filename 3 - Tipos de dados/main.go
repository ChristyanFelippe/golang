package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int8 = -128
	fmt.Println(numero)
	var numero2 uint8 = 255
	fmt.Println(numero2)

	//alias
	//rune = int32
	//byte = uint8
	var numero3 rune = 12345
	fmt.Println(numero3)
	var numero4 byte = 255
	fmt.Println(numero4)

	// float
	var numero5 float32 = 12344.333
	fmt.Println(numero5)
	var numero6 float32 = 132332344.333
	fmt.Println(numero6)

	var string1 string = "fksdopfka"
	fmt.Println(string1)

	char := 'B'
	fmt.Println(char)
	// com aspas simples, ele usa como char, pegando o valor da tabela ascii

	// Fim de strings
	var texto1 string
	fmt.Println(texto1)

	var boolean bool
	fmt.Println(boolean)

	var err error
	fmt.Println(err)

	var err1 error = errors.New("Erro internal")
	fmt.Println(err1)
}
