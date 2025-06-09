package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	UserID       int64      `json:"id_usuario"`
	UserCode     string     `json:"codigo_usuario"`
	UserName     string     `json:"nome_usuario"`
	UserEmail    string     `json:"email_usuario"`
	UserPassword string     `json:"senha_usuario"`
	Cnpj_Cpf     string     `json:"cnpj_cpf"`
	UserUUID     uuid.UUID  `json:"uuid_usuario"`
	CreatedAt    time.Time  `json:"data_inclusao"`
	CreatedBy    *int       `json:"id_usuario_inclusao"`
	UpdatedAt    *time.Time `json:"data_alteracao"`
	UpdatedBy    *int       `json:"id_usuario_alteracao"`
	DeletedAt    *time.Time `json:"data_exclusao"`
	DeletedBy    *int       `json:"id_usuario_exclusao"`
	IsActive     bool       `json:"registro_ativo"`
}
