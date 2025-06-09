package services

import (
	"database/sql"
	"errors"
	"fmt"
	"supercotacao/backend/models"
	"supercotacao/backend/repository"
	"supercotacao/backend/utils"
)

func CreateTree(input models.Tree, db *sql.DB) (models.Tree, error) {
	treeExists, err := repository.CheckTreeExists(db, input.NameTree)
	if err != nil {
		utils.Logger.Info("Erro ao verificar existência da arvore ", err)
		return input, errors.New("Erro ao verificar existência da arvore. Tente novamente mais tarde")
	} else if treeExists {
		return input, fmt.Errorf("Arvore com nome '%s' já existente na base de dados, tente novamente com outro nome ou use a arvore já existente.", input.NameTree)
	}

	if input.NameTree == "" {
		return input, fmt.Errorf("Nome da arvore está inválido ou vazio, adicione um nome válido.")
	}

	treeCreated, err := repository.InsertTree(db, input)
	if err != nil {
		utils.Logger.Info("Erro ao criar nova arvore ", err)
		return models.Tree{}, errors.New("erro ao criar nova arvore. tente novamente mais tarde")
	}

	return treeCreated, nil
}
