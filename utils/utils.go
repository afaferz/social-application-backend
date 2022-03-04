package utils

import (
	"os"
)

func GetSecret() string {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		//That's surely a big secret this way...
		secret = "sdmalncnjsdsmf"
	}
	return secret
}
