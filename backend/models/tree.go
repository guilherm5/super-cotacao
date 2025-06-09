package models

import (
	"time"
)

type Tree struct {
	TreeID    int        `json:"id_arvore"`
	NameTree  string     `json:"nome_arvore"`
	CreatedAt time.Time  `json:"data_inclusao"`
	CreatedBy *int64     `json:"id_usuario_inclusao"`
	UpdatedAt *time.Time `json:"data_alteracao"`
	UpdatedBy *int       `json:"id_usuario_alteracao"`
	DeletedAt *time.Time `json:"data_exclusao"`
	DeletedBy *int       `json:"id_usuario_exclusao"`
}
