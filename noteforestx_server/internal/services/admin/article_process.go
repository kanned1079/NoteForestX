package admin

import (
	"errors"
	"fmt"
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
func (s *AdminService) preProcessArticleData(ctx *gin.Context) (*dto.AddOrUpdateArticleRequestDto, int, error) {
	var reqData dto.AddOrUpdateArticleRequestDto
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, http.StatusBadRequest, err
	}

	reqData.Title = strings.TrimSpace(reqData.Title)
	reqData.Content = strings.TrimSpace(reqData.Content)
	if reqData.Title == "" || reqData.Content == "" {
		return nil, http.StatusBadRequest, errors.New("title and content cannot be empty")
	}

	//s.utils.Logger.PrintInfo(reqData)
	return &reqData, http.StatusOK, nil
}

// AddArticle 新建文章
func (s *AdminService) AddArticle(ctx *gin.Context) {
	req, code, err := s.preProcessArticleData(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"message": err.Error()})
		return
	}

	db := dao.ExistingDbDaoInst.DbDao

	// 使用事务保证一致性
	err = db.Transaction(func(tx *gorm.DB) error {
		// 检查标题是否存在
		var exist models.Article
		if err := tx.Where("title = ?", req.Title).First(&exist).Error; err == nil {
			return errors.New("article already exists: " + req.Title)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 处理标签
		var tags []*models.ArticleTag
		for _, t := range req.Tags {
			var tag models.ArticleTag

			if t.ID == "new" {
				// 先按 name 查，防止重复创建
				if err := tx.Where("name = ?", t.Name).First(&tag).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						tag = models.ArticleTag{
							ID:   uuid.New().String(),
							Name: t.Name,
						}
						if err := tx.Create(&tag).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			} else {
				if err := tx.First(&tag, "id = ?", t.ID).Error; err != nil {
					return err
				}
			}

			tags = append(tags, &tag)
		}

		// 创建文章
		newArticle := models.Article{
			ID:       uuid.New().String(),
			Slug:     req.Slug,
			Title:    req.Title,
			Top:      req.Top,
			Status:   models.ArticleStatus(req.Status),
			Content:  req.Content,
			Tags:     tags,
			ImageUrl: req.ImageUrl,
		}

		if err := tx.Create(&newArticle).Error; err != nil {
			return err
		}

		// 返回 JSON
		ctx.JSON(http.StatusOK, newArticle)
		return nil
	})

	if err != nil {
		if strings.HasPrefix(err.Error(), "article already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
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

	db := dao.ExistingDbDaoInst.DbDao

	// 使用事务保证一致性
	err = db.Transaction(func(tx *gorm.DB) error {
		var article models.Article
		if err := tx.Preload("Tags").First(&article, "id = ?", uid).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("article not found")
			}
			return err
		}

		// 更新基本字段
		article.Title = req.Title
		article.Slug = req.Slug
		article.Top = req.Top
		article.Status = models.ArticleStatus(req.Status)
		article.Content = req.Content
		article.ImageUrl = req.ImageUrl

		// 处理标签
		var tags []*models.ArticleTag
		for _, t := range req.Tags {
			var tag models.ArticleTag
			if t.ID == "new" {
				// 按 name 查，防止重复创建
				if err := tx.Where("name = ?", t.Name).First(&tag).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						tag = models.ArticleTag{
							ID:   uuid.New().String(),
							Name: t.Name,
						}
						if err := tx.Create(&tag).Error; err != nil {
							return err
						}
					} else {
						return err
					}
				}
			} else {
				// 已有 ID，检查是否存在
				if err := tx.First(&tag, "id = ?", t.ID).Error; err != nil {
					return err
				}
			}
			tags = append(tags, &tag)
		}

		// 替换多对多关系
		if err := tx.Model(&article).Association("Tags").Replace(tags); err != nil {
			return err
		}

		// 保存文章
		if err := tx.Save(&article).Error; err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, article)
		return nil
	})

	if err != nil {
		if err.Error() == "article not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

// GetArticleList 获取文章列表（分页 + 搜索 + 标签 + 状态）
func (s *AdminService) GetArticleList(ctx *gin.Context) {
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
	if req.Status != nil && *req.Status != "" {
		db = db.Where("status = ?", *req.Status)
	}

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

	// 9603 4603 3243

	// ========== 分页查询 ==========
	var articles []models.Article
	if err := db.
		Order("created_at DESC").
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Select("id, slug, title, top, status, created_at, updated_at").
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

	ctx.JSON(http.StatusOK, gin.H{
		"id":      uid,
		"message": "success",
		"article": article,
	})
}

func (AdminService) isValidArticleStatus(s string) bool {
	switch s {
	case "draft", "published", "hidden":
		return true
	default:
		return false
	}
}

func (s *AdminService) PatchArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.PatchArticleRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 至少要有一个字段
	if req.Top == nil && req.Status == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	updates := map[string]interface{}{}

	if req.Top != nil {
		updates["top"] = *req.Top
	}

	if req.Status != nil {
		if !s.isValidArticleStatus(*req.Status) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid status value",
			})
			return
		}
		updates["status"] = *req.Status
	}

	db := dao.ExistingDbDaoInst.DbDao

	if err := db.Model(&models.Article{}).
		Where("id = ?", uid).
		Updates(updates).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "article updated",
	})
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
