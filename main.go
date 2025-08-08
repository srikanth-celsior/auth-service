package main

import (
	"auth-service/routers"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	_ = godotenv.Load()
	app := iris.New()
	routers.RegisterRoutes(app)
	_ = app.Listen(":3002")
}
