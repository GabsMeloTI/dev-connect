package helper

import (
	"golang.org/x/crypto/chacha20poly1305"
	"log"
	"os"
)

func GetSignatureString() string {
	key := os.Getenv("SIGNATURE_KEY")
	if len(key) != chacha20poly1305.KeySize {
		log.Fatal("a chave de assinatura deve ter exatamente 32 caracteres")
	}
	return key
}
