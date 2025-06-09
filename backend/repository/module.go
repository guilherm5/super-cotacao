package repository

import (
	"database/sql"
	"supercotacao/backend/models"
)

func InsertModule(db *sql.DB, module models.Module) (models.Module, error) {
	query := `
		INSERT INTO public.sis_modulo (
			nome_modulo, url_icone_modulo, descricao_modulo, id_usuario_inclusao
		)
		VALUES ($1, $2, $3, $4)
		RETURNING nome_modulo, url_icone_modulo, descricao_modulo, id_usuario_inclusao
	`

	err := db.QueryRow(
		query,
		&module.ModuleName,
		&module.ModuleUrlIcon,
		&module.ModuleDescription,
		&module.CreatedBy,
	).Scan(
		&module.ModuleName,
		&module.ModuleUrlIcon,
		&module.ModuleDescription,
		&module.CreatedBy,
	)

	return module, err
}

func CheckModuleExists(db *sql.DB, nomeModulo string) (bool, error) {
	var exists bool
	query := `
		SELECT	EXISTS(
			SELECT	1
			FROM	public.sis_modulo
			WHERE	nome_modulo = $1
		)
	`

	err := db.QueryRow(query, nomeModulo).Scan(&exists)
	return exists, err
}
