package router

import (
	"github.com/gin-gonic/gin"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/middlewares"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin"
)

/*
	{
	            "id": "190d874d-2744-4cef-ae68-2dbf1628e7e2",
	            "name": "オリジナル10000users入り",
	        },
	        {
	            "id": "e5b09780-771f-4112-bec9-528255091da9",
	            "name": "夏服",
	        },
	        {
	            "id": "c149c73e-771a-443e-9f70-ea6045cc07a6",
	            "name": "カメラ女子",
	        },
	        {
	            "id": "57d70b5a-83bd-46b0-b913-9bfd642e46ce",
	            "name": "花と女の子",
	        },
	        {
	            "id": "cc9f849c-4170-4be1-9518-beb4492613a0",
	            "name": "女の子",
	        }
*/

func (this *RouterInstance) RegisterAdminRoutes(v1 *gin.RouterGroup) {
	//adminRouter := v1.Group("/admin", middlewares.RequireRole("admin"))
	adminRouter := v1.Group("/admin", middlewares.RequireAuth(), middlewares.RequireRole(string(models.UserRoleAdmin)))
	var adminService = &admin.AdminService{
		Db:  dao.ExistingDbDaoInst.DbDao,
		Rdb: dao.ExistingDbDaoInst.RdbDao,
	}

	//var illustratonService = &illustration.IllustrationService{
	//	Db:  dao.ExistingDbDaoInst.DbDao,
	//	Rdb: dao.ExistingDbDaoInst.RdbDao,
	//}

	//adminRouter.GET("/document", adminService.GetDocumentList)
	//adminRouter.GET("document/:id", adminService.GetDocumentById)
	//adminRouter.PUT("document/:id", adminService.UpdateDocumentById)
	//adminRouter.POST("document", adminService.AddNewDocument)
	//adminRouter.DELETE("document/:id", adminService.RemoveDocumentById)

	//adminRouter.GET("/illustration/:id", illustratonService.GetIllustrationById)
	////adminRouter.GET("/illustration", adminService.GetIllustrationList)
	//adminRouter.PUT("/illustration/:id", adminService.UpdateIllustrationById)
	//adminRouter.POST("/illustration", adminService.AddNewIllustration)
	//adminRouter.DELETE("/illustration/:id", adminService.RemoveIllustrationById)
	//
	//adminRouter.GET("/illustration_tag", adminService.GetIllustrationTagList)
	//adminRouter.GET("/illustration_tag/:id", adminService.GetIllustrationTagById)
	//adminRouter.PUT("/illustration_tag/:id", adminService.UpdateIllustrationTagById)
	//adminRouter.POST("/illustration_tag", adminService.AddNewIllustrationTag)
	//adminRouter.DELETE("/illustration_tag/:id", adminService.RemoveIllustrationTagById)
	//
	//// sample c76f9e88-d35a-445e-8832-cf4bca4b633d
	//adminRouter.GET("/illustration_author", adminService.GetIllustrationAuthorList)
	//adminRouter.POST("/illustration_author", adminService.AddNewIllustrationAuthor)
	//adminRouter.PUT("/illustration_author/:id", adminService.UpdateIllustrationAuthorById)
	//adminRouter.DELETE("/illustration_author/:id", adminService.RemoveIllustrationAuthorById)

	adminRouter.GET("/tag", adminService.GetTagsSearchList)

	adminRouter.GET("/article/:id", adminService.GetArticleById)
	adminRouter.GET("/article", adminService.GetArticleList)
	adminRouter.POST("/article", adminService.AddArticle)
	adminRouter.PUT("/article/:id", adminService.UpdateArticleById)
	adminRouter.PATCH("/article/:id", adminService.PatchArticleById)
	adminRouter.DELETE("/article/:id", adminService.RemoveArticleById)

}
