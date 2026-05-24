package storage

import (
	"database/sql"
	"errors"
	"projeto-saude/database"
	"projeto-saude/models"
)

var ErrNotFound = errors.New("Consulta não encontrada")

func GetAll() ([]models.Consulta, error) {
	rows, err := database.DB.Query(
		"SELECT id, paciente_id, medico_id, data, hora, descricao, status FROM consultas",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Consulta
	for rows.Next() {
		var c models.Consulta
		rows.Scan(&c.ID, &c.Paciente, &c.Medico, &c.Data, &c.Hora, &c.Descricao, &c.Status)
		lista = append(lista, c)
	}
	if lista == nil {
		lista = []models.Consulta{}
	}
	return lista, nil
}
func GetByID(id int) (models.Consulta, error) {
	var c models.Consulta
	err := database.DB.QueryRow(
		"SELECT id, paciente, medico, data, hora, descricao, status FROM consultas WHERE id = $1", id,
	).Scan(&c.ID, &c.Paciente, &c.Medico, &c.Data, &c.Hora, &c.Descricao, &c.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Consulta{}, ErrNotFound
	}
	return c, err
}
func Create(c models.Consulta) (models.Consulta, error) {
	err := database.DB.QueryRow(
		"INSERT INTO consultas (paciente, medico, data, hora, descricao, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		c.Paciente, c.Medico, c.Data, c.Hora, c.Descricao, c.Status,
	).Scan(&c.ID)
	return c, err
}
func Update(id int, c models.Consulta) (models.Consulta, error) {
	res, err := database.DB.Exec(
		"UPDATE consultas SET paciente = $1, medico = $2, data = $3, hora = $4, descricao = $5, status = $6 WHERE id = $7",
		c.Paciente, c.Medico, c.Data, c.Hora, c.Descricao, c.Status, id,
	)
	if err != nil {
		return models.Consulta{}, err
	}
	linhas, _ := res.RowsAffected()
	if linhas == 0 {
		return models.Consulta{}, ErrNotFound
	}
	c.ID = id
	return c, nil
}
func Delete(id int) error {
	res, err := database.DB.Exec("DELETE FROM consultas WHERE id = $1", id)
	if err != nil {
		return err
	}
	linhas, _ := res.RowsAffected()
	if linhas == 0 {
		return ErrNotFound
	}
	return nil
}
