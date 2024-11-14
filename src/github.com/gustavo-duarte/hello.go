package main

import "fmt"

//Evita de criar variaveis a todo momento!
const prefixHello = "Hello, "

func Hello(nome string) string {
	if nome == "" {
		nome = "World"
	}
	return prefixHello + nome
}

func main() {
	fmt.Println(Hello("mundo"))
}
