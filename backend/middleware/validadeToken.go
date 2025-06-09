package middleware

import (
	"log"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"supercotacao/backend/utils"
)

func ValidadeToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		secret := os.Getenv("SECRET")
		if secret == "" {
			utils.Logger.Error("token inválido")
			context.JSON(401, gin.H{
				"erro": "token inválido",
				"code": 401,
			})
			context.Abort()
			return
		}

		authorization := context.GetHeader("Authorization")
		if authorization == "" {
			utils.Logger.Error("token jwt não fornecido")
			context.JSON(401, gin.H{
				"erro": "erro ao recuperar autenticação do usuário",
				"code": 401,
			})
			context.Abort()
			return
		}

		// Suporte a prefixo "Bearer "
		const bearerPrefix = "Bearer "
		if len(authorization) <= len(bearerPrefix) || authorization[:len(bearerPrefix)] != bearerPrefix {
			utils.Logger.Error("formato inválido para token JWT")
			context.JSON(401, gin.H{
				"erro": "formato inválido para autenticação. Use Bearer <token>",
				"code": 401,
			})
			context.Abort()
			return
		}

		tokenString := authorization[len(bearerPrefix):]

		auth, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !auth.Valid {
			utils.Logger.Error("token inválido ", err)
			context.JSON(401, gin.H{
				"erro": "autenticação inválida, realize o login novamente",
				"code": 401,
			})
			context.Abort()
			return
		}

		claims, ok := auth.Claims.(jwt.MapClaims)
		if !ok {
			utils.Logger.Error("Erro ao obter claims", err)
			context.JSON(400, gin.H{
				"erro": "erro interno ao obter parâmetros do usuário",
				"code": 400,
			})
			context.Abort()
			return
		}

		idStr, ok := claims["user_id"].(string)
		if !ok {
			log.Println("erro ao processar id_usuario no token JWT")
			context.JSON(400, gin.H{
				"erro": "erro ao processar id_usuario no token JWT",
			})
			context.Abort()
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Println("erro ao converter id_usuario para int64", err)
			context.JSON(400, gin.H{
				"erro": "erro ao converter id_usuario",
			})

			context.Abort()
			return
		}

		context.Set("user_id", id)
		context.Next()
	}
}
