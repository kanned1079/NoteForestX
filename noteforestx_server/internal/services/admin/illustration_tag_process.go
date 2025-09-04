package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"
	"strings"
)

// doc process crud

func (this *AdminService) isIllustrationTagNameExists(name string) (bool, error) {
	var count int64
	if err := this.Db.Model(&models.IllustrationTag{}).
		Where("name = ?", name).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// AddNewIllustrationTag response a tag data
func (this *AdminService) AddNewIllustrationTag(ctx *gin.Context) {
	var req dto.AddNewIllustrationTagRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "name cannot be empty"})
		return
	}

	// 检查是否已存在同名标签
	exists, err := this.isIllustrationTagNameExists(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"message": "tag with the same name already exists"})
		return
	}

	// 创建新标签
	tag := models.IllustrationTag{
		Id:   uuid.New(),
		Name: name,
	}

	if err := this.Db.Create(&tag).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

// UpdateIllustrationTagById response a tag data
func (this *AdminService) UpdateIllustrationTagById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var req dto.UpdateIllustrationTagByIdRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "name cannot be empty"})
		return
	}

	exists, err := this.isIllustrationTagNameExists(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"message": "tag with the same name already exists"})
		return
	}

	var tag models.IllustrationTag
	if err := this.Db.First(&tag, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "tag not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	tag.Name = name
	if err := this.Db.Save(&tag).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (this *AdminService) GetIllustrationTagList(ctx *gin.Context) {
	var req dto.GetIllustrationTagListRequestDto
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	var tags []models.IllustrationTag
	query := this.Db.Model(&models.IllustrationTag{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		query = query.Where("name LIKE ?", search)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := query.Offset((req.Page - 1) * req.Size).Limit(req.Size).
		Order("created_at desc").Find(&tags).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  tags,
		"total": total,
		"page":  req.Page,
		"size":  req.Size,
	})
}

func (this *AdminService) GetIllustrationTagById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var tag models.IllustrationTag
	if err := this.Db.First(&tag, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "tag not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (this *AdminService) RemoveIllustrationTagById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	if err := this.Db.Delete(&models.IllustrationTag{}, "id = ?", uid).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
