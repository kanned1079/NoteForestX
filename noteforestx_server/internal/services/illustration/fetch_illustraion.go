package illustration

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"noteforestx_server/internal/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (this *IllustrationService) FetchIllustration(ctx *gin.Context) {
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
