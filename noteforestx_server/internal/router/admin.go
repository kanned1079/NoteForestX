package router

import (
	"github.com/gin-gonic/gin"
	"noteforestx_server/internal/middlewares"
	"noteforestx_server/internal/services/admin"
)

func (this *RouterInstance) RegisterAdminRoutes(v1 *gin.RouterGroup) {
	adminRouter := v1.Group("/admin", middlewares.RequireRole("admin"))
	var adminService admin.AdminService

	adminRouter.GET("/document", adminService.GetDocumentList)
	adminRouter.GET("document/:id", adminService.GetDocumentById)
	adminRouter.PUT("document/:id", adminService.UpdateDocumentById)
	adminRouter.POST("document", adminService.AddNewDocument)
	adminRouter.DELETE("document/:id", adminService.RemoveDocumentById)

	adminRouter.GET("/illustration", adminService.GetIllustrationList)
	adminRouter.GET("illustration/:id", adminService.GetIllustrationById)
	adminRouter.PUT("illustration/:id", adminService.UpdateIllustrationById)
	adminRouter.POST("illustration", adminService.AddNewIllustration)
	adminRouter.DELETE("illustration/:id", adminService.RemoveIllustrationById)

	//adminRouter.GET("/illustration_tag", adminService.GetIllustrationTagList)
	//adminRouter.GET("illustration_tag/:id", adminService.GetIllustrationTagById)
	//adminRouter.PUT("illustration_tag/:id", adminService.UpdateIllustrationTagById)
	//adminRouter.POST("illustration_tag", adminService.AddNewIllustrationTag)
	//adminRouter.DELETE("illustration_tag/:id", adminService.RemoveIllustrationTagById)

}
