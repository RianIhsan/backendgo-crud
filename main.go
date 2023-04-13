package main

import (
	"github.com/RianIhsan/fullstack-goreact/bootstrap"
	"github.com/RianIhsan/fullstack-goreact/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitApp(app)
}
