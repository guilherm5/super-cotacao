package utils

import (
	"fmt"
	"net/mail"
)

func VerifyMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		fmt.Printf("Erro ao verificar email: %s. Erro: %s", address, err)
		return false
	}
	return true
}
