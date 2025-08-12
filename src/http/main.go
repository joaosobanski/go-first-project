package main

import (
	"fmt"
	"net/http"
)

func main() {
	iniciarMonitoramento()
}

// outras funções
// restante do código omitido

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
