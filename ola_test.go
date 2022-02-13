package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}

	}

	t.Run("em ingles", func(t *testing.T) {
		resultado := Ola("Matt", "ingles")
		esperado := "Hello, Matt"

		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("Jeanluca", "frances")
		esperado := "Bonjour, Jeanluca"

		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("em espanhol", func(t *testing.T) {

		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("diga olá para as pessoas", func(t *testing.T) {

		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diga 'Olá, mundo' caso a string passada seja vazia", func(t *testing.T) {

		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)

	})

}
