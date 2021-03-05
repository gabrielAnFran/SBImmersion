package crud

import (
	"cadastroCidade/src/banco"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Cidade struct{
	ID 		string 		`json: "id"`
	Nome 	string		`json: "nome"`
	Uf 		string		`json: "uf"`
}
//CreateCity cria cidade
func CreateCity(c *gin.Context){

	fmt.Println("endpoint hit")
 	var reqBody Cidade
	if erro := c.ShouldBindJSON(&reqBody); erro != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : true,
			"message" : "erro ao receber corpo da requisição",
		})
		return
	}

	fmt.Println("endpoint hit 2")

	res, erro := banco.DBCLient.Exec("INSERT INTO usuario(nome, uf) VALUES (?, ?);",
reqBody.Nome, reqBody.Uf)

	if len(reqBody.Uf) != 2{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message" : "UF invalido",
		})
		
		return
	}
	if erro != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message" : "erro ao inserirr",
		})
	}

	id, erro := res.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{
			"error": false,
			"id" : id ,

	})

}

//GetAll retorna todas as cidades listadas
func GetAll(c *gin.Context){

	c.Query("uf")
	c.JSON(200, gin.H{
		"MESSAGE" :"GET ENDPOINT HIT",
	})

	var cidades [] Cidade

	linhas, erro := banco.DBCLient.Query("SELECT id, nome, uf FROM  usuario;" )
	if erro != nil{c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error" : true,
		"message" : "erro ao rjasonar",
	})
	return
}
	for linhas.Next(){
		var cidade Cidade
		if erro := linhas.Scan(&cidade.ID, &cidade.Nome, &cidade.Uf); erro != nil{
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error" : true,
				"message" : "erro ao receber corpo da for",
			})
			return
		}

		cidades = append(cidades, cidade)

	}

	c.JSON(http.StatusOK, cidades)

}
//GetOne retorna uma unica cidade
func GetOne(c *gin.Context){
	parametro := c.Param("id")
	id, _ := strconv.Atoi(parametro)
	row := banco.DBCLient.QueryRow("SELECT id, nome, uf FROM  usuario WHERE id =?;", id )

	var cidade Cidade
	if erro := row.Scan(&cidade.ID, &cidade.Nome, &cidade.Uf); erro !=nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : true,
			"message" : erro.Error(),
		})
		return
	}

c.JSON(http.StatusOK, cidade)
}

func PutOne(c *gin.Context){
	
	Parametro := c.Param("id")
	var reqBody Cidade

	id, _ := strconv.Atoi(Parametro)
	if erro := c.ShouldBindJSON(&reqBody); erro != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error" : true,
			"message" : "erro ao receber corpo da requisição",
		})
		return
	}

	_, erro := banco.DBCLient.Exec("UPDATE usuario set nome = ?, uf = ? WHERE id = ?",
reqBody.Nome, reqBody.Uf, id)
	if erro != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message" : "erro ao transformar",
		})
	}

}

func DeleteOne(c *gin.Context){
	parametro := c.Param("id")
	id, _ := strconv.Atoi(parametro)
	banco.DBCLient.QueryRow("DELETE FROM  usuario WHERE id =?;", id )
	c.JSON(http.StatusNoContent, id)
		
}

func GetCityByState(c *gin.Context){
	parametro := c.Param("uf")

	
	var cidades [] Cidade

	linhas, erro := banco.DBCLient.Query("SELECT id, nome, uf FROM  usuario WHERE uf= ?;", parametro )
	if erro != nil{c.JSON(http.StatusUnprocessableEntity, gin.H{
		"error" : true,
		"message" : "erro ao rjasonar",
	})
	return
}
	for linhas.Next(){
		var cidade Cidade
		if erro := linhas.Scan(&cidade.ID, &cidade.Nome, &cidade.Uf); erro != nil{
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error" : true,
				"message" : "erro ao receber corpo da for",
			})
			return
		}

		cidades = append(cidades, cidade)

	}

	c.JSON(http.StatusOK, cidades)

}
