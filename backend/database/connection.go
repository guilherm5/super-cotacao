package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println("Erro ao ler variaveis de ambiente para iniciar conexão com o banco de dados", err)
		return nil, errors.New("Erro iniciar conexão com banco de dados")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	stringConnectionDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=verify-full",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", stringConnectionDB)
	if err != nil {
		log.Println("Erro ao realizar conexão com o banco de dados", err)
		return nil, errors.New("Erro ao realizar a conexão com o banco de dados")
	}

	err = db.Ping()
	if err != nil {
		log.Println("Erro ao verificar conexão com o banco de dados", err)
		return nil, errors.New("Erro ao verificar conexão com o banco de dados")
	}
	log.Println("Conexão com o banco de dados realizada com sucesso")
	return db, nil
}
