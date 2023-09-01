package main

import (
	"src/EjercicioTren/handlers"

	"github.com/gin-gonic/gin"
)

var (
	r             *gin.Engine
	trenesHandler *handlers.TrenHandler
)

func main() {
	r = gin.Default()

	iniciar()

	mapping()

	r.Run(":8080")
}

func iniciar() {
	trenesHandler = handlers.NewTrenHandler()
}

func mapping() {
	r.PUT("/trenes/:idTren", trenesHandler.InsertarVagon)
	r.GET("/trenes/:pesoVagon", trenesHandler.ObtenerSegunPeso)
	r.GET("/trenes/:id", trenesHandler.CalcularCosto) //se pasa como /:id?kilometros
}
