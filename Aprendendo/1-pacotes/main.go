package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Printf("blabla")
	//Letras maiusculas e minusculas importante. Maiusculo = publico // minusculo -- private
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("christyanmend@gmail.com")
	fmt.Println(err) // retornando nil, se e-mail estiver correto

	// go get github.com/badoux/checkmail
}
