package admin

import (
	"net/http"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *AdminService) GetTagsSearchList(ctx *gin.Context) {
	// 从 query 获取 keyword
	keyword := ctx.Query("search")

	db := dao.ExistingDbDaoInst.DbDao.Model(&models.ArticleTag{})

	// 如果 keyword 不为空，模糊搜索标签名
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	// 查询总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查询列表，按创建时间倒序
	var tagList []models.ArticleTag
	if err := db.Order("created_at DESC").Find(&tagList).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total":   total,
		"tags":    tagList,
		"message": "success",
	})
}
