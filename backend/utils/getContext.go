package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserFromContextPtr(context *gin.Context) (*int64, error) {
	idUsuario, exists := context.Get("user_id")
	if !exists {
		Logger.Info("ID do Usuário não encontrado no contexto")
		return nil, fmt.Errorf("erro ao recuperar ID do Usuário. Tente novamente mais tarde")
	}

	idUserAction, ok := idUsuario.(int64)
	if !ok {
		Logger.Info("Erro ao converter ID do Usuário para int64")
		return nil, fmt.Errorf("erro ao processar ID do Usuário. Tente novamente mais tarde")
	}

	return &idUserAction, nil
}
