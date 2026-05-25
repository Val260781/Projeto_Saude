package models

type Medico struct {
	ID            int    `json:"id"`
	Nome          string `json:"nome"`
	Especialidade string `json:"especialidade"`
	CRM           string `json:"crm"`
}
