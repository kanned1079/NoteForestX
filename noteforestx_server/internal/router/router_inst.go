package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"noteforestx_server/internal/middlewares"
	"noteforestx_server/utils"
	"os"
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

	this.RegisterPublicRoutes(v1)
	this.RegisterAdminRoutes(v1)

}

func (this *RouterInstance) StartAndServe(port string) {
	if err := this.Router.Run(fmt.Sprintf(":%s", port)); err != nil {
		this.util.Logger.PrintError("failed to start api services", err)
		os.Exit(1)
	}
}
