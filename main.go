package main

import (
	"auth-service/routers"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	_ = godotenv.Load()
	app := iris.New()
	routers.RegisterRoutes(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = app.Listen(":" + port)
}
