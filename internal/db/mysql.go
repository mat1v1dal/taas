package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"taas/internal/models"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	// Cargar .env si existe
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	// Variables de entorno
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, name)

	// Conexión
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = database

	// Migraciones
	err = DB.AutoMigrate(
		&models.User{},
		&models.Place{},
		&models.Trip{},
		&models.TripParticipation{},
		&models.Vehicle{},
		&models.UserPassenger{},
		&models.UserDriver{},
	)
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	log.Println("Database connected and migrated")
}
