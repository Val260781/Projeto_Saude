package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"projeto-saude/database"
	"strconv"
)

type Consulta struct {
	ID            int    `json:"id"`
	Paciente      string `json:"paciente"`
	Medico        string `json:"medico"`
	Especialidade string `json:"especialidade"`
	Data          string `json:"data"`
	Hora          string `json:"hora"`
	Descricao     string `json:"descricao"`
	Status        string `json:"status"`
}

func ListarConsultas(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, paciente, medico, data, hora, descricao, status, especialidade FROM consultas ORDER BY data ASC, hora ASC")
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar consultas")
		return
	}
	defer rows.Close()

	var consultas []Consulta
	for rows.Next() {
		var c Consulta
		err := rows.Scan(&c.ID, &c.Paciente, &c.Medico, &c.Data, &c.Hora, &c.Descricao, &c.Status, &c.Especialidade)
		if err != nil {
			responderErro(w, http.StatusInternalServerError, "Erro ao ler consulta")
			return
		}
		consultas = append(consultas, c)
	}

	if err := rows.Err(); err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao iterar consultas")
		return
	}

	responderJSON(w, http.StatusOK, consultas)
}

func BuscarConsulta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var c Consulta
	err = database.DB.QueryRow("SELECT id, paciente, medico, data, hora, descricao, status, especialidade FROM consultas WHERE id = $1", id).
		Scan(&c.ID, &c.Paciente, &c.Medico, &c.Data, &c.Hora, &c.Descricao, &c.Status, &c.Especialidade)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responderErro(w, http.StatusNotFound, "Consulta nao encontrada")
		} else {
			responderErro(w, http.StatusInternalServerError, "Erro ao buscar consulta")
		}
		return
	}

	responderJSON(w, http.StatusOK, c)
}

func CriarConsulta(w http.ResponseWriter, r *http.Request) {
	var c Consulta
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		responderErro(w, http.StatusBadRequest, "Requisicao invalida")
		return
	}

	if c.Paciente == "" || c.Medico == "" || c.Data == "" {
		responderErro(w, http.StatusBadRequest, "Campos obrigatorios: paciente, medico e data")
		return
	}

	if c.Status == "" {
		c.Status = "agendada"
	}

	err := database.DB.QueryRow(
		"INSERT INTO consultas (paciente, medico, especialidade, data, hora, descricao, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		c.Paciente, c.Medico, c.Especialidade, c.Data, c.Hora, c.Descricao, c.Status,
	).Scan(&c.ID)

	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar consulta")
		return
	}

	responderJSON(w, http.StatusCreated, c)
}

func AtualizarConsulta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}

	var c Consulta
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		responderErro(w, http.StatusBadRequest, "Requisicao invalida")
		return
	}

	_, err = database.DB.Exec(
		"UPDATE consultas SET paciente=$1, medico=$2, especialidade=$3, data=$4, hora=$5, descricao=$6, status=$7 WHERE id=$8",
		c.Paciente, c.Medico, c.Especialidade, c.Data, c.Hora, c.Descricao, c.Status, id,
	)
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao atualizar consulta")
		return
	}

	c.ID = id
	responderJSON(w, http.StatusOK, c)
}

func DeletarConsulta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}

	result, err := database.DB.Exec("DELETE FROM consultas WHERE id = $1", id)
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao deletar consulta")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		responderErro(w, http.StatusNotFound, "Consulta nao encontrada")
		return
	}

	responderJSON(w, http.StatusOK, map[string]string{"mensagem": "Consulta removida com sucesso"})
}
