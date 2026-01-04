package router

import (
	"github.com/gin-gonic/gin"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/middlewares"
	"noteforestx_server/internal/services/user"
)

func (this *RouterInstance) RegisterUserRoutes(v1 *gin.RouterGroup) {
	userRouter := v1.Group("/user", middlewares.RequireAuth())

	userServices := user.UserService{
		Db:  dao.ExistingDbDaoInst.DbDao,
		Rdb: dao.ExistingDbDaoInst.RdbDao,
	}

	userRouter.PATCH("/:id/username", userServices.UpdateUsername)

	userRouter.PATCH("/:id/password", userServices.UpdateUserPassword)

}
