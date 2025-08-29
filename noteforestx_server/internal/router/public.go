package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (this *RouterInstance) RegisterPublicRoutes(v1 *gin.RouterGroup) {
	public := v1.Group("public")

	public.GET("test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
}
