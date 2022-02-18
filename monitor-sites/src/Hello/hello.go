package main

import (
	"fmt"
)

func main() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão atual: ", versao)
	//fmt.Println("O tipo da variável é: ", reflect.TypeOf(nome))
	fmt.Println("\n1 - Iniciar o monitoranento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")

	var comando int
	fmt.Scan(&comando)

	if comando == 1 {

	} else if comando == 2 {

	} else if comando == 3 {
		fmt.Println("\nSaindo...")
	} else {
		fmt.Println("\nComando inválido!")
	}
}
