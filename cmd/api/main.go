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

	r.GET("/ping", handlers.Pong)

	// Iniciar servidor
	r.Run(":8080")
}
