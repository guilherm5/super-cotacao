package repository

import (
	"database/sql"

	"supercotacao/backend/models"
)

func SelectUserLogin(input models.User, db *sql.DB) (models.User, error) {
	var user models.User
	query := "SELECT * FROM public.sis_usuario WHERE email_usuario = $1"

	err := db.QueryRow(query, input.UserEmail).Scan(
		&user.UserID,
		&user.UserCode,
		&user.UserName,
		&user.UserEmail,
		&user.UserPassword,
		&user.Cnpj_Cpf,
		&user.UserUUID,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
		&user.DeletedAt,
		&user.DeletedBy,
		&user.IsActive,
	)
	return user, err
}
