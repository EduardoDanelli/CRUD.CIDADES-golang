package models

// Cidades seta os valores do id, nome, siglaEstado
type Cidades struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	SiglaEstado string `json:"siglaEstado" gorm:"column:siglaestado"`
}
