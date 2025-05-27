package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"supercotacao/backend/models"
	"supercotacao/backend/repository"
	"supercotacao/backend/utils"
)

func CreateUser(input models.User, db *sql.DB) (models.User, error) {
	userExists, err := repository.CheckUserExists(db, input.UserEmail)
	if err != nil {
		utils.Logger.Info("Erro ao verificar existência do usuário ", err)
		return input, errors.New("Erro ao verificar existência do usuário. Tente novamente mais tarde")
	} else if userExists {
		return input, fmt.Errorf("Usuário com e-mail '%s' já existente na base de dados, tente novamente com outro email.", input.UserEmail)
	}

	if !utils.VerifyMailAddress(input.UserEmail) {
		utils.Logger.Info("E-mail é inválido. Tente novamente com um e-mail válido ", err)
		return input, errors.New("E-mail é inválido. Tente novamente com um e-mail válido")
	}

	hashedPassword, err := utils.CreateHash(input.UserPassword)
	if err != nil {
		utils.Logger.Info("Erro ao validar senha para criar novo usuário ", err)
		return models.User{}, errors.New("Erro ao validar senha para criar novo usuário. Tente novamente mais tarde")
	}

	userCreated, err := repository.InsertUser(db, input, hashedPassword, utils.GenerateCode())
	if err != nil {
		utils.Logger.Info("Erro ao criar novo usuário ", err)
		return models.User{}, errors.New("erro ao criar novo usuário. tente novamente mais tarde")
	}

	log.Println("Usuário criado com sucesso")

	return userCreated, nil
}
