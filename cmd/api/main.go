package main

import (
	"log"
	"taas/internal/db"
	"taas/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Iniciando servidor Terapia-as-a-Service...")

	db.InitDB()

	r := gin.Default()

	// Handlers de usuario
	userHandler := handlers.NewUserHandler()

	r.POST("/api/users", userHandler.CreateUser)
	r.POST("/api/users/passenger", userHandler.RegisterPassenger)
	r.POST("/api/users/driver", userHandler.RegisterDriver)

	// Iniciar servidor
	r.Run(":8080")
}
