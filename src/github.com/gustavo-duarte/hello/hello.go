package main

import "fmt"

//Evita de criar variaveis a todo momento!
const prefixHello = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrances = "Bonjour, "

func prefixSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "frances":
		prefixo = prefixHelloFrances
	case "spanish":
		prefixo = prefixHelloSpanish
	default:
		prefixo = prefixHello
	}
	return
}

func Hello(nome string, idioma string) string {
	if nome == "" {
		nome = "World"
	}
	return prefixSaudacao(idioma) + nome
}

func main() {
	fmt.Println(Hello("Django", ""))
}
