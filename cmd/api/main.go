package main

import (
	"log"
	"taas/internal/db"
	"taas/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "taas/cmd/api/docs" // Importa la documentación generada por swaggo
)

// @title TaaS API
// @version 1.0
// @description API de Terapia-as-a-Service para manejo de usuarios, conductores, pasajeros y vehículos.
// @host localhost:8080
// @BasePath /api
func main() {
	log.Println("Iniciando servidor Terapia-as-a-Service...")

	db.InitDB()

	r := gin.Default()

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handlers
	userHandler := handlers.NewUserHandler()
	vehicleHandler := handlers.NewVehicleHandler()

	// Rutas de usuario
	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.POST("/users/passenger", userHandler.RegisterPassenger)
		api.POST("/users/driver", userHandler.RegisterDriver)
		api.POST("/users/driver/assign-vehicle", userHandler.AssingVehicleToDriver)

		api.POST("/vehicles", vehicleHandler.CreateVehicle)
	}

	r.Run(":8080")
}
