package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
var DBCLient *sql.DB
//ConectaDB conecta com o bando de dados
func ConectaDB(){
	stringConexao := "golang:golang@/cadastroSimples?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil{
		panic(erro.Error())
	}

	if erro = db.Ping(); erro != nil{
		panic(erro.Error())
	}

	DBCLient = db
	fmt.Println(db)
}