package services

import (
	"database/sql"
	"errors"
	"fmt"
	"supercotacao/backend/models"
	"supercotacao/backend/repository"
	"supercotacao/backend/utils"
)

func CreateModule(input models.Module, db *sql.DB) (models.Module, error) {
	moduleExists, err := repository.CheckModuleExists(db, input.ModuleName)
	if err != nil {
		utils.Logger.Info("Erro ao verificar existência do módulo ", err)
		return input, errors.New("Erro ao verificar existência do módulo. Tente novamente mais tarde")
	} else if moduleExists {
		return input, fmt.Errorf("Módulo com nome '%s' já existente na base de dados, tente novamente com outro nome ou use o módulo já existente.", input.ModuleName)
	}

	moduleCreated, err := repository.InsertModule(db, input)
	if err != nil {
		utils.Logger.Info("Erro ao criar novo módulo ", err)
		return models.Module{}, errors.New("erro ao criar novo módulo. tente novamente mais tarde")
	}

	return moduleCreated, nil
}
