package helper

import (
	"os"
)

func GetSignatureString() string {
	return os.Getenv("TOKEN_SIGNATURE")
}
