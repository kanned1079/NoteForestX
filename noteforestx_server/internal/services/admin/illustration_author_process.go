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

// author process crud

func (this *AdminService) isAuthorHasExistingIllustration(uuid uuid.UUID) (bool, error) {
	var count int64
	if err := this.Db.Model(&models.Illustration{}).Where("author_id = ?", uuid).Count(&count).
		Error; err != nil {
		return true, err // 查找错误也返回true 防止误删除
	}
	return count > 0, nil
}

func (this *AdminService) isIllustrationAuthorExistsByUuid(uuid uuid.UUID) (bool, error) {
	var count int64
	if err := this.Db.Model(&models.IllustrationAuthor{}).
		Where("id = ?", uuid).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (this *AdminService) isIllustrationAuthorExistsByLink(link string, excludeId *uuid.UUID) (bool, error) {
	var count int64
	db := this.Db.Model(&models.IllustrationAuthor{}).Where("link = ?", link)
	if excludeId != nil {
		db = db.Where("id <> ?", *excludeId)
	}
	if err := db.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (this *AdminService) getAuthorInfoFromCtx(ctx *gin.Context) (err error, name, link string, related bool) {
	var req dto.AddNewIllustrationAuthorRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err, "", "", false
	}
	name = strings.TrimSpace(req.Name)
	if name == "" {
		return errors.New("name cannot be empty"), "", "", false
	}
	link = strings.TrimSpace(req.Link)
	if link == "" {
		return errors.New("link cannot be empty"), "", "", false
	}
	return
}

func (this *AdminService) AddNewIllustrationAuthor(ctx *gin.Context) {
	err, name, link, _ := this.getAuthorInfoFromCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 检查是否已存在相同画师
	exists, err := this.isIllustrationAuthorExistsByLink(link, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"message": "author already exists"})
		return
	}

	author := models.IllustrationAuthor{
		Id:   uuid.New(),
		Name: name,
		Link: link,
	}

	if err := this.Db.Create(&author).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (this *AdminService) UpdateIllustrationAuthorById(ctx *gin.Context) {
	uid, err := this.utils.GetAndParseParamUuid("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	// 2. 检查作者是否存在
	exists, err := this.isIllustrationAuthorExistsByUuid(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to query author"})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "author not found"})
		return
	}

	// 3. 从请求体获取 name 和 link
	err, name, link, _ := this.getAuthorInfoFromCtx(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 4. 检查 link 是否被其他作者占用
	linkExists, err := this.isIllustrationAuthorExistsByLink(link, &uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to check link uniqueness"})
		return
	}
	if linkExists {
		ctx.JSON(http.StatusConflict, gin.H{"message": "link already exists"})
		return
	}

	// 5. 更新作者信息
	updateData := map[string]interface{}{
		"name": name,
		"link": link,
	}
	if err := this.Db.Model(&models.IllustrationAuthor{}).
		Where("id = ?", uid).
		Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update author"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "author updated successfully",
	})
}

//func (this *AdminService) GetIllustrationAuthorList(ctx *gin.Context) {
//	var req dto.GetIllustrationAuthorListRequestDto
//	if err := ctx.ShouldBindQuery(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
//		return
//	}
//
//	type IllustrationAuthorWithRelated struct { // 如果req.Related是Ture就統計該作者所有插畫數量
//		Id        uuid.UUID      `json:"id"`
//		Name      string         `json:"name"`
//		Link      string         `json:"link"`
//		Related   int64          `json:"related"`
//		CreatedAt *time.Time     `json:"created_at"`
//		UpdatedAt *time.Time     `json:"updated_at"`
//		DeletedAt gorm.DeletedAt `json:"deleted_at"`
//	}
//
//	if req.Related { // 統計該作者所有插畫數量
//
//	}
//
//	if req.Page <= 0 {
//		req.Page = 1
//	}
//	if req.Size <= 0 {
//		req.Size = 10
//	}
//	offset := (req.Page - 1) * req.Size
//
//	db := this.Db.Model(&models.IllustrationAuthor{})
//
//	if req.Search != "" {
//		// 模糊查询 name
//		db = db.Where("name LIKE ?", "%"+req.Search+"%")
//	}
//
//	var authors []models.IllustrationAuthor
//	var total int64
//	if err := db.Count(&total).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to count authors"})
//		return
//	}
//
//	// 分页查询
//	if err := db.Order("created_at DESC").Offset(offset).Limit(req.Size).Find(&authors).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch authors"})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"list":  authors,
//		"total": total,
//		"page":  req.Page,
//		"size":  req.Size,
//	})
//}

func (this *AdminService) GetIllustrationAuthorList(ctx *gin.Context) {
	var req dto.GetIllustrationAuthorListRequestDto
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
	offset := (req.Page - 1) * req.Size

	db := this.Db.Model(&models.IllustrationAuthor{})

	if req.Search != "" {
		db = db.Where("name LIKE ?", "%"+req.Search+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to count authors"})
		return
	}

	if req.Related {
		// 使用自定义结构体
		type IllustrationAuthorWithRelated struct {
			Id        uuid.UUID      `json:"id"`
			Name      string         `json:"name"`
			Link      string         `json:"link"`
			Related   int64          `json:"related"`
			CreatedAt *time.Time     `json:"created_at"`
			UpdatedAt *time.Time     `json:"updated_at"`
			DeletedAt gorm.DeletedAt `json:"deleted_at"`
		}

		var authors []models.IllustrationAuthor
		if err := db.Order("created_at DESC").Offset(offset).Limit(req.Size).Find(&authors).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch authors"})
			return
		}

		result := make([]IllustrationAuthorWithRelated, 0, len(authors))
		for _, a := range authors {
			var count int64
			this.Db.Model(&models.Illustration{}).Where("author_id = ?", a.Id).Count(&count)
			result = append(result, IllustrationAuthorWithRelated{
				Id:        a.Id,
				Name:      a.Name,
				Link:      a.Link,
				Related:   count,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.UpdatedAt,
				DeletedAt: a.DeletedAt,
			})
		}

		// 即使没有数据，也返回空数组
		if result == nil {
			result = []IllustrationAuthorWithRelated{}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"list":  result,
			"total": total,
			"page":  req.Page,
			"size":  req.Size,
		})
	} else {
		// 直接返回原始模型
		var authors []models.IllustrationAuthor
		if err := db.Order("created_at DESC").Offset(offset).Limit(req.Size).Find(&authors).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch authors"})
			return
		}

		// 确保返回空数组而不是 nil
		if authors == nil {
			authors = []models.IllustrationAuthor{}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"list":  authors,
			"total": total,
			"page":  req.Page,
			"size":  req.Size,
		})
	}
}

func (this *AdminService) GetIllustrationAuthorById(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	var existingAuthor models.IllustrationAuthor
	if err := this.Db.First(&existingAuthor, "id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "auth not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, existingAuthor)
}

func (this *AdminService) RemoveIllustrationAuthorById(ctx *gin.Context) {
	uid, err := this.utils.GetAndParseParamUuid("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	this.utils.Logger.PrintInfo(uid)

	existing, err := this.isAuthorHasExistingIllustration(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if existing {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "this author has illustration existed, you cannot delete now",
		})
		return
	}

	if err := this.Db.Delete(&models.IllustrationAuthor{}, "id = ?", uid).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
