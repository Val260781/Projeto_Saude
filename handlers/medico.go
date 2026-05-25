package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"projeto-saude/models"
	"projeto-saude/storage"
	"strconv"
)

func ListarMedicos(w http.ResponseWriter, r *http.Request) {
	medicos, err := storage.GetAllMedicos()
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar medicos")
		return
	}
	responderJSON(w, http.StatusOK, medicos)
}

func BuscarMedico(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}
	medico, err := storage.GetMedicoByID(id)
	if errors.Is(err, storage.ErrNotFound) {
		responderErro(w, http.StatusNotFound, "Medico nao encontrado")
		return
	}
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar medico")
		return
	}
	responderJSON(w, http.StatusOK, medico)
}

func CriarMedico(w http.ResponseWriter, r *http.Request) {
	var novo models.Medico
	if err := json.NewDecoder(r.Body).Decode(&novo); err != nil {
		responderErro(w, http.StatusBadRequest, "Corpo da requisicao invalido")
		return
	}
	if novo.Nome == "" || novo.Especialidade == "" || novo.CRM == "" {
		responderErro(w, http.StatusBadRequest, "Campos obrigatorios: nome, especialidade, crm")
		return
	}
	criado, err := storage.CreateMedico(novo)
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar medico")
		return
	}
	responderJSON(w, http.StatusCreated, criado)
}

func AtualizarMedico(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}
	var atualizado models.Medico
	if err := json.NewDecoder(r.Body).Decode(&atualizado); err != nil {
		responderErro(w, http.StatusBadRequest, "Corpo da requisicao invalido")
		return
	}
	medico, err := storage.UpdateMedico(id, atualizado)
	if errors.Is(err, storage.ErrNotFound) {
		responderErro(w, http.StatusNotFound, "Medico nao encontrado")
		return
	}
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao atualizar medico")
		return
	}
	responderJSON(w, http.StatusOK, medico)
}

func DeletarMedico(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responderErro(w, http.StatusBadRequest, "ID invalido")
		return
	}
	err = storage.DeleteMedico(id)
	if errors.Is(err, storage.ErrNotFound) {
		responderErro(w, http.StatusNotFound, "Medico nao encontrado")
		return
	}
	if err != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao deletar medico")
		return
	}
	responderJSON(w, http.StatusOK, map[string]string{"mensagem": "Medico removido com sucesso"})
}
