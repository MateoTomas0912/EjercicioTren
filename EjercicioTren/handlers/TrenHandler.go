package handlers

import (
	"net/http"
	"src/EjercicioTren/dto"
	"src/EjercicioTren/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrenHandler struct {
	trenHandler *services.TrenService
}

func NewTrenHandler() *TrenHandler {
	return &TrenHandler{
		trenHandler: services.NewTrenService(),
	}
}

func (handler *TrenHandler) InsertarVagon(c *gin.Context) {
	idTren, err := strconv.Atoi(c.Param("idTren"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro idTren no numerico"})

		return
	}

	var vagon dto.Vagon
	err = c.BindJSON(&vagon)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el body"})

		return
	}

	tren := handler.trenHandler.InsertarVagon(idTren, vagon)
	if tren != nil {
		c.JSON(http.StatusOK, tren)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo insertar el vagon"})
	}
}

func (handler *TrenHandler) ObtenerSegunPeso(c *gin.Context) {
	pesoVagon, err := strconv.Atoi(c.Param("pesoVagon"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro peso no numerico"})

		return
	}

	tren := handler.trenHandler.ObtenerSegunPeso(pesoVagon)
	if tren != nil {
		c.JSON(http.StatusOK, tren)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el tren"})
	}
}

func (handler *TrenHandler) CalcularCosto(c *gin.Context) {
	// Obtener el ID de tren de los parámetros de ruta
	idTren, err := strconv.Atoi(c.Param("idTren"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro id no numerico"})

		return
	}

	// Obtener la cantidad de kilómetros de los queryParams
	kilometersStr := c.DefaultQuery("kilometros", "0")
	cantKm, err := strconv.Atoi(kilometersStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Kilometros no es entero")
		return
	}

	costo := handler.trenHandler.CalcularCosto(idTren, cantKm)
	if costo != 0 {
		c.JSON(http.StatusOK, costo)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se encontro ese tren"})
	}
}
