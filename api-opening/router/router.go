package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Inicialize() {
	//Inicializa o route usando as configurações default do GIN
	router := gin.Default()
	//Definindo um rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"ola":     "eai lindo",
		})
	})

	//Roda o servidor
	router.Run(":3000") //por padrão ela usa a rota 8080
}
