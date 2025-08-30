package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"
	"strings"
)

// doc process crud

// preProcessNewDocumentData frontend process doc data
func (this *AdminService) preProcessNewDocumentData(ctx *gin.Context) (reqData *dto.AddOrUpdateNewDocumentRequestDto, httpCode int, err error) {
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		return nil, http.StatusBadRequest, err
	}
	reqData.Title = strings.TrimSpace(reqData.Title)
	reqData.SubTitle = strings.TrimSpace(reqData.SubTitle)
	if len(reqData.Title) <= 0 || len(reqData.Content) <= 0 {
		return nil, http.StatusBadRequest, errors.New("the title and content cannot be empty")
	}

	this.utils.Logger.PrintInfo(reqData)
	return reqData, http.StatusOK, nil
}

// 新增 Document
func (this *AdminService) AddNewDocument(ctx *gin.Context) {
	req, code, err := this.preProcessNewDocumentData(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	if result := dao.ExistingDbDaoInst.DbDao.Model(&models.Document{}).Where("title = ?", req.Title); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "this document has already existed: " + req.Title,
		})
		return
	}

	newDoc := models.Document{
		Id:       uuid.New(),
		Title:    req.Title,
		Subtitle: req.SubTitle,
		Category: req.Category,
		Content:  req.Content,
		ImageURL: req.ImageUrl,
		Show:     false, // 默认不展示
		Pined:    false,
	}

	if err := dao.ExistingDbDaoInst.DbDao.Create(&newDoc).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newDoc)
}

// 更新 Document
func (this *AdminService) UpdateDocumentById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	req, code, err := this.preProcessNewDocumentData(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	var doc models.Document
	if err := dao.ExistingDbDaoInst.DbDao.First(&doc, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	doc.Title = req.Title
	doc.Subtitle = req.SubTitle
	doc.Category = req.Category
	doc.Content = req.Content
	doc.ImageURL = req.ImageUrl
	//doc.UpdatedAt = &time.Time{}

	if err := dao.ExistingDbDaoInst.DbDao.Save(&doc).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, doc)
}

// 获取 Document 列表（分页 + 搜索）
func (this *AdminService) GetDocumentList(ctx *gin.Context) {
	var req dto.GetDocumentListRequestDto
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

	var docs []models.Document
	query := dao.ExistingDbDaoInst.DbDao.Model(&models.Document{})

	if req.Search != "" {
		query = query.Where("title LIKE ?", "%"+req.Search+"%")
	}

	var total int64
	query.Count(&total)

	if err := query.Order("created_at desc").
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Find(&docs).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  docs,
		"total": total,
		"page":  req.Page,
		"size":  req.Size,
	})
}

// 获取 Document by ID
func (this *AdminService) GetDocumentById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var doc models.Document
	if err := dao.ExistingDbDaoInst.DbDao.First(&doc, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, doc)
}

// 删除 Document
func (this *AdminService) RemoveDocumentById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := dao.ExistingDbDaoInst.DbDao.Delete(&models.Document{}, "id = ?", uid).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "document deleted"})
}
