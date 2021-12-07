package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	// Executa a aplicação
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
