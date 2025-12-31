package utils

import (
	"noteforestx_server/internal/config"

	"golang.org/x/crypto/bcrypt"
)

type Logger struct {
	running  bool
	stopChan chan struct{}
}

type Utils struct {
	Logger               Logger
	jwtSecret            []byte
	accessTokenExpiredIn int
}

func (this *Utils) CheckPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	return err == nil
}

func NewUtils() *Utils {
	return &Utils{
		Logger:               Logger{},
		jwtSecret:            []byte(config.ExistingAppConfig.Runtime.JwtSecret),
		accessTokenExpiredIn: config.ExistingAppConfig.Runtime.AccessTokenExpiredIn,
	}
}
