package middlewares

import (
	"net/http"
	"noteforestx_server/internal/models"
	"noteforestx_server/utils"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = utils.CtxUserIDKey // "auth.userId"
const CtxUserRoleKey = "auth.userRole"

func RequireAuth() gin.HandlerFunc {
	var u utils.Utils
	return func(ctx *gin.Context) {
		claims, err := u.ExtractAccessTokenClaims(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "未授权: " + err.Error()})
			ctx.Abort()
			return
		}

		// 放入 ctx，供后续 handler / util 使用
		ctx.Set(CtxUserIDKey, claims.UserID)
		ctx.Set(CtxUserRoleKey, claims.Role)

		ctx.Next()
	}
}

// 只允许指定角色访问
func RequireRole(requiredRoles ...string) gin.HandlerFunc {
	var u utils.Utils
	return func(ctx *gin.Context) {
		claims, err := u.ExtractAccessTokenClaims(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "未授权: " + err.Error()})
			ctx.Abort()
			return
		}

		roleClaim := claims.Role
		for _, r := range requiredRoles {
			if roleClaim == models.UserRole(r) {
				ctx.Set(CtxUserIDKey, claims.UserID)
				ctx.Set(CtxUserRoleKey, roleClaim)
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, gin.H{"message": "无权限访问"})
		ctx.Abort()
	}
}
