package admin

import (
	"errors"
	"net/http"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (this *AdminService) isIllustrationTagUsed(tagId uuid.UUID) (bool, error) {
	var count int64
	err := this.Db.Model(&models.IllustrationTagMapping{}).
		Where("tag_id = ?", tagId).
		Count(&count).Error
	if err != nil {
		// 查询错误也返回 true，防止误删除
		return true, err
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

	ctx.JSON(http.StatusOK, gin.H{
		"tag":     tag,
		"message": "success",
	})
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
		req.Size = 20
	}

	type TagWithRelated struct {
		Id        uuid.UUID      `json:"id"`
		Name      string         `json:"name"`
		Related   int64          `json:"related"` // 關聯插畫數量
		CreatedAt *time.Time     `json:"created_at"`
		UpdatedAt *time.Time     `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`
	}

	// ==== 構建搜尋條件 ====
	search := ""
	if req.Search != "" {
		search = "%" + req.Search + "%"
	}

	// ==== 獲取總數 ====
	var total int64
	countQuery := this.Db.Model(&models.IllustrationTag{}).Where("deleted_at IS NULL")
	if search != "" {
		countQuery = countQuery.Where("name LIKE ?", search)
	}
	if err := countQuery.Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// ==== 如果請求要帶關聯數 ====
	if req.Related {
		tagList := make([]TagWithRelated, 0) // ✅ 确保即使无结果也返回 []
		query := this.Db.
			Table("x_illustration_tag AS t").
			Select(`
				t.id,
				t.name,
				t.created_at,
				t.updated_at,
				t.deleted_at,
				COUNT(m.illustration_id) AS related
			`).
			Joins(`LEFT JOIN x_illustration_tag_mapping AS m 
					ON t.id = m.tag_id AND m.deleted_at IS NULL`).
			Where(`t.deleted_at IS NULL`)

		// 搜尋條件
		if search != "" {
			query = query.Where("t.name LIKE ?", search)
		}

		err := query.
			Group("t.id").
			Order("t.created_at DESC").
			Offset((req.Page - 1) * req.Size).
			Limit(req.Size).
			Scan(&tagList).Error

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"list":  tagList,
			"total": total,
			"page":  req.Page,
			"size":  req.Size,
		})
		return
	}

	// ==== 普通標籤列表 ====
	tags := make([]models.IllustrationTag, 0) // ✅ 确保即使无结果也返回 []
	query := this.Db.Model(&models.IllustrationTag{}).Where("deleted_at IS NULL")
	if search != "" {
		query = query.Where("name LIKE ?", search)
	}
	if err := query.
		Order("created_at DESC").
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Find(&tags).Error; err != nil {
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

//func (this *AdminService) GetIllustrationTagList(ctx *gin.Context) {
//	var req dto.GetIllustrationTagListRequestDto
//	if err := ctx.ShouldBindQuery(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
//		return
//	}
//
//	if req.Page <= 0 {
//		req.Page = 1
//	}
//	if req.Size <= 0 {
//		req.Size = 20
//	}
//
//	type tagWithRelated struct {
//		Related   int64          `json:"related"` // 在原模型定義上加的關聯作品數量
//		Id        uuid.UUID      `json:"id"`
//		Name      string         `json:"name"`
//		CreatedAt *time.Time     `json:"created_at"`
//		UpdatedAt *time.Time     `json:"updated_at"`
//		DeletedAt gorm.DeletedAt `json:"deleted_at"`
//	}
//
//	if req.Related { // 如果是true則查詢該tag關聯了多少個插畫集
//
//	}
//
//	var tags []models.IllustrationTag
//	query := this.Db.Model(&models.IllustrationTag{})
//
//	if req.Search != "" {
//		search := "%" + req.Search + "%"
//		query = query.Where("name LIKE ?", search)
//	}
//
//	var total int64
//	if err := query.Count(&total).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
//		return
//	}
//
//	if err := query.Offset((req.Page - 1) * req.Size).Limit(req.Size).
//		Order("created_at desc").Find(&tags).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"list":  tags,
//		"total": total,
//		"page":  req.Page,
//		"size":  req.Size,
//	})
//}

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

	used, err := this.isIllustrationTagUsed(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if used {
		ctx.JSON(http.StatusConflict, gin.H{"message": "this tag is in use by one or more illustrations"})
		return
	}

	if err := this.Db.Delete(&models.IllustrationTag{}, "id = ?", uid).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
