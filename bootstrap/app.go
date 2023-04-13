package bootstrap

import (
	"log"
	"os"

	"github.com/RianIhsan/fullstack-goreact/database/migrations"
	"github.com/RianIhsan/fullstack-goreact/database/storage"
	"github.com/RianIhsan/fullstack-goreact/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func InitApp(app *fiber.App) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storage.ConnectDB(config)
	if err != nil {
		log.Fatal("Tidak bisa mengakses Database")
	}
	err = migrations.MigrateUsers(db)
	if err != nil {
		log.Fatal("Data gagal dimigrasikan ke database")
	}
	repo := repository.Repository{
		DB: db,
	}
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	repo.SetupRoute(app)
	app.Listen(":8080")
}
