package main

import (
	"log"
	"taas/internal/db"
	"taas/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Iniciando servidor Terapia-as-a-Service...")

	// Conexión a base de datos
	db.InitDB()

	r := gin.Default()

	// Rutas públicas
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Rutas protegidas (a futuro, usar middleware JWT)
	r.GET("/me", handlers.Me)

	// Iniciar servidor
	r.Run(":8080")
}
