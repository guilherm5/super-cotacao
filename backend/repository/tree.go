package repository

import (
	"database/sql"
	"supercotacao/backend/models"
)

func InsertTree(db *sql.DB, tree models.Tree) (models.Tree, error) {
	query := `
		INSERT INTO public.sis_arvore (
			nome_arvore, id_usuario_inclusao
		)
		VALUES ($1, $2)
		RETURNING nome_arvore, id_usuario_inclusao
	`

	err := db.QueryRow(
		query,
		&tree.NameTree,
		&tree.CreatedBy,
	).Scan(
		&tree.NameTree,
		&tree.CreatedBy,
	)

	return tree, err
}

func CheckTreeExists(db *sql.DB, nameTree string) (bool, error) {
	var exists bool
	query := `
		SELECT	EXISTS(
			SELECT	1
			FROM	public.sis_arvore
			WHERE	nome_arvore = $1
		)
	`

	err := db.QueryRow(query, nameTree).Scan(&exists)
	return exists, err
}
