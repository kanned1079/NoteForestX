package router

import (
	"noteforestx_server/internal/services/public"

	"github.com/gin-gonic/gin"
)

func (this *RouterInstance) RegisterPublicRoutes(v1 *gin.RouterGroup) {
	publicRouter := v1.Group("/public")
	var publicService public.PublicService

	publicRouter.POST("/user/login", publicService.UserLogin)
	//publicRouter.POST("/user/register", publicService.UserRegister)

}
