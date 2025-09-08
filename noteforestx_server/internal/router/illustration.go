package router

import (
	"github.com/gin-gonic/gin"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/services/illustration"
)

func (this *RouterInstance) RegisterIllustrationRoutes(v1 *gin.RouterGroup) {
	illustrationRouter := v1.Group("/illustration")
	illustrationService := &illustration.IllustrationService{
		Db:  dao.ExistingDbDaoInst.DbDao,
		Rdb: dao.ExistingDbDaoInst.RdbDao,
	}

	illustrationRouter.GET("", illustrationService.GetIllustrationList)
	illustrationRouter.GET("/file/:file_name", illustrationService.FetchIllustrationByFilename)

}
