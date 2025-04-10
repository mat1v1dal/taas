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
	User          models.User     `json:"user"`
	DriverLicense string          `json:"driver_license"`
	Vehicle       *models.Vehicle `json:"vehicle"`
}

func (h *UserHandler) RegisterDriver(c *gin.Context) {
	var req DriverRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := h.UserService.RegisterDriver(&req.User, req.DriverLicense, req.Vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar conductor"})
		return
	}

	c.JSON(http.StatusCreated, req.User)
}
