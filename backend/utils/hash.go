package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Erro ao gerar hash da senha:", err)
		return "", err
	}

	return string(hash), nil
}
