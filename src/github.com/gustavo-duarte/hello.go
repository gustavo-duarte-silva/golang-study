package main

import "fmt"

func Hello(nome string) string {
	return "Olá, " + nome
}

func main() {
	fmt.Println(Hello("mundo"))
}
