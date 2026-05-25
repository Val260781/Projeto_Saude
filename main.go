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

	// Rotas de Consultas
	mux.HandleFunc("GET /consultas", handlers.ListarConsultas)
	mux.HandleFunc("GET /consultas/{id}", handlers.BuscarConsulta)
	mux.HandleFunc("POST /consultas", handlers.CriarConsulta)
	mux.HandleFunc("PUT /consultas/{id}", handlers.AtualizarConsulta)
	mux.HandleFunc("DELETE /consultas/{id}", handlers.DeletarConsulta)

	// Rotas de Médicos
	mux.HandleFunc("GET /medicos", handlers.ListarMedicos)
	mux.HandleFunc("GET /medicos/{id}", handlers.BuscarMedico)
	mux.HandleFunc("POST /medicos", handlers.CriarMedico)
	mux.HandleFunc("PUT /medicos/{id}", handlers.AtualizarMedico)
	mux.HandleFunc("DELETE /medicos/{id}", handlers.DeletarMedico)

	// Rotas de Pacientes
	mux.HandleFunc("GET /pacientes", handlers.ListarPacientes)
	mux.HandleFunc("GET /pacientes/{id}", handlers.BuscarPaciente)
	mux.HandleFunc("POST /pacientes", handlers.CriarPaciente)
	mux.HandleFunc("PUT /pacientes/{id}", handlers.AtualizarPaciente)
	mux.HandleFunc("DELETE /pacientes/{id}", handlers.DeletarPaciente)

	// Arquivos estáticos
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	porta := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", porta)

	//Aplicando corsMiddleware
	log.Fatal(http.ListenAndServe(porta, corsMiddleware(mux)))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
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
