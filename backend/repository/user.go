package repository

import (
	"database/sql"
	"supercotacao/backend/models"
)

/*Tem um bug, quando scanear o retorno, ele ainda insere no banco, bug ou feature? não sei kk*/
func InsertUser(db *sql.DB, user models.User, hashedPassword string, codigoUsuario string) (models.User, error) {
	query := `
		INSERT INTO public.sis_usuario (
			nome_usuario, email_usuario, senha_usuario, codigo_usuario, cnpj_cpf, registro_ativo
		) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_usuario, email_usuario, codigo_usuario, registro_ativo, uuid_usuario
	`

	err := db.QueryRow(
		query,
		user.UserName,
		user.UserEmail,
		hashedPassword,
		codigoUsuario,
		user.Cnpj_Cpf,
		user.IsActive,
	).Scan(
		&user.UserID,
		&user.UserEmail,
		&user.UserCode,
		&user.IsActive,
		&user.UserUUID,
	)

	return user, err
}

func CheckUserExists(db *sql.DB, email string) (bool, error) {
	var exists bool
	query := `
		SELECT	EXISTS(
			SELECT	1
			FROM	public.sis_usuario
			WHERE	email_usuario = $1
		)
	`

	err := db.QueryRow(query, email).Scan(&exists)
	return exists, err
}
