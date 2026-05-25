package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"projeto-saude/database"
	"projeto-saude/models"
	"strconv"
)

func ListarPacientes(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, nome, cpf, email, telefone, data_nascimento, endereco, data_cadastro
		FROM pacientes
		ORDER BY nome ASC
	`)
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar pacientes")
		return
	}
	defer rows.Close()

	var pacientes []models.Paciente
	for rows.Next() {
		var p models.Paciente
		err := rows.Scan(&p.ID, &p.Nome, &p.CPF, &p.Email, &p.Telefone, &p.DataNascimento, &p.Endereco, &p.DataCadastro)
		if err != nil {
			responderErro(w, http.StatusInternalServerError, "Erro ao ler paciente")
			return
		}
		pacientes = append(pacientes, p)
	}

	if err := rows.Err(); err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao iterar pacientes")
		return
	}

	responderJSON(w, http.StatusOK, pacientes)
}

func BuscarPaciente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var p models.Paciente
	err = database.DB.QueryRow(`
		SELECT id, nome, cpf, email, telefone, data_nascimento, endereco, data_cadastro
		FROM pacientes WHERE id = $1
	`, id).Scan(&p.ID, &p.Nome, &p.CPF, &p.Email, &p.Telefone, &p.DataNascimento, &p.Endereco, &p.DataCadastro)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responderErro(w, http.StatusNotFound, "Paciente nao encontrado")
		} else {
			responderErro(w, http.StatusInternalServerError, "Erro ao buscar paciente")
		}
		return
	}

	responderJSON(w, http.StatusOK, p)
}

func CriarPaciente(w http.ResponseWriter, r *http.Request) {
	var p models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		responderErro(w, http.StatusBadRequest, "Requisicao invalida")
		return
	}

	if p.Nome == "" || p.CPF == "" {
		responderErro(w, http.StatusBadRequest, "Campos obrigatorios: nome e cpf")
		return
	}

	err := database.DB.QueryRow(`
		INSERT INTO pacientes (nome, cpf, email, telefone, data_nascimento, endereco)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, data_cadastro
	`, p.Nome, p.CPF, p.Email, p.Telefone, p.DataNascimento, p.Endereco,
	).Scan(&p.ID, &p.DataCadastro)

	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar paciente")
		return
	}

	responderJSON(w, http.StatusCreated, p)
}

func AtualizarPaciente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var p models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		responderErro(w, http.StatusBadRequest, "Requisicao invalida")
		return
	}

	_, err = database.DB.Exec(`
		UPDATE pacientes
		SET nome=$1, cpf=$2, email=$3, telefone=$4, data_nascimento=$5, endereco=$6
		WHERE id=$7
	`, p.Nome, p.CPF, p.Email, p.Telefone, p.DataNascimento, p.Endereco, id)

	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao atualizar paciente")
		return
	}

	p.ID = id
	responderJSON(w, http.StatusOK, p)
}

func DeletarPaciente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}

	result, err := database.DB.Exec("DELETE FROM pacientes WHERE id = $1", id)
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao deletar paciente")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		responderErro(w, http.StatusNotFound, "Paciente nao encontrado")
		return
	}

	responderJSON(w, http.StatusOK, map[string]string{"mensagem": "Paciente removido com sucesso"})
}
