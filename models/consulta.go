package models

type Consulta struct {
	ID        int    `json:"id"`
	Paciente  string `json:"paciente"`
	Medico    string `json:"medico"`
	Data      string `json:"data"`
	Hora      string `json:"hora"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
}
