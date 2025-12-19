package admin

import (
	"errors"
	"net/http"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// preProcessArticleData 前端请求预处理
func (s *AdminService) preProcessArticleData(ctx *gin.Context) (reqData *dto.AddOrUpdateArticleRequestDto, httpCode int, err error) {
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, http.StatusBadRequest, err
	}
	reqData.Title = strings.TrimSpace(reqData.Title)
	reqData.Content = strings.TrimSpace(reqData.Content)
	if len(reqData.Title) == 0 || len(reqData.Content) == 0 {
		return nil, http.StatusBadRequest, errors.New("title and content cannot be empty")
	}
	s.utils.Logger.PrintInfo(reqData)
	return reqData, http.StatusOK, nil
}

// AddArticle 新增文章
func (s *AdminService) AddArticle(ctx *gin.Context) {
	req, code, err := s.preProcessArticleData(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"message": err.Error()})
		return
	}

	// 检查标题是否存在
	var exist models.Article
	if err := dao.ExistingDbDaoInst.DbDao.Where("title = ?", req.Title).First(&exist).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "article already exists: " + req.Title})
		return
	}

	// 处理标签
	var tags []*models.ArticleTag
	for _, t := range req.Tags {
		var tag models.ArticleTag
		if err := dao.ExistingDbDaoInst.DbDao.Where("name = ?", t).First(&tag).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tag = models.ArticleTag{
					ID:   uuid.New().String(),
					Name: t,
				}
				dao.ExistingDbDaoInst.DbDao.Create(&tag)
			}
		}
		tags = append(tags, &tag)
	}

	newArticle := models.Article{
		ID:      uuid.New().String(),
		Slug:    req.Slug,
		Title:   req.Title,
		Top:     req.Top,
		Status:  models.ArticleStatus(req.Status),
		Content: req.Content,
		Tags:    tags,
	}

	if err := dao.ExistingDbDaoInst.DbDao.Create(&newArticle).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newArticle)
}

// UpdateArticleById 更新文章
func (s *AdminService) UpdateArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req, code, err := s.preProcessArticleData(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"message": err.Error()})
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

	article.Title = req.Title
	article.Slug = req.Slug
	article.Top = req.Top
	article.Status = models.ArticleStatus(req.Status)
	article.Content = req.Content

	// 更新标签
	var tags []*models.ArticleTag
	for _, t := range req.Tags {
		var tag models.ArticleTag
		if err := dao.ExistingDbDaoInst.DbDao.Where("name = ?", t).First(&tag).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tag = models.ArticleTag{
					ID:   uuid.New().String(),
					Name: t,
				}
				dao.ExistingDbDaoInst.DbDao.Create(&tag)
			}
		}
		tags = append(tags, &tag)
	}
	// 替换多对多关系
	dao.ExistingDbDaoInst.DbDao.Model(&article).Association("Tags").Replace(tags)

	if err := dao.ExistingDbDaoInst.DbDao.Save(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// GetArticleList 获取文章列表（分页 + 搜索 + 标签 + 状态）
func (s *AdminService) GetArticleList(ctx *gin.Context) {
	var req dto.GetArticleListRequestDto
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	var articles []models.Article
	query := dao.ExistingDbDaoInst.DbDao.Preload("Tags").Model(&models.Article{})

	if req.Search != "" {
		query = query.Where("title LIKE ?", "%"+req.Search+"%")
	}
	if req.Tag != "" {
		query = query.Joins("JOIN article_tags at ON at.article_id = articles.id").
			Where("at.name = ?", req.Tag)
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	var total int64
	query.Count(&total)

	if err := query.Order("created_at desc").
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
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
func (s *AdminService) GetArticleById(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, article)
}

// RemoveArticleById 删除文章（软删除）
func (s *AdminService) RemoveArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := dao.ExistingDbDaoInst.DbDao.Delete(&models.Article{}, "id = ?", uid).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "article deleted"})
}
