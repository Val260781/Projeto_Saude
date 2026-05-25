package storage

import (
	"projeto-saude/database"
	"projeto-saude/models"
)

func GetAllMedicos() ([]models.Medico, error) {
	rows, err := database.DB.Query(
		"SELECT id, nome, especialidade, crm FROM medicos ORDER BY nome",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Medico
	for rows.Next() {
		var m models.Medico
		rows.Scan(&m.ID, &m.Nome, &m.Especialidade, &m.CRM)
		lista = append(lista, m)
	}

	if lista == nil {
		lista = []models.Medico{}
	}
	return lista, nil
}

func GetMedicoByID(id int) (models.Medico, error) {
	var m models.Medico
	err := database.DB.QueryRow(
		"SELECT id, nome, especialidade, crm FROM medicos WHERE id = $1", id,
	).Scan(&m.ID, &m.Nome, &m.Especialidade, &m.CRM)
	return m, err
}

func CreateMedico(m models.Medico) (models.Medico, error) {
	err := database.DB.QueryRow(
		"INSERT INTO medicos (nome, especialidade, crm) VALUES ($1, $2, $3) RETURNING id",
		m.Nome, m.Especialidade, m.CRM,
	).Scan(&m.ID)
	return m, err
}

func UpdateMedico(id int, m models.Medico) (models.Medico, error) {
	res, err := database.DB.Exec(
		"UPDATE medicos SET nome=$1, especialidade=$2, crm=$3 WHERE id=$4",
		m.Nome, m.Especialidade, m.CRM, id,
	)
	if err != nil {
		return models.Medico{}, err
	}
	linhas, _ := res.RowsAffected()
	if linhas == 0 {
		return models.Medico{}, ErrNotFound
	}
	m.ID = id
	return m, nil
}

func DeleteMedico(id int) error {
	res, err := database.DB.Exec("DELETE FROM medicos WHERE id = $1", id)
	if err != nil {
		return err
	}
	linhas, _ := res.RowsAffected()
	if linhas == 0 {
		return ErrNotFound
	}
	return nil
}
