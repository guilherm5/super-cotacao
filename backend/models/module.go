package models

import (
	"time"
)

type Module struct {
	ModuleID          int        `json:"id_modulo"`
	ModuleName        string     `json:"nome_modulo"`
	ModuleUrlIcon     string     `json:"url_icone_modulo"`
	ModuleDescription string     `json:"descricao_modulo"`
	CreatedAt         time.Time  `json:"data_inclusao"`
	CreatedBy         *int64     `json:"id_usuario_inclusao"`
	UpdatedAt         *time.Time `json:"data_alteracao"`
	UpdatedBy         *int       `json:"id_usuario_alteracao"`
	DeletedAt         *time.Time `json:"data_exclusao"`
	DeletedBy         *int       `json:"id_usuario_exclusao"`
}
