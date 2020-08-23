package main

import "testing"

func TestSoma(t *testing.T) {
	tabelas := []struct {
		x int
		y int
		n int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{0, 1, 1},
	}

	for _, tabela := range tabelas {
		total := Soma(tabela.x, tabela.y)
		if total != tabela.n {
			t.Errorf("Soma de (%d+%d) incorreta, obtido: %d, esperado: %d.", tabela.x, tabela.y, total, tabela.n)
		}
	}
}
