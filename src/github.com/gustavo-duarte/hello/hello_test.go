package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	verifyMessage := func(t *testing.T, results string, target string) {
		t.Helper()
		if results != target {
			t.Errorf("Results: '%s' | Target: '%s'", results, target)
		}
	}
	t.Run("Diz olá para as pessoas", func(t *testing.T) {
		results := Hello("Chris", "")
		target := "Hello, Chris"
		verifyMessage(t, results, target)
	})

	t.Run("Diz 'Olá, mundo' quando for uma string vazia", func(t *testing.T) {
		results := Hello("", "")
		target := "Hello, World"
		verifyMessage(t, results, target)
	})

	t.Run("Em espanhol", func(t *testing.T) {
		results := Hello("Elodie", "spanish")
		target := "Hola, Elodie"
		verifyMessage(t, results, target)
	})

	t.Run("Frances", func(t *testing.T) {
		results := Hello("Finn", "frances")
		target := "Bonjour, Finn"
		verifyMessage(t, results, target)
	})
}
