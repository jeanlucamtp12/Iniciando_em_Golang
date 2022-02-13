package main

import "testing"

func TestOla(t *testing.T) {

	verifcaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}

	}
	t.Run("diga olá para as pessoas", func(t *testing.T) {

		resultado := Ola("Chris")
		esperado := "Olá, Chris"

		verifcaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diga 'Olá, mundo' caso a string passada seja vazia", func(t *testing.T) {

		resultado := Ola("")
		esperado := "Olá, mundo"

		verifcaMensagemCorreta(t, resultado, esperado)

	})

}
