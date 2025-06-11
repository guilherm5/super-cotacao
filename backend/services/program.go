package services

import (
	"database/sql"
	"errors"
	"fmt"
	"supercotacao/backend/models"
	"supercotacao/backend/repository"
	"supercotacao/backend/utils"
)

func CreateProgram(input models.Program, db *sql.DB) (models.Program, error) {
	programExists, err := repository.CheckProgramExists(db, input.NameProgram)
	if err != nil {
		utils.Logger.Info("Erro ao verificar existência do programa ", err)
		return input, errors.New("Erro ao verificar existência do programa. Tente novamente mais tarde")
	} else if programExists {
		return input, fmt.Errorf("Programa com nome '%s' já existente na base de dados, tente novamente com outro nome ou use a arvore já existente.", input.NameProgram)
	}

	if input.NameProgram == "" {
		return input, fmt.Errorf("Nome do programa está inválido ou vazio, adicione um nome válido.")
	}

	programCreated, err := repository.InsertProgram(db, input)
	if err != nil {
		utils.Logger.Info("Erro ao criar nova arvore ", err)
		return models.Program{}, errors.New("erro ao criar novo programa. tente novamente mais tarde")
	}

	return programCreated, nil
}
