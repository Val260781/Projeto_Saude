package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"projeto-saude/models"
	"projeto-saude/storage"
	"strconv"
)

func responderjson(w http.ResponseWriter, status int, dado any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dado)
}

func responderErro(w http.ResponseWriter, status int, mensagem string) {
	responderjson(w, status, map[string]string{"erro": mensagem})
}

func ListarConsultas(w http.ResponseWriter, r *http.Request) {
	consultas, err := storage.GetAll()
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar consultas")
		return
	}
	responderjson(w, http.StatusOK, consultas)
}

func BuscarConsulta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}
	consulta, err := storage.GetByID(id)
	if errors.Is(err, storage.ErrNotFound) {
		responderErro(w, http.StatusNotFound, "Consulta nao encontrada")
		return
	}
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar consulta")
		return
	}
	responderjson(w, http.StatusOK, consulta)
}

func CriarConsulta(w http.ResponseWriter, r *http.Request) {
	var nova models.Consulta
	if err := json.NewDecoder(r.Body).Decode(&nova); err != nil {
		responderErro(w, http.StatusBadRequest, "Requisicao invalido")
		return
	}
	if nova.Paciente == "" || nova.Medico == "" || nova.Data == "" {
		responderErro(w, http.StatusBadRequest, "Campos obrigatorios: paciente, medico e data")
		return
	}
	if nova.Status == "" {
		nova.Status = "agendada"
	}
	criada, err := storage.Create(nova)
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar consulta")
		return
	}
	responderjson(w, http.StatusCreated, criada)
}

func AtualizarConsulta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}
	var atualizada models.Consulta
	if err := json.NewDecoder(r.Body).Decode(&atualizada); err != nil {
		responderErro(w, http.StatusBadRequest, "Requisicao invalido")
		return
	}
	consulta, err := storage.Update(id, atualizada)
	if errors.Is(err, storage.ErrNotFound) {
		responderErro(w, http.StatusNotFound, "Consulta nao encontrada")
		return
	}
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao atualizar consulta")
		return
	}
	responderjson(w, http.StatusOK, consulta)
}

func DeletarConsulta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}
	err = storage.Delete(id)
	if errors.Is(err, storage.ErrNotFound) {
		responderErro(w, http.StatusNotFound, "Consulta nao encontrada")
		return
	}
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao deletar consulta")
		return
	}
	responderjson(w, http.StatusOK, map[string]string{"mensagem": "Consulta removida com sucesso!"})
}
