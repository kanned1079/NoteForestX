package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("permission denied")
	ErrInvalidUserID = errors.New("invalid user id")
)

const CtxUserIDKey = "auth.userId"

func (u *Utils) EnsureSameUser(ctx *gin.Context, paramID string) error {
	// 1️⃣ 校验 param id
	paramID = strings.TrimSpace(paramID)
	if paramID == "" {
		return ErrInvalidUserID
	}

	targetID, err := uuid.Parse(paramID)
	if err != nil {
		return ErrInvalidUserID
	}

	// 2️⃣ 从 context 中取 token user id
	userIDAny, exists := ctx.Get(CtxUserIDKey)
	if !exists {
		return ErrUnauthorized
	}

	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		return ErrUnauthorized
	}

	// 3️⃣ 比较
	if userID != targetID {
		return ErrForbidden
	}

	return nil
}
