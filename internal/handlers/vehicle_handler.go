package handlers

import (
	"taas/internal/models"
	"taas/internal/services"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	VehicleService *services.VehicleService
}

func NewVehicleHandler() *VehicleHandler {
	return &VehicleHandler{
		VehicleService: &services.VehicleService{},
	}
}

// CreateVehicle godoc
// @Summary Crear vehículo
// @Description Crea un nuevo vehículo en el sistema
// @Tags Vehículos
// @Accept json
// @Produce json
// @Param vehicle body models.Vehicle true "Datos del vehículo"
// @Success 201 {object} models.Vehicle
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /vehicles [post]
func (h *VehicleHandler) CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if err := h.VehicleService.CreateVehicle(&vehicle); err != nil {
		c.JSON(500, gin.H{"error": "Error creating vehicle"})
		return
	}

	c.JSON(201, vehicle)
}
