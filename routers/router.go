package routers

import (
	"auth-service/models"
	"auth-service/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application) {
	app.Post("/login", loginHandler)
}

func loginHandler(ctx iris.Context) {
	var req struct {
		UserID string `json:"user_id"`
	}
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StopWithStatus(400)
		return
	}
	token, err := GenerateJWT(req.UserID)
	if err != nil {
		ctx.StopWithStatus(500)
		return
	}
	ctx.JSON(iris.Map{"token": token})
}

func GenerateJWT(userId string) (string, error) {
	claims := models.CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 96)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecretKey, err := utils.GetSecret("JWT_SECRET", os.Getenv("PROJECT_ID"))
	if err != nil {
		return "", err
	}
	return token.SignedString([]byte(jwtSecretKey))
}
