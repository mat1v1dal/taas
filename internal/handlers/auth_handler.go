package handlers

import (
	"net/http"
	"taas/internal/db"
	"taas/internal/models"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// En un caso real, deberías hashear la contraseña con bcrypt
	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.Password,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado", "user": user})
}

func Login(c *gin.Context) {
	// Esta función será implementada luego con autenticación
	c.JSON(http.StatusOK, gin.H{"message": "Login placeholder"})
}

func Me(c *gin.Context) {
	// Esta función será protegida por JWT en el futuro
	c.JSON(http.StatusOK, gin.H{"message": "Perfil placeholder"})
}
