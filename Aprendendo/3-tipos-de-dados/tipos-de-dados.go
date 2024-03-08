package main

import (
	"errors"
	"fmt"
)

func main() {
	// outros exemplos int16, int32 int64
	var numero1 int8 = 127 // valor maximo
	var numero2 int8 = -128
	fmt.Println(numero1)
	fmt.Println(numero2)
	// outros exemplos uint16, uint32 uint64
	// uint = somente valores positivos
	var num3 uint8 = 0
	var num4 uint8 = 255
	fmt.Println(num3)
	fmt.Println(num4)

	//rune
	// rune = INT32
	var num5 rune = 10000
	fmt.Println(num5)

	// byte
	// Byte = uint8
	var num6 byte = 123
	fmt.Println(num6)

	//float32, float64
	var num7 float32 = 323232.444
	fmt.Println(num7)

	var texto1 string = "texto1"
	fmt.Println(texto1)

	texto2 := "texto2"
	fmt.Println(texto2)

	char := 'F'
	fmt.Println(char) // retorno é o valor em tabela ascii

	var zero uint16
	fmt.Println(zero)

	bool1 := true
	bool2 := false
	fmt.Println(bool1, bool2)

	var error error = errors.New("Erro interno")
	fmt.Println(error)
}
