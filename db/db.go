package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" 
)

func Conecta() *sql.DB {
	parametros := "root:root1234@tcp(localhost:3306)/erp"

	db, erro := sql.Open("mysql", parametros)
	if erro != nil {
		fmt.Println("Erro ao abrir a conexão:", erro)
	}

	return db
}