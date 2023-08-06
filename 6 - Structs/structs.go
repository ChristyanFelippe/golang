package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	fmt.Println(u)
	u.nome = "Christyan"
	u.idade = 30
	fmt.Println(u)

	EnderecoExemplo := endereco{"Rua lobos", 0}

	usuario2 := usuario{"Christyan", 30, EnderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Christyan"}
	fmt.Println(usuario3)
}
