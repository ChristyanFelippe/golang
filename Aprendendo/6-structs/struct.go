package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua            string
	numero_da_casa uint8
}

func main() {
	var u usuario
	u.nome = "Christyan"
	u.idade = 31
	fmt.Println(u.nome, u.idade)
	fmt.Println(u)

	endereco_exemplo := endereco{"rua dos lobos", 33}

	usuario2 := usuario{"maria", 22, endereco_exemplo}
	fmt.Println(usuario2)

	usuario3_faltando_parametro := usuario{nome: "cabrito"}
	fmt.Println(usuario3_faltando_parametro)

	usuario3_faltando_parametro_idade := usuario{idade: 33}
	fmt.Println(usuario3_faltando_parametro_idade)
}
