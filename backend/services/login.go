package services

import (
	"database/sql"
	"errors"
	"os"
	"strconv"
	"supercotacao/backend/repository"
	"supercotacao/backend/utils"
	"time"

	"supercotacao/backend/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CompareAndGenerateToken(input models.User, db *sql.DB) (string, error) {
	secret := os.Getenv("SECRET")
	user, err := repository.SelectUserLogin(input, db)
	if err != nil {
		utils.Logger.Info("erro ao buscar usuário para iniciar fluxo de login ", err)
		return "", errors.New("usuário não encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(input.UserPassword)); err != nil {
		utils.Logger.Info("erro ao comparar senha ", err)
		return "", errors.New("senha inválida")
	}

	claims := jwt.MapClaims{
		"user_id":   strconv.FormatInt(user.UserID, 10),
		"user_name": user.UserName,
		"exp":       time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		utils.Logger.Info("erro ao gerar token ", err)
		return "", errors.New("erro ao gerar token")
	}

	return tokenString, nil
}
