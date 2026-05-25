package models

type Paciente struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	CPF            string `json:"cpf"`
	Email          string `json:"email"`
	Telefone       string `json:"telefone"`
	DataNascimento string `json:"data_nascimento"`
	Endereco       string `json:"endereco"`
	DataCadastro   string `json:"data_cadastro"`
}
