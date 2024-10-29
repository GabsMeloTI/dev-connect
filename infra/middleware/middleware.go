package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"treads/pkg/token"
)

func CheckAuthorization(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		bearerToken := c.Request().Header.Get("Authorization")
		tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

		maker, err := token.NewPasetoMaker(os.Getenv("TOKEN_SIGNATURE"))
		if err != nil {
			return c.JSON(http.StatusBadGateway, err.Error())
		}

		tokenPayload, err := maker.VerifyToken(tokenStr)
		if err != nil {
			return c.JSON(http.StatusBadGateway, err.Error())
		}
		c.Set("token_user_id", tokenPayload.UserID)
		c.Set("token_user_name", tokenPayload.Username)
		c.Set("token_email", tokenPayload.Email)
		c.Set("token_bio", tokenPayload.Bio)
		c.Set("token_avatar", tokenPayload.Avatar)
		c.Set("token_expiry_at", tokenPayload.ExpiredAt)

		return handlerFunc(c)
	}
}
