package main

import (
	"github.com/Siddheshk02/Insta-share/lib"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	//app := fiber.New()

	engine := html.New("./index", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Static("/", "./")
	app.Static("/instashare", "./index.html")

	app.Get("/file-transfer", func(c *fiber.Ctx) error {
		return c.SendString("File Transfer!")
	})

	app.Post("/upload", lib.UploadFile)

	app.Listen(":3000")
}
