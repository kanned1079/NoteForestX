package admin

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (this *AdminService) isAuthorExisted(authorId string) bool {
	uid, err := uuid.Parse(authorId)
	if err != nil {
		// 无效 UUID 视为不存在
		return false
	}
	var count int64
	err = this.Db.Model(&models.IllustrationAuthor{}).
		Where("id = ?", uid).
		Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (this *AdminService) isTagsExisted(tags []string) bool {
	if len(tags) == 0 {
		return true
	}
	// Convert to UUID slice
	var uuids []uuid.UUID
	for _, t := range tags {
		uid, err := uuid.Parse(t)
		if err != nil {
			return false
		}
		uuids = append(uuids, uid)
	}

	var count int64
	err := this.Db.Model(&models.IllustrationTag{}).
		Where("id IN ?", uuids).
		Count(&count).Error
	if err != nil {
		return false
	}

	return count == int64(len(uuids))
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

// AddNewIllustration POST:admin/illustration
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

	// 4. 校验作者是否存在
	if !this.isAuthorExisted(dto.AuthorId) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "author does not exist"})
		return
	}

	// 5. 校验 tags 是否存在
	if !this.isTagsExisted(dto.TagsId) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "one or more tags do not exist"})
		return
	}

	// 6. 保存文件并压缩
	_, err = this.saveAndCompressFile(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 7. 构建 Illustration 实体
	authorUUID, _ := uuid.Parse(dto.AuthorId)
	illustration := models.Illustration{
		Id:      uuid.New(),
		PixivId: pixivId,
		//FilePath: paths["original"],
		FilePath:    file.Filename, // 直接保存文件名即可 因为有图片压缩就不需要有前面的清晰度前缀
		Name:        dto.Name,
		AuthorId:    authorUUID,
		Link:        dto.Link,
		Description: dto.Description,
		Limited:     dto.Limited,
	}

	// 8. 查询标签并绑定
	if len(dto.TagsId) > 0 {
		var tags []models.IllustrationTag
		var uuids []uuid.UUID
		for _, t := range dto.TagsId {
			if uid, err := uuid.Parse(t); err == nil {
				uuids = append(uuids, uid)
			}
		}
		if len(uuids) > 0 {
			if err := this.Db.Where("id IN ?", uuids).Find(&tags).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query tags"})
				return
			}
			illustration.Tags = tags
		}
	}

	// 9. 保存数据库
	if err := this.Db.Create(&illustration).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save illustration"})
		return
	}

	// 10. 返回成功
	ctx.JSON(http.StatusOK, gin.H{
		"message": "illustration uploaded successfully",
		"data":    illustration,
	})
}

// UpdateIllustrationById PUT:admin/illustration/:id
func (this *AdminService) UpdateIllustrationById(ctx *gin.Context) {
	// 1. 获取插画ID
	uid, err := this.utils.GetAndParseParamUuid("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	// 2. 绑定请求体
	var dto dto.UpdateIllustrationByIdRequestDto
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. 开启事务
	tx := this.Db.Begin()
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to start transaction"})
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "panic occurred"})
		}
	}()

	// 4. 查询现有插画
	var illustration models.Illustration
	if err := tx.Preload("Tags").First(&illustration, "id = ?", uid).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "illustration not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	// 5. 校验作者是否存在
	if !this.isAuthorExisted(dto.AuthorId) {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "author does not exist"})
		return
	}

	// 6. 校验标签是否存在
	if !this.isTagsExisted(dto.TagsId) {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "one or more tags do not exist"})
		return
	}

	// 7. 更新基本字段
	authorUUID, _ := uuid.Parse(dto.AuthorId)
	updateData := map[string]interface{}{
		"name":        dto.Name,
		"author_id":   authorUUID,
		"link":        dto.Link,
		"limited":     dto.Limited,
		"description": dto.Description,
	}

	if err := tx.Model(&illustration).Updates(updateData).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update illustration"})
		return
	}

	// 8. 更新标签关系
	if dto.TagsId != nil {
		// 查询标签
		var tags []models.IllustrationTag
		var uuids []uuid.UUID
		for _, t := range dto.TagsId {
			if uid, err := uuid.Parse(t); err == nil {
				uuids = append(uuids, uid)
			}
		}
		if len(uuids) > 0 {
			if err := tx.Where("id IN ?", uuids).Find(&tags).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to query tags"})
				return
			}
		}

		// 替换标签关系
		if err := tx.Model(&illustration).Association("Tags").Replace(tags); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update tags"})
			return
		}
	}

	// 9. 提交事务
	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to commit transaction"})
		return
	}

	// 10. 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "illustration updated successfully",
		"illustration": illustration,
	})
}

// 删除三种清晰度的文件
func (this *AdminService) deleteIllustrationFiles(fileName string) error {
	baseDir := config.ExistingAppConfig.Illustration.SaveDir
	paths := []string{
		filepath.Join(baseDir, "original", fileName),
		filepath.Join(baseDir, "medium", fileName),
		filepath.Join(baseDir, "small", fileName),
	}

	for _, p := range paths {
		if err := os.Remove(p); err != nil {
			// 如果文件不存在，不算错误
			if !os.IsNotExist(err) {
				return fmt.Errorf("failed to delete file %s: %w", p, err)
			}
		}
	}
	return nil
}

// 移动原图到 trash，并删除 medium/small
func (this *AdminService) moveOriginalToTrash(fileName string) error {
	baseDir := config.ExistingAppConfig.Illustration.SaveDir
	srcPath := filepath.Join(baseDir, "original", fileName)
	trashDir := filepath.Join(baseDir, "trash")
	dstPath := filepath.Join(trashDir, fileName)

	// 确保 trash 目录存在
	if err := os.MkdirAll(trashDir, 0755); err != nil {
		return fmt.Errorf("failed to create trash dir: %w", err)
	}

	// 移动原图
	if err := os.Rename(srcPath, dstPath); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to move original file to trash: %w", err)
		}
	}

	// 删除 medium 和 small
	for _, sub := range []string{"medium", "small"} {
		p := filepath.Join(baseDir, sub, fileName)
		if err := os.Remove(p); err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("failed to delete %s file: %w", sub, err)
			}
		}
	}

	return nil
}

// RemoveIllustrationById DELETE:admin/illustration/:id?remove_file=false
func (this *AdminService) RemoveIllustrationById(ctx *gin.Context) {
	id, err := this.utils.GetAndParseParamUuid("id", ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	paraData := struct {
		RemoveFile bool `form:"remove_file" json:"remove_file"` // 默认为 false，true 时删除所有文件
	}{}
	if err := ctx.ShouldBindQuery(&paraData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid param"})
		return
	}

	// 查询插画
	var illustration models.Illustration
	if err := this.Db.Preload("Tags").First(&illustration, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "illustration not found"})
		return
	}

	// 启动事务
	tx := this.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 先清理标签关系
	if err := tx.Model(&illustration).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to clear tags relation"})
		return
	}

	if paraData.RemoveFile {
		// 删除三种清晰度的文件
		baseDir := config.ExistingAppConfig.Illustration.SaveDir
		paths := []string{
			filepath.Join(baseDir, "original", illustration.FilePath),
			filepath.Join(baseDir, "medium", illustration.FilePath),
			filepath.Join(baseDir, "small", illustration.FilePath),
		}
		for _, p := range paths {
			_ = os.Remove(p) // 忽略错误（比如文件不存在）
		}
	} else {
		// 移动原图到 trash 并删除 medium/small
		if err := this.moveOriginalToTrash(illustration.FilePath); err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to move original file to trash"})
			return
		}
		baseDir := config.ExistingAppConfig.Illustration.SaveDir
		_ = os.Remove(filepath.Join(baseDir, "medium", illustration.FilePath))
		_ = os.Remove(filepath.Join(baseDir, "small", illustration.FilePath))
	}

	// 删除插画记录
	if err := tx.Delete(&illustration).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete illustration"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to commit transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

/*
	type GetIllustrationListRequestDto struct {
		Page          int    `form:"page" json:"page"`
		Size          int    `form:"size" json:"size"`
		SearchAs      string `form:"search_as" json:"search_as"`           // "author"作者 | "tag"标签名 | "name"插画名 默认使用tag
		SearchContent string `form:"search_content" json:"search_content"` // 搜寻的内容 留空全部
		Sort          string `form:"sort" json:"sort"`                     // "ASC" | "DESC" 默认按照created_at降序
	}
*/

//func (this *AdminService) GetIllustrationList(ctx *gin.Context) {
//
//}

//func (this *AdminService) GetIllustrationById(ctx *gin.Context) {
//
//}
