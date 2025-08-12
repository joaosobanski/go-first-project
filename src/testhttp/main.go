package main

import (
	"fmt"
	"net/http"
)

func main() {
	iniciarMonitoramento()
}

//hello.go
//restante do arquivo

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

// restante do código omitido

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://httpbin.org/status/200",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}

	fmt.Println("")
}
