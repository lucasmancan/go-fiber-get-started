package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasmancan/go-get-started/src/usecases"
)

func main() {

	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Get("/api", GetCars)

	app.Listen(":3000")
}

func GetCars(c *fiber.Ctx) error {
	err := c.JSON(usecases.NewSaveUserUseCase().Array)
	return err
}

type Car struct {
	Id   int
	Name string
}
