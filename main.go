package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Defina o valor unitario da maçã: ")

	//var precoUnitarioMaca string - assim tu inicializa a variavel vazia

	//assim tu ja inicializa e ja usa a variavel
	precoUnitarioMaca := ""
	fmt.Scan(&precoUnitarioMaca)

	fmt.Println(precoUnitarioMaca)

	precoUnitarioMacaFloat, err := strconv.ParseFloat(precoUnitarioMaca, 64)
	if err != nil {
		fmt.Println("Erro ao converter preço da maça de string para decimal", err.Error())
		return
	}

	fmt.Println("Quantos reais de maça a cliente deseja??")

	var pedidoCliente string
	fmt.Scan(&pedidoCliente) //R$10.50

	valorPagamento := strings.TrimPrefix(pedidoCliente, "R$")

	valorPgtoFloat, err := strconv.ParseFloat(valorPagamento, 64)
	if err != nil {
		fmt.Println("Erro ao converter o valor do pagamento do cliente de string para decimal", err.Error())
		return
	}

	qtdMacaAComprar := valorPgtoFloat / precoUnitarioMacaFloat

	valorPgtoFormatado := strconv.FormatFloat(valorPgtoFloat, 'f', 2, 64)

	qtdMacaInt := int(qtdMacaAComprar)
	fmt.Println("Quantas maças é possível comprar?", qtdMacaInt)

	fmt.Printf("A cliente pode comprar %d maçãs com o valor de %s", qtdMacaInt, valorPgtoFormatado)
}
