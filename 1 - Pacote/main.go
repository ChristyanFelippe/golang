package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	auxiliar.SegundoEscrever()
	retorno := checkmail.ValidateFormat("christyanmend@gmail.com")
	fmt.Println(retorno)

}
