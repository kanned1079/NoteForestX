package article

import (
	"errors"
	"net/http"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetArticleList 获取文章列表（分页 + 搜索 + 标签 + 状态）
// published ONLY
func (s *ArticleService) GetArticleList(ctx *gin.Context) {
	var req dto.GetArticleListRequestDto
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 默认分页
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	db := dao.ExistingDbDaoInst.DbDao.
		Model(&models.Article{}).
		Preload("Tags")

	// ========== 状态过滤 ==========
	//if req.Status != nil && *req.Status != "" {
	//	db = db.Where("status = ?", *req.Status)
	//}

	db = db.Where("status = ?", "published")

	// ========== 标签 ID 过滤（最高优先级） ==========
	if req.TagId != nil && *req.TagId != "" {
		subQuery := dao.ExistingDbDaoInst.DbDao.
			Model(&models.ArticleTagRel{}).
			Select("article_id").
			Where("tag_id = ?", *req.TagId)

		// ⚠️ 子查询必须用 IN (?)
		db = db.Where("id IN (?)", subQuery)

	} else if req.Tag != nil && *req.Tag != "" {
		// ========== 标签名模糊搜索 ==========
		var tagIDs []string
		if err := dao.ExistingDbDaoInst.DbDao.
			Model(&models.ArticleTag{}).
			Where("name LIKE ?", "%"+*req.Tag+"%").
			Pluck("id", &tagIDs).Error; err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(tagIDs) == 0 {
			// 没有匹配标签，直接返回空
			ctx.JSON(http.StatusOK, gin.H{
				"list":  []models.Article{},
				"total": 0,
				"page":  req.Page,
				"size":  req.Size,
			})
			return
		}

		subQuery := dao.ExistingDbDaoInst.DbDao.
			Model(&models.ArticleTagRel{}).
			Select("article_id").
			Where("tag_id IN ?", tagIDs)

		db = db.Where("id IN (?)", subQuery)

	} else if req.Search != nil && *req.Search != "" {
		// ========== 标题模糊搜索 ==========
		db = db.Where("title LIKE ?", "%"+*req.Search+"%")
	}

	// ========== 查询总数 ==========
	var total int64
	if err := db.Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ========== 分页查询 ==========
	var articles []models.Article
	if err := db.
		Order("top DESC, created_at DESC").
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Select("id, slug, title, top, created_at, updated_at").
		Find(&articles).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  articles,
		"total": total,
		"page":  req.Page,
		"size":  req.Size,
	})
}

// GetArticleById 根据 ID 获取文章
func (s *ArticleService) GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var article models.Article
	if err := dao.ExistingDbDaoInst.DbDao.Preload("Tags").First(&article, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":      uid,
		"message": "success",
		"article": article,
	})
}
