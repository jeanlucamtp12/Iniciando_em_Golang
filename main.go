package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hello, "
const espanhol = "espanhol"
const frances = "frances"
const ingles = "ingles"

func prefixoSaudacao(idioma string) (prefixo string) {

	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return

}

func Ola(nome string, idioma string) string {

	if nome == "" {
		nome = "mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
