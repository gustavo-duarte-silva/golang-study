package main

import "fmt"

//Evita de criar variaveis a todo momento!
const prefixHello = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrances = "Bonjour, "

func Hello(nome string, idioma string) string {
	prefixo := prefixHello
	if nome == "" {
		nome = "World"
	}
	switch idioma {
	case "frances":
		prefixo = prefixHelloFrances
	case "spanish":
		prefixo = prefixHelloSpanish
	}
	return prefixo + nome
}

func main() {
	fmt.Println(Hello("Django", ""))
}
