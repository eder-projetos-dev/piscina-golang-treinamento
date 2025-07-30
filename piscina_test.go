package main

import "testing"

/*

CENÁRIOS DE TESTE

	Funcionalidade: Cálculo do volume da piscina e seleção de tratamentos

	Cenário: Cálculo de volume para piscina retangular com dimensões válidas
	Dado que o usuário selecionou a opção 1 (Retangular)

	Quando informa comprimento 5m, largura 3m e profundidade 1.5m
	Então o sistema deve calcular o volume como 22.5 m3

*/

func TestVolumeRetangular(t *testing.T) {
	resultado := CalcularVolumeRetangular(5, 3, 1.5)
	esperado := 22.5

	if resultado != esperado {
		t.Errorf("esperado %.2f, mas obteve %.2f", esperado, resultado)
	}
}

/*

Cenário: Tentativa de entrada com valor negativo para diâmetro redondo
Dado que o usuário selecionou a opção 3 (Redonda)

Quando informa diâmetro -4m e profundidade 2m
Então o sistema deve exibir "Erro: Valores não podem ser negativos"

*/

func TestVolumeRedondoValoresNegativos(t *testing.T) {
	diametro := -4.0
	profundidade := 2.0

	volume := CalcularVolumeRedondo(diametro, profundidade)
	esperado := -1.0

	if volume != esperado {
		t.Errorf("Esperado %.2f para entrada inválida, mas obteve %.2f", esperado, volume)
	}
}

/*

Cenário: Seleção de opção de tratamento inválida
Dado que o volume M foi calculado como 30m3

Quando o usuário seleciona a opção 7
Então o sistema deve mostrar "Opção inválida - escolha de 1 a 6"

*/

func TestOpcaoInvalida(t *testing.T) {
	M := 30.0
	opcao := 7

	resultado := ExecutarTratamento(opcao, M)
	esperado := "\nOpção inválida - escolha de 1 a 6"

	if resultado != esperado {
		t.Errorf("Esperado: %q, mas obteve: %q", esperado, resultado)
	}
}

/*

Funcionalidade: Tratamentos químicos da piscina

Cenário: Cálculo de dosagem semanal para piscina média
Dado M = 50m3 e opção 1 selecionada

Quando processa dosagens
Então deve retornar 700g algistático inicial
E 300g semanal de algistático

*/

func TestCalcularDosagensParaPiscinaMedia(t *testing.T) {
	M := 50.0

	dosagemInicial, dosagemSemanal, _ := CalcularDosagens(M)

	if dosagemInicial != 700 {
		t.Errorf("Esperado 700g de algistático inicial, mas foi %.2f", dosagemInicial)
	}

	if dosagemSemanal != 300 {
		t.Errorf("Esperado 300g de algistático semanal, mas foi %.2f", dosagemSemanal)
	}
}
