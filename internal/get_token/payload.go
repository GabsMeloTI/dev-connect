package get_token

import (
	"github.com/labstack/echo/v4"
	"time"
	"treads/internal/model"
)

func GetPayloadToken(c echo.Context) model.PayloadDTO {
	strUserID, _ := c.Get("token_user_id").(string)
	strUserName, _ := c.Get("token_user_name").(string)
	strExpiryAt, _ := c.Get("token_expiry_at").(time.Time)
	strEmail, _ := c.Get("token_email").(string)

	return model.PayloadDTO{
		UserID:       strUserID,
		UserNickname: strUserName,
		ExpiryAt:     strExpiryAt,
		Email:        strEmail,
	}
}
