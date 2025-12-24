package router

import (
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/services/article"

	"github.com/gin-gonic/gin"
)

func (this *RouterInstance) RegisterArticleRoutes(v1 *gin.RouterGroup) {
	articleRouter := v1.Group("/article")
	articleService := &article.ArticleService{
		Db:  dao.ExistingDbDaoInst.DbDao,
		Rdb: dao.ExistingDbDaoInst.RdbDao,
	}

	articleRouter.GET("/", articleService.GetArticleList)
	articleRouter.GET("/:id", articleService.GetArticleById)

	//illustrationRouter.GET("/", illustrationService.GetIllustrationList)
	//illustrationRouter.GET("/:id", illustrationService.FetchIllustrationByIllId)
	//illustrationRouter.GET("/file/:file_name", illustrationService.FetchIllustrationByFilename)
	//
	//illustrationRouter.GET("/author/:author_id", illustrationService.FetchIllustrationsByAuthorId)
}
