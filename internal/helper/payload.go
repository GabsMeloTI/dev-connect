package helper

import (
	"fmt"
	"reflect"
	"time"
	"treads/internal/model"

	"github.com/labstack/echo/v4"
)

func GetPayloadToken(c echo.Context) model.PayloadDTO {
	strUserID, _ := c.Get("token_user_id").(string)
	strUserName, _ := c.Get("token_user_name").(string)
	strName, _ := c.Get("token_user_name").(string)
	strUserEmail, _ := c.Get("token_user_name").(string)
	strExpiryAt, _ := c.Get("token_expiry_at").(time.Time)

	return model.PayloadDTO{
		UserID:   strUserID,
		Name:     strName,
		Username: strUserName,
		Email:    strUserEmail,
		ExpiryAt: strExpiryAt,
	}
}

func WhoMapper[T any](data T, payload model.PayloadDTO) (T, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return data, fmt.Errorf("input data must be a struct")
	}

	whoField := v.FieldByName("who")
	if !whoField.IsValid() {
		return data, fmt.Errorf("struct doesn't have a field named 'who'")
	}

	whoField.Set(reflect.ValueOf(payload.Username))

	return data, nil
}
