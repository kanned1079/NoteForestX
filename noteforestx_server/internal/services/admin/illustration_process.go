package admin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// doc process crud

func (this *AdminService) saveToRedis(size string, file *multipart.FileHeader) error {
	key := "illustration:" + size + ":"

	if !config.ExistingAppConfig.RedisConfig.Enabled {
		return nil
	}

	// 1. 打开文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 2. 读取文件内容到内存
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		return err
	}

	// 3. 构造 Redis key，这里用 PixivId 或文件名
	key = key + strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))

	expire := 24 * time.Hour

	redisCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := this.Rdb.Set(redisCtx, key, buf.Bytes(), expire).Err(); err != nil {
		return err
	}

	return nil
}

func (this *AdminService) checkIllustrationMetadata(dto dto.AddNewIllustrationRequestDto) {

}

func (this *AdminService) isIllustrationExisted(illustrationId string) bool {
	var count int64
	err := this.Db.Model(&models.Illustration{}).Where("pixiv_id = ?", illustrationId).Count(&count).Error
	if err != nil {
		// 出错时也返回 true，防止重复插入
		return true
	}
	return count > 0
}

func (this *AdminService) saveAndCompressFile(file *multipart.FileHeader) (map[string]string, error) {
	// 保存的路径
	baseDir := config.ExistingAppConfig.Illustration.SaveDir
	paths := map[string]string{
		"original": filepath.Join(baseDir, "original", file.Filename),
		"medium":   filepath.Join(baseDir, "medium", file.Filename),
		"small":    filepath.Join(baseDir, "small", file.Filename),
	}

	// 创建目录
	for _, p := range paths {
		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			return nil, fmt.Errorf("failed to create dir: %w", err)
		}
	}

	// 保存原图（直接保存，不解码）
	if err := ctxSaveUploadedFile(file, paths["original"]); err != nil {
		return nil, err
	}

	// 打开原图进行压缩
	srcFile, err := file.Open()
	if err != nil {
		os.Remove(paths["original"])
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		os.Remove(paths["original"])
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	// 事务式保存压缩图
	savedFiles := []string{paths["original"]}
	defer func() {
		// 如果最终出错，删除已保存文件
		if err != nil {
			for _, f := range savedFiles {
				_ = os.Remove(f)
			}
		}
	}()

	// save medium
	medium := imaging.Resize(img, config.ExistingAppConfig.Illustration.CompressedMediumPixel, 0, imaging.Lanczos)
	if err := imaging.Save(medium, paths["medium"], imaging.JPEGQuality(85)); err != nil {
		return nil, fmt.Errorf("failed to save medium: %w", err)
	}
	savedFiles = append(savedFiles, paths["medium"])

	// save small
	small := imaging.Resize(img, config.ExistingAppConfig.Illustration.CompressedSmallPixel, 0, imaging.Lanczos)
	if err := imaging.Save(small, paths["small"], imaging.JPEGQuality(80)); err != nil {
		return nil, fmt.Errorf("failed to save small: %w", err)
	}
	savedFiles = append(savedFiles, paths["small"])

	paths["fileName"] = file.Filename

	return paths, nil
}

// 兼容 ctx.SaveUploadedFile 的逻辑，不依赖 ctx
func ctxSaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	return nil
}

func (this *AdminService) AddNewIllustration(ctx *gin.Context) {
	// 1. 绑定表单数据到 DTO
	var dto dto.AddNewIllustrationRequestDto
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 获取上传的文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// 3. 生成 PixivId (用文件名去掉后缀)
	ext := filepath.Ext(file.Filename)
	pixivId := strings.TrimSuffix(file.Filename, ext)

	if this.isIllustrationExisted(pixivId) {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "this illustration has existed. pixiv_id: " + pixivId,
		})
		return
	}

	// 4. 保存文件
	_, err = this.saveAndCompressFile(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 5. 转换 tags 为 JSON
	tagsJSON, err := json.Marshal(dto.TagsId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to encode tags"})
		return
	}

	// 6. 插入数据库
	illustration := models.Illustration{
		Id:            uuid.New(),
		PixivId:       pixivId,
		TagsId:        datatypes.JSON(tagsJSON),
		FilePath:      file.Filename,
		Name:          dto.Name,
		Author:        dto.Author,
		PixivAuthorId: dto.PixivAuthorId,
		PixivLink:     dto.PixivLink,
	}

	if err := this.Db.Create(&illustration).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save illustration"})
		return
	}

	// 7. 返回成功
	ctx.JSON(http.StatusOK, gin.H{
		"message": "illustration uploaded successfully",
		"data":    illustration,
	})
}

func (this *AdminService) UpdateIllustrationById(ctx *gin.Context) {

}

func (this *AdminService) GetIllustrationList(ctx *gin.Context) {

}

func (this *AdminService) GetIllustrationById(ctx *gin.Context) {

}

func (this *AdminService) RemoveIllustrationById(ctx *gin.Context) {

}
