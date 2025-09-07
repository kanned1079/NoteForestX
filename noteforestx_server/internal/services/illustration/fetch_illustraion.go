package illustration

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/illustration/dto"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
type GetIllustrationListRequestDto struct {
	Page          int    `form:"page" json:"page"`
	Size          int    `form:"size" json:"size"`
	SearchAs      string `form:"search_as" json:"search_as"`           // "author"作者 | "tag"标签名 | "name"插画名 默认使用tag
	SearchContent string `form:"search_content" json:"search_content"` // 搜寻的内容 留空查询所有即可也不需要根据搜索类型选择
	Sort          string `form:"sort" json:"sort"`                     // "ASC" | "DESC" 默认按照created_at降序
}
*/

// GetIllustrationList GET:/api/v1/illustration?page=&size=&search_as=&search_content=&sort=
func (this *IllustrationService) GetIllustrationList(ctx *gin.Context) {
	var paraReq dto.GetIllustrationListRequestDto
	if err := ctx.ShouldBindQuery(&paraReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}

	// 默认分页
	if paraReq.Page <= 0 {
		paraReq.Page = 1
	}
	if paraReq.Size <= 0 {
		paraReq.Size = 10
	}
	offset := (paraReq.Page - 1) * paraReq.Size

	// 默认排序
	sortOrder := "i.created_at DESC"
	if paraReq.Sort == "ASC" {
		sortOrder = "i.created_at ASC"
	}

	// 给主表加别名 "i"
	db := this.Db.Model(&models.Illustration{}).Table(models.Illustration{}.TableName() + " AS i").
		Preload("Author").Preload("Tags")

	// 搜索处理
	if paraReq.SearchContent != "" {
		search := "%" + paraReq.SearchContent + "%"
		switch paraReq.SearchAs {
		case "author":
			authorTable := models.IllustrationAuthor{}.TableName()
			db = db.Joins("LEFT JOIN "+authorTable+" AS a ON a.id = i.author_id").
				Where("a.name LIKE ?", search)
		case "tag":
			tagMappingTable := models.IllustrationTagMapping{}.TableName()
			tagTable := models.IllustrationTag{}.TableName()
			db = db.Joins("LEFT JOIN "+tagMappingTable+" AS itm ON i.id = itm.illustration_id").
				Joins("LEFT JOIN "+tagTable+" AS t ON t.id = itm.tag_id").
				Where("t.name LIKE ?", search)
		case "name":
			db = db.Where("i.name LIKE ?", search)
		default:
			tagMappingTable := models.IllustrationTagMapping{}.TableName()
			tagTable := models.IllustrationTag{}.TableName()
			db = db.Joins("LEFT JOIN "+tagMappingTable+" AS itm ON i.id = itm.illustration_id").
				Joins("LEFT JOIN "+tagTable+" AS t ON t.id = itm.tag_id").
				Where("t.name LIKE ?", search)
		}
	}

	// 查询总数
	var total int64
	if err := db.Group("i.id").Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "查询失败: " + err.Error()})
		return
	}

	// 查询分页数据
	var illustrations []models.Illustration
	if err := db.Group("i.id").Order(sortOrder).Offset(offset).Limit(paraReq.Size).Find(&illustrations).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "查询失败: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  illustrations,
		"total": total,
		"page":  paraReq.Page,
		"size":  paraReq.Size,
	})
}

func (this *IllustrationService) FetchIllustrationByFilename(ctx *gin.Context) {
	// 1. 从 URL 获取文件名
	illustrationFileName := ctx.Param("file_name")
	if illustrationFileName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file_name is required"})
		return
	}

	// 2. 获取 size 查询参数，默认 original
	size := ctx.DefaultQuery("size", "original")
	if size != "small" && size != "medium" && size != "original" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "size must be small, medium, or original"})
		return
	}

	fileName := filepath.Base(illustrationFileName)

	// 3. 优先从 Redis 读取
	data, err := this.ReadFromRedis(size, fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
		return
	}
	if data != nil {
		// Redis 命中，直接返回
		ctx.Data(http.StatusOK, "image/jpeg", data)
		return
	}

	// 4. Redis 没有 -> 磁盘读取
	baseDir := config.ExistingAppConfig.Illustration.SaveDir
	var filePath string
	switch size {
	case "original":
		filePath = filepath.Join(baseDir, "original", fileName)
	case "medium":
		filePath = filepath.Join(baseDir, "medium", fileName)
	case "small":
		filePath = filepath.Join(baseDir, "small", fileName)
	}

	data, err = os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	// 5. 异步保存到 Redis（不阻塞响应）
	go func(data []byte, size, fileName string) {
		if !config.ExistingAppConfig.RedisConfig.Enabled {
			return
		}
		key := "illustration:" + size + ":" + strings.TrimSuffix(fileName, filepath.Ext(fileName))
		redisCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = this.Rdb.Set(redisCtx, key, data, 24*time.Hour).Err()
	}(data, size, fileName)

	// 6. 返回文件数据
	ctx.Data(http.StatusOK, "image/jpeg", data)
}
