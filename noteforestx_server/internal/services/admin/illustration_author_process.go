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

func (this *AdminService) getAuthorInfoFromCtx(ctx *gin.Context) (err error, name, link string) {
	var req dto.AddNewIllustrationAuthorRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err, "", ""
	}
	name = strings.TrimSpace(req.Name)
	if name == "" {
		return errors.New("name cannot be empty"), "", ""
	}
	link = strings.TrimSpace(req.Link)
	if link == "" {
		return errors.New("link cannot be empty"), "", ""
	}
	return
}

func (this *AdminService) AddNewIllustrationAuthor(ctx *gin.Context) {
	err, name, link := this.getAuthorInfoFromCtx(ctx)
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
	err, name, link := this.getAuthorInfoFromCtx(ctx)
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

func (this *AdminService) GetIllustrationAuthorList(ctx *gin.Context) {
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
	offset := (req.Page - 1) * req.Size

	db := this.Db.Model(&models.IllustrationAuthor{})

	if req.Search != "" {
		// 模糊查询 name
		db = db.Where("name LIKE ?", "%"+req.Search+"%")
	}

	var authors []models.IllustrationAuthor
	var total int64
	if err := db.Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to count authors"})
		return
	}

	// 分页查询
	if err := db.Order("created_at DESC").Offset(offset).Limit(req.Size).Find(&authors).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch authors"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  authors,
		"total": total,
		"page":  req.Page,
		"size":  req.Size,
	})
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
