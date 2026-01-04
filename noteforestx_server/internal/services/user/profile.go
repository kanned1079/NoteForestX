package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/user/dto"
	"noteforestx_server/utils"
	"strings"
)

func (s *UserService) UpdateUsername(ctx *gin.Context) {
	// 1️⃣ 获取 URL 参数
	userIDParam := strings.TrimSpace(ctx.Param("id"))
	if userIDParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	// 2️⃣ 请求体绑定
	var req dto.UpdateUsernameRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	// 3️⃣ username 校验
	username := strings.TrimSpace(req.Username)
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "username cannot be empty",
		})
		return
	}
	if len(username) > 32 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "username too long",
		})
		return
	}

	// 4️⃣ 权限校验（确保 token 用户和 param 用户相同）
	if err := s.utils.EnsureSameUser(ctx, userIDParam); err != nil {
		switch err {
		case utils.ErrUnauthorized:
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		case utils.ErrForbidden:
			ctx.JSON(http.StatusForbidden, gin.H{"message": "permission denied"})
		case utils.ErrInvalidUserID:
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		}
		return
	}

	// 5️⃣ 查询用户是否存在
	userID, _ := uuid.Parse(userIDParam) // 已经在 EnsureSameUser 校验过，这里直接解析
	var user models.User
	if err := s.Db.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "database error",
		})
		return
	}

	// 6️⃣ 更新 username（单字段更新）
	if err := s.Db.Model(&user).Update("username", username).Error; err != nil {
		// 唯一索引冲突
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "username already exists",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update username",
		})
		return
	}

	// 7️⃣ 返回成功
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"new_username": username,
	})
}

// UpdateUserPassword PATCH: /api/v1/user/:id/password
func (s *UserService) UpdateUserPassword(ctx *gin.Context) {
	// 1️⃣ 解析 URL 参数（UUID）
	userIDParam := strings.TrimSpace(ctx.Param("id"))
	if userIDParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	// 2️⃣ 绑定请求体
	var req dto.UpdateUserPasswordRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	previousPassword := strings.TrimSpace(req.PreviousPassword)
	newPassword := strings.TrimSpace(req.NewPassword)

	if previousPassword == "" || newPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "password cannot be empty"})
		return
	}

	// 3️⃣ 校验新密码长度 / 强度
	if len(newPassword) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "new password must be at least 6 characters"})
		return
	}

	// 4️⃣ 权限校验
	if err := s.utils.EnsureSameUser(ctx, userIDParam); err != nil {
		switch err {
		case utils.ErrUnauthorized:
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		case utils.ErrForbidden:
			ctx.JSON(http.StatusForbidden, gin.H{"message": "permission denied"})
		case utils.ErrInvalidUserID:
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		}
		return
	}

	// 5️⃣ 查询用户是否存在
	userID, _ := uuid.Parse(userIDParam)
	var user models.User
	if err := s.Db.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	// 6️⃣ 验证旧密码
	if !s.utils.CheckPasswordHash(previousPassword, user.Password) {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "previous password incorrect"})
		return
	}

	// 7️⃣ 生成新密码哈希
	newHash, err := s.utils.HashPassword(newPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash password"})
		return
	}

	// 8️⃣ 更新数据库
	if err := s.Db.Model(&user).Update("password", newHash).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update password"})
		return
	}

	// 9️⃣ 返回成功
	ctx.JSON(http.StatusOK, gin.H{
		"message": "password updated successfully",
	})
}
