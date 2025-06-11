package models

import "time"

type Program struct {
	ProgramID   int        `json:"id_programa"`
	NameProgram string     `json:"nome_programa"`
	CreatedAt   time.Time  `json:"data_inclusao"`
	CreatedBy   *int64     `json:"id_usuario_inclusao"`
	UpdatedAt   *time.Time `json:"data_alteracao"`
	UpdatedBy   *int       `json:"id_usuario_alteracao"`
	DeletedAt   *time.Time `json:"data_exclusao"`
	DeletedBy   *int       `json:"id_usuario_exclusao"`
}
