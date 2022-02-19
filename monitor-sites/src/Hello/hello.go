package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoranento = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
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
	fmt.Println(" ")

	return comandoLido
}

func initMonitoramento() {
	fmt.Println("\nMonitorando...")
	sites := []string{
		"https://www.google.com.br",
		"https://www.youtube.com"}

	for i := 0; i < monitoranento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		for i, site := range sites {
			fmt.Println("testando", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
		fmt.Println(" ")
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "site está com Problema", resp.StatusCode)
	}
}
