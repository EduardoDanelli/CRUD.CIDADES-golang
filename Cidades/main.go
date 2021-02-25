package main

import (
	"teste/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/teste", controllers.CriarCidade)
	router.GET("/teste", controllers.ListarCidades)
	router.GET("/teste/:id", controllers.BuscarCidade)
	router.PUT("/teste/:id", controllers.EditarCidade)
	router.DELETE("/teste/:id", controllers.ExcluirCidade)

	router.Run(":8081")
}
