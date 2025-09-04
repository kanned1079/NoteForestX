package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (this *RouterInstance) RegisterUserRoutes(v1 *gin.RouterGroup) {
	userRouter := v1.Group("/user")

	userRouter.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "user",
		})
	})
}
