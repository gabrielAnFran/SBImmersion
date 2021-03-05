package main

import (
	"cadastroCidade/src/banco"
	"cadastroCidade/src/crud"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	banco.ConectaDB()

	r.POST("/cidades", crud.CreateCity)
	r.GET("/cidades", crud.GetAll )
	r.GET("/cidades/:id", crud.GetOne )
	r.GET("/estados/:uf", crud.GetCityByState)
	r.PUT("/cidades/:id",crud.PutOne)
	r.DELETE("/cidades/:id", crud.DeleteOne)
	

	if erro := r.Run(":8081"); erro != nil {
		log.Fatal(erro.Error())
	}

}
