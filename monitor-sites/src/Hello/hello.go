package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	//var comando int
	comando := lerComando()

	switch comando {
	case 1:
		initMonitoramento()
	case 2:
		fmt.Println("\nExibindo Logs...")
	case 3:
		fmt.Println("\nSaindo...")
		os.Exit(0)
	default:
		fmt.Println("\nNão conheço esse programa...")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão atual: ", versao)
}

func exibeMenu() {
	fmt.Println("\n1 - Iniciar o monitoranento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func initMonitoramento() {
	fmt.Println("\nMonitorando...")
	site := "https://www.mrs.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
