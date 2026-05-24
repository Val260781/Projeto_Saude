package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"projeto-saude/database"
	"projeto-saude/handlers"
)

func main() {
	carregarEnv(".env")
	database.Conectar()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /consultas", handlers.ListarConsultas)
	mux.HandleFunc("GET /consultas/{id}", handlers.BuscarConsulta)
	mux.HandleFunc("POST /consultas/", handlers.CriarConsulta)
	mux.HandleFunc("PUT /consultas/{id}", handlers.AtualizarConsulta)
	mux.HandleFunc("DELETE /consultas/{id}", handlers.DeletarConsulta)

	porta := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", porta)
	log.Fatal(http.ListenAndServe(porta, mux))
}

func carregarEnv(arquivo string) {
	dados, err := os.ReadFile(arquivo)
	if err != nil {
		return
	}
	for _, linha := range splitlinhas(string(dados)) {
		if len(linha) == 0 || linha[0] == '#' {
			continue
		}
		for i, c := range linha {
			if c == '=' {
				os.Setenv(linha[:i], linha[i+1:])
				break
			}
		}
	}

}

func splitlinhas(s string) []string {
	var linhas []string
	inicio := 0
	for i, c := range s {
		if c == '\n' {
			linhas = append(linhas, s[inicio:i])
			inicio = i + 1
		}
	}
	if inicio < len(s) {
		linhas = append(linhas, s[inicio:])
	}
	return linhas
}
