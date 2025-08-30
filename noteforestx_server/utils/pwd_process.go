package utils

import (
	"golang.org/x/crypto/bcrypt"
	"noteforestx_server/internal/config"
)

func (this *Utils) HashPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.ExistingAppConfig.Runtime.BcryptCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
