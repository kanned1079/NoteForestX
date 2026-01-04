package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/models"
	"strings"
	"time"
)

type AccessTokenClaims struct {
	UserID uuid.UUID       `json:"uid"`
	Role   models.UserRole `json:"role"`
	jwt.RegisteredClaims
}

//func (this *Utils) GenerateAccessToken(user models.User) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"id":   user.Id,
//		"role": user.Role,
//		"exp":  time.Now().Add(time.Hour * time.Duration(config.ExistingAppConfig.Runtime.AccessTokenExpiredIn)).Unix(), // 6 小时有效期
//	})
//	//this.Logger.PrintInfo("jwt secret: ", config.AppCfg.Runtime.JwtSecret)
//
//	return token.SignedString([]byte(config.ExistingAppConfig.Runtime.JwtSecret))
//}

func (u *Utils) GenerateAccessToken(user models.User) (string, error) {
	now := time.Now()

	claims := AccessTokenClaims{
		UserID: user.Id,   // uuid.UUID
		Role:   user.Role, // models.UserRole（建议是 string enum）
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(
				now.Add(time.Hour * time.Duration(
					config.ExistingAppConfig.Runtime.AccessTokenExpiredIn,
				)),
			),
			// 可选，但推荐
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(
		[]byte(config.ExistingAppConfig.Runtime.JwtSecret),
	)
}

//func (this *Utils) ExtractTokenClaims(ctx *gin.Context) (jwt.MapClaims, error) {
//	authHeader := ctx.GetHeader("Authorization")
//	//this.Logger.PrintInfo(authHeader)
//	if authHeader == "" {
//		return nil, errors.New("缺少 Authorization 头")
//	}
//	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
//
//	// 使用 ParseWithClaims 明确指定 MapClaims
//	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		//this.Logger.PrintInfo("jwt secret: ", config.AppCfg.Runtime.JwtSecret)
//		return []byte(config.ExistingAppConfig.Runtime.JwtSecret), nil
//	})
//	if err != nil || !token.Valid {
//		return nil, errors.New("无效的 Token")
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return nil, errors.New("无法解析 Token Claims")
//	}
//	//log.Println(claims)
//	return claims, nil
//}

func (u *Utils) ExtractAccessTokenClaims(ctx *gin.Context) (*AccessTokenClaims, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("missing authorization header")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("invalid authorization header format")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 防止 alg 被篡改
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v",
					token.Header["alg"],
				)
			}

			return []byte(config.ExistingAppConfig.Runtime.JwtSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
