package utils

import "noteforestx_server/internal/config"

type Logger struct {
}

type Utils struct {
	Logger               Logger
	jwtSecret            []byte
	accessTokenExpiredIn int
}

func NewUtils() *Utils {
	return &Utils{
		Logger:               Logger{},
		jwtSecret:            []byte(config.ExistingAppConfig.Runtime.JwtSecret),
		accessTokenExpiredIn: config.ExistingAppConfig.Runtime.AccessTokenExpiredIn,
	}
}
