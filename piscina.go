package main

import (
	"fmt"
	"math"
)

func CalcularVolumeRetangular(comprimento, largura, profundidade float64) float64 {
	if comprimento < 0 || largura < 0 || profundidade < 0 {
		return -1
	}
	return comprimento * largura * profundidade
}

func CalcularVolumeRetangularInterativo() float64 {
	var comprimento, largura, profundidade float64

	fmt.Println("\nFormato selecionado: Piscina Retangular")

	fmt.Print("\nDigite o comprimento: ")
	fmt.Scanf("%f", &comprimento)

	fmt.Print("Digite o largura: ")
	fmt.Scanf("%f", &largura)

	fmt.Print("Digite o profundidade: ")
	fmt.Scanf("%f", &profundidade)

	volume := CalcularVolumeRetangular(comprimento, largura, profundidade)

	if volume < 0 {
		fmt.Println("Erro: Valores não podem ser negativos.")
	} else {
		fmt.Printf("\nVolume = %.2f m³\n", volume)
	}

	return volume
}

func CalcularVolumeOval(diametroMaior, diametroMenor, profundidade float64) float64 {
	if diametroMaior < 0 || diametroMenor < 0 || profundidade < 0 {
		return -1
	}
	return diametroMaior * diametroMenor * profundidade * 0.785
}

func CalcularVolumeOvalInterativo() float64 {
	var diametroMaior, diametroMenor, profundidade float64

	fmt.Println("\nFormato selecionado: Piscina Oval:")

	fmt.Print("\nDigite o diâmetro maior: ")
	fmt.Scanf("%f", &diametroMaior)

	fmt.Print("Digite o diâmetro menor: ")
	fmt.Scanf("%f", &diametroMenor)

	fmt.Print("Digite a profundidade: ")
	fmt.Scanf("%f", &profundidade)

	volume := CalcularVolumeOval(diametroMaior, diametroMenor, profundidade)

	if volume < 0 {
		fmt.Println("Erro: Valores não podem ser negativos.")
	} else {
		fmt.Printf("\nVolume = %.2f m³\n", volume)
	}

	return volume
}

func CalcularVolumeRedondo(diametro, profundidade float64) float64 {
	if diametro < 0 || profundidade < 0 {
		return -1
	}
	raio := diametro / 2
	return math.Pi * raio * raio * profundidade
}

func CalcularVolumeRedondoInterativo() float64 {
	var diametro, profundidade float64

	fmt.Println("\nFormato selecionado: Piscina Redonda")

	fmt.Print("\nDigite o diâmetro: ")
	fmt.Scanf("%f", &diametro)

	fmt.Print("\nDigite a profundidade: ")
	fmt.Scanf("%f", &profundidade)

	volume := CalcularVolumeRedondo(diametro, profundidade)

	if volume < 0 {
		fmt.Println("\nErro: Valores não podem ser negativos.")
	} else {
		fmt.Printf("\nVolume = %.2f m³\n", volume)
	}

	return volume
}

func CalcularDosagens(M float64) (dosagemInicialAlgistatico, dosagemSemanalAlgistatico, dosagemCloroLiquido float64) {
	dosagemInicialAlgistatico = M * 14
	dosagemSemanalAlgistatico = M * 6
	dosagemCloroLiquido = M * 25
	return
}

func CalcularDecantacao(M float64) (sulfatoAluminio, barrilhaLeve float64) {
	sulfatoAluminio = M * 60
	barrilhaLeve = M * 30
	return
}

// Retorna a dosagem recomendada e um booleano indicando se é necessária correção
func CalcularCorrecaoPH(M, phAtual float64) (dosagem float64, precisaCorrigir bool) {
	var phRecomendado float64

	switch {
	case phAtual < 6.5:
		phRecomendado = 15.8
	case phAtual < 6.8:
		phRecomendado = 12.7
	case phAtual < 7.2:
		phRecomendado = 9.5
	case phAtual > 8.4:
		phRecomendado = 19
	case phAtual >= 8 && phAtual <= 8.4:
		phRecomendado = 12.7
	case phAtual > 7.6 && phAtual < 8:
		phRecomendado = 6.5
	default:
		return 0, false // não precisa corrigir
	}

	dosagem = M * phRecomendado
	return dosagem, true
}

func CalcularSuperdosagens(M float64) (algicida, cloro float64) {
	algicida = M * 16
	cloro = M * 50
	return
}

func ExecutarTratamento(opcao int, M float64) string {
	switch opcao {
	case 1:
		algInicial, algSemanal, cloro := CalcularDosagens(M)
		print("\n")
		return fmt.Sprintf(`--- Dosagens Diárias e Semanais ---
Dosagem Inicial de Algistático: %.2f ml
Dosagem Semanal de Algistático: %.2f ml
Dosagem Diária de Cloro Líquido: %.2f ml`, algInicial, algSemanal, cloro)

	case 2:
		sulfato, barrilha := CalcularDecantacao(M)
		return fmt.Sprintf(`--- Decantação para Aspirar ---
Sulfato de Alumínio: %.2f gramas
Barrilha Leve: %.2f gramas

Instruções:
- Dissolva o sulfato de alumínio em um balde com água e espalhe sobre a piscina.
- Dissolva a barrilha leve separadamente e aplique igualmente.
- Mantenha a piscina sem filtrar por 12 a 24 horas.
- Aspire o fundo da piscina após a decantação.`, sulfato, barrilha)

	case 3:
		return "\n--- Correção de pH --- (interação com usuário necessária)"

	case 4:
		algicida, cloro := CalcularSuperdosagens(M)
		return fmt.Sprintf(`--- Superdosagens ---
Superdosagem de Algicida: %.2f ml
Superdosagem de Cloro Líquido: %.2f ml

Instruções:
- Aplique o algicida diretamente na piscina, espalhando por toda a superfície.
- Em seguida, aplique o cloro líquido de forma uniforme.
- Mantenha a piscina sem uso por 24 a 48 horas.
- Após esse período, filtre e aspire o fundo se necessário.`, algicida, cloro)

	case 5:
		return "--- Opção ainda será implementada ---"

	case 6:
		return "\nEncerrando o programa. Até logo!"

	default:
		return "\nOpção inválida - escolha de 1 a 6"
	}
}

func main() {

	var volume float64
	var formatoPiscina int

	fmt.Println("\n===== Cálculo de Volume da Piscina =====")
	fmt.Println("\nEscolha o Formato da Piscina:")
	fmt.Println("\n1 - Retangular")
	fmt.Println("2 - Oval")
	fmt.Println("3 - Redonda")
	fmt.Print("\nDigite a opção desejada: ")
	fmt.Scanf("%d", &formatoPiscina)

	if formatoPiscina < 1 || formatoPiscina > 3 {
		fmt.Println("Opção inválida. Por favor, escolha 1, 2 ou 3.")
		return
	}

	fmt.Println("\nInforme as Dimensões da Piscina:")

	switch formatoPiscina {
	case 1:
		volume = CalcularVolumeRetangularInterativo()
	case 2:
		volume = CalcularVolumeOvalInterativo()
	case 3:
		volume = CalcularVolumeRedondoInterativo()
	}

	if volume < 0 {
		fmt.Println("Erro: volume inválido.")
		return
	}

	/*
		Parte 2: Escolha de Opções para Tratamento da Piscina
		Nesta parte, o código permite ao usuário escolher entre várias opções de
		tratamento de piscina com base no volume calculado anteriormente.
	*/

	M := volume

	var opcao int

	for {
		fmt.Println("\nEscolha uma das opções abaixo:")
		fmt.Println("\nOpção 1: Dosagens Diárias e Semanais")
		fmt.Println("Opção 2: Decantação para Aspirar")
		fmt.Println("Opção 3: Correção de pH")
		fmt.Println("Opção 4: Superdosagens")
		fmt.Println("Opção 5: Precisa ser implementada")
		fmt.Println("Opção 6: FIM")

		fmt.Print("\nDigite o número da opção: ")
		fmt.Scanf("%d", &opcao)

		resultado := ExecutarTratamento(opcao, M)
		fmt.Println(resultado)

		if opcao == 6 {
			break
		}
	}

}
