package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryohidaka/go-pixiv"
)

func (this *RouterInstance) RegisterPixivRoutes(v1 *gin.RouterGroup) {
	pixivRouter := v1.Group("/pixiv")
	//articleService := &article.ArticleService{
	//	Db:  dao.ExistingDbDaoInst.DbDao,
	//	Rdb: dao.ExistingDbDaoInst.RdbDao,
	//}

	pixivRouter.GET("/login", func(ctx *gin.Context) {
		this.util.Logger.PrintInfo("PIXIV Login")
		app, err := pixiv.NewApp("<YOUR_REFRESH_TOKEN>")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		user, _ := app.UserDetail(11)
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})

	})

	//illustrationRouter.GET("/", illustrationService.GetIllustrationList)
	//illustrationRouter.GET("/:id", illustrationService.FetchIllustrationByIllId)
	//illustrationRouter.GET("/file/:file_name", illustrationService.FetchIllustrationByFilename)
	//
	//illustrationRouter.GET("/author/:author_id", illustrationService.FetchIllustrationsByAuthorId)
}
