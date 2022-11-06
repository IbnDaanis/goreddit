package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ibndaanis/golang-rest/internal/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDataBase()
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Port
	port := ":" + os.Getenv("PORT")

	err := app.Listen(port)

	if err != nil {
		log.Fatal("There was an error with the server: ", err)
	}

}
