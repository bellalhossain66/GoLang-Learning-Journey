package main

import(
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"go-fiber-mvc-app/config"
	"go-fiber-mvc-app/router"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	router.AppRoute(app)

	log.Printf("==> 🌎  Listening on port %s. Visit http://localhost:%s/ in your browser.", os.Getenv("APP_PORT"), os.Getenv("APP_PORT"))
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Fail to start server %v", err)
	}
}