package handlers

import (
	"net/http"
	"taas/internal/models"
	"taas/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserService: &services.UserService{},
	}
}

// CreateUser godoc
// @Summary Crear un nuevo usuario
// @Description Crea un usuario base sin rol asignado (ni conductor ni pasajero)
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param user body models.User true "Datos del usuario"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// RegisterPassenger godoc
// @Summary Registrar pasajero
// @Description Registra un usuario y le asigna el rol de pasajero
// @Tags Pasajeros
// @Accept json
// @Produce json
// @Param user body models.User true "Datos del usuario pasajero"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/passenger [post]
func (h *UserHandler) RegisterPassenger(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := h.UserService.RegisterPassenger(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar pasajero"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

type DriverRegistrationRequest struct {
	User          models.User `json:"user"`
	DriverLicense string      `json:"driver_license"`
}

// RegisterDriver godoc
// @Summary Registrar conductor
// @Description Registra un usuario y le asigna el rol de conductor (sin vehículo)
// @Tags Conductores
// @Accept json
// @Produce json
// @Param driver body DriverRegistrationRequest true "Datos del conductor"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/driver [post]
func (h *UserHandler) RegisterDriver(c *gin.Context) {
	var req DriverRegistrationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := h.UserService.RegisterDriver(&req.User, req.DriverLicense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar conductor"})
		return
	}

	c.JSON(http.StatusCreated, req.User)
}

// AssignVehicleToDriver godoc
// @Summary Asignar vehículo a conductor
// @Description Asigna un vehículo existente a un conductor ya registrado
// @Tags Conductores
// @Accept json
// @Produce json
// @Param assignment body struct{DriverID uint `json:"driver_id"`; VehicleID uint `json:"vehicle_id"`} true "ID del conductor y del vehículo"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/driver/assign-vehicle [post]
type VehicleAssignmentRequest struct {
	DriverID  uint `json:"driver_id"`
	VehicleID uint `json:"vehicle_id"`
}

func (h *UserHandler) AssingVehicleToDriver(c *gin.Context) {
	var req VehicleAssignmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := h.UserService.RegisterVehicleToDriver(req.DriverID, req.VehicleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al asignar vehículo al conductor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vehículo asignado correctamente"})
}
