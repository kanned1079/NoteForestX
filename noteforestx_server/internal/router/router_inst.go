package router

import (
	"noteforestx_server/internal/middlewares"
	"noteforestx_server/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type RouterInstance struct {
	Id     int32
	Router *gin.Engine
	Mode   string
	util   utils.Utils
}

func NewRouterInstance(id int32, ginMode string) *RouterInstance {
	gin.SetMode(ginMode)
	return &RouterInstance{
		Id:     id,
		Router: gin.Default(),
		Mode:   ginMode,
	}
}

func (this *RouterInstance) RegisterApiServices() {
	this.util.Logger.PrintInfo("start api services")
	this.Router.Use(middlewares.AllowRequestTypeCors())

	apiPrefix := this.Router.Group("/api")
	v1 := apiPrefix.Group("v1")

	//this.RegisterPublicRoutes(apiPrefix)
	//this.RegisterAdminRoutes(apiPrefix)
	//this.RegisterUserRoutes(apiPrefix)
	//this.RegisterArticleRoutes(apiPrefix)

	this.RegisterPublicRoutes(v1)
	this.RegisterAdminRoutes(v1)
	this.RegisterUserRoutes(v1)
	this.RegisterArticleRoutes(v1)

	//this.RegisterIllustrationRoutes(v1)
	//this.RegisterPixivRoutes(v1)

}

func (this *RouterInstance) StartAndServe(addr string) {
	if err := this.Router.Run(addr); err != nil {
		this.util.Logger.PrintError("failed to start api services", err)
		os.Exit(1)
	}
}
