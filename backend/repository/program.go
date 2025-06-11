package repository

import (
	"database/sql"
	"supercotacao/backend/models"
)

func InsertProgram(db *sql.DB, program models.Program) (models.Program, error) {
	query := `
		INSERT INTO public.sis_programa (
			nome_programa, id_usuario_inclusao
		)
		VALUES ($1, $2)
		RETURNING nome_programa, id_usuario_inclusao
	`

	err := db.QueryRow(
		query,
		&program.NameProgram,
		&program.CreatedBy,
	).Scan(
		&program.NameProgram,
		&program.CreatedBy,
	)

	return program, err
}

func CheckProgramExists(db *sql.DB, nameProgram string) (bool, error) {
	var exists bool
	query := `
		SELECT	EXISTS(
			SELECT	1
			FROM	public.sis_programa
			WHERE	nome_programa = $1
		)
	`

	err := db.QueryRow(query, nameProgram).Scan(&exists)
	return exists, err
}
