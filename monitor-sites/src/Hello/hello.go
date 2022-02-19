package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoranento = 1
const delay = 5

func main() {
	exibeIntroducao()
	//registraLog("site-falso", false)

	for {
		exibeMenu()
		//var comando int
		comando := lerComando()

		switch comando {
		case 1:
			initMonitoramento()
		case 2:
			fmt.Println("\nExibindo Logs...")
			imprimeLogs()
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

	lerSiteArquivo()
	lerSiteArquivo()

	sites := lerSiteArquivo()

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
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ops!!!Ocorreu um erro: ", err)
	}
	fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "site está com Problema", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSiteArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ops!!!Ocorreu um erro: ", err)
	}

	read := bufio.NewReader(arquivo)
	for {
		linha, err := read.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile(
		"log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ops!!!Ocorreu um erro: ", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	//fmt.Println(arquivo)
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFIle("log.txt")
	if err != nil {
		fmt.Println("Ops!!!Ocorreu um erro: ", err)
	}
	fmt.Println(string(arquivo))
}
