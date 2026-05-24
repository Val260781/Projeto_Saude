package models

type Consulta struct {
	ID        int    `json:"id"`
	Paciente  string `json:"paciente_id"`
	Medico    string `json:"medico_id"`
	Data      string `json:"data"`
	Hora      string `json:"hora"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
}
