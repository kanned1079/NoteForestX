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

// 支持自定義新文件名
func (this *AdminService) saveAndCompressFile(file *multipart.FileHeader, newFileName string) (fileName string, w, h int, err error) {
	// 保存的路径
	baseDir := config.ExistingAppConfig.Illustration.SaveDir
	paths := map[string]string{
		"original": filepath.Join(baseDir, "original", newFileName),
		"medium":   filepath.Join(baseDir, "medium", newFileName),
		"small":    filepath.Join(baseDir, "small", newFileName),
	}

	// 创建目录
	for _, p := range paths {
		if err = os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			err = fmt.Errorf("failed to create dir: %w", err)
			return
		}
	}

	// 保存原图
	if err = ctxSaveUploadedFile(file, paths["original"]); err != nil {
		return
	}

	// 打开原图进行解码
	srcFile, openErr := file.Open()
	if openErr != nil {
		_ = os.Remove(paths["original"])
		err = fmt.Errorf("failed to open uploaded file: %w", openErr)
		return
	}
	defer srcFile.Close()

	img, _, decodeErr := image.Decode(srcFile)
	if decodeErr != nil {
		_ = os.Remove(paths["original"])
		err = fmt.Errorf("failed to decode image: %w", decodeErr)
		return
	}

	// 保存原始宽高
	w = img.Bounds().Dx()
	h = img.Bounds().Dy()

	// 事务式保存压缩图
	savedFiles := []string{paths["original"]}
	defer func() {
		if err != nil {
			for _, f := range savedFiles {
				_ = os.Remove(f)
			}
		}
	}()

	// 保存 medium
	medium := imaging.Resize(img, config.ExistingAppConfig.Illustration.CompressedMediumPixel, 0, imaging.Lanczos)
	if saveErr := imaging.Save(medium, paths["medium"], imaging.JPEGQuality(85)); saveErr != nil {
		err = fmt.Errorf("failed to save medium: %w", saveErr)
		return
	}
	savedFiles = append(savedFiles, paths["medium"])

	// 保存 small
	small := imaging.Resize(img, config.ExistingAppConfig.Illustration.CompressedSmallPixel, 0, imaging.Lanczos)
	if saveErr := imaging.Save(small, paths["small"], imaging.JPEGQuality(80)); saveErr != nil {
		err = fmt.Errorf("failed to save small: %w", saveErr)
		return
	}
	savedFiles = append(savedFiles, paths["small"])

	// 返回文件名（這裏用 newFileName 而不是原始的 file.Filename）
	fileName = newFileName
	return
}

// ctx.SaveUploadedFile 的簡化版，不依賴 ctx
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

	if _, err = io.Copy(out, src); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	return nil
}

// AddNewIllustration POST:admin/illustration
//func (this *AdminService) AddNewIllustration(ctx *gin.Context) {
//	// 1. 绑定表单数据到 DTO
//	var dto dto.AddNewIllustrationRequestDto
//	if err := ctx.ShouldBind(&dto); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 2. 获取上传的文件
//	file, err := ctx.FormFile("file")
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
//		return
//	}
//
//	// 3. 生成 PixivId (用文件名去掉后缀)
//	ext := filepath.Ext(file.Filename)
//	pixivId := strings.TrimSuffix(file.Filename, ext)
//
//	if this.isIllustrationExisted(pixivId) {
//		ctx.JSON(http.StatusConflict, gin.H{
//			"message": "this illustration has existed. pixiv_id: " + pixivId,
//		})
//		return
//	}
//
//	// 4. 校验作者是否存在
//	if !this.isAuthorExisted(dto.AuthorId) {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "author does not exist"})
//		return
//	}
//
//	// 5. 校验 tags 是否存在
//	if !this.isTagsExisted(dto.TagsId) {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "one or more tags do not exist"})
//		return
//	}
//
//	illustrationUuid := uuid.New() // 图片文件的uuid
//
//	// 6. 保存文件并压缩
//	// 补全重新设置文件的文件名 在已知文件的前面加上前面的uuid_
//	// 比如说有这样一张插画 134997716_p0.jpg 那么这个文件的名需要变成 uuid_134997716_p0.jpg
//	// 其中文件名中的p0就是这一张插画属于这一个插画集中的第几张 p0即表示是第一张 那么其对应在IllustrationImage表中的Order也就是0
//	// 你还是需要修改这个接口方法获取文件的部分代码，来支持允许传多个插画文件
//	// 134997716 也就是这个插画集的id 它是一个string 并且长度不一定 我想让它作为index作为数据库中的唯一属性 来代替Id          uuid.UUID           `json:"id" gorm:"type:char(36);primaryKey"`
//	// 另外我看你在新的IllustrationImage表中有图片文件的宽高 那么还需要你写一个方法来获取这个宽高
//	_, err = this.saveAndCompressFile(file) // 保存一个文件到磁盘
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 7. 构建 Illustration 实体
//	// todo 你需要按照上面的描述来重新修改这个部分 这里的Id将会变为插画集的id 他是一个string
//	authorUUID, _ := uuid.Parse(dto.AuthorId)
//	illustration := models.Illustration{
//		Id:      illustrationUuid,
//		PixivId: pixivId,
//		//FilePath: paths["original"],
//		FilePath:    file.Filename, // 直接保存文件名即可 因为有图片压缩就不需要有前面的清晰度前缀
//		Name:        dto.Name,
//		AuthorId:    authorUUID,
//		Link:        dto.Link,
//		Description: dto.Description,
//		Limited:     dto.Limited,
//	}
//
//	// 8. 查询标签并绑定
//	if len(dto.TagsId) > 0 {
//		var tags []models.IllustrationTag
//		var uuids []uuid.UUID
//		for _, t := range dto.TagsId {
//			if uid, err := uuid.Parse(t); err == nil {
//				uuids = append(uuids, uid)
//			}
//		}
//		if len(uuids) > 0 {
//			if err := this.Db.Where("id IN ?", uuids).Find(&tags).Error; err != nil {
//				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query tags"})
//				return
//			}
//			illustration.Tags = tags
//		}
//	}
//
//	// 9. 保存数据库
//	if err := this.Db.Create(&illustration).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save illustration"})
//		return
//	}
//
//	// 10. 返回成功
//	ctx.JSON(http.StatusOK, gin.H{
//		"message": "illustration uploaded successfully",
//		"data":    illustration,
//	})
//}

// AddNewIllustration POST: admin/illustration
func (this *AdminService) AddNewIllustration(ctx *gin.Context) {
	// 1. 绑定表单数据到 DTO
	var dto dto.AddNewIllustrationRequestDto
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 获取上传的多个文件
	form, _ := ctx.MultipartForm()
	files := form.File["files"] // 这里前端字段名用 "files" 上传多个文件
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "files are required"})
		return
	}

	// 3. pixivId 或插画集 ID 用作 Illustration 主键
	// 从第一个文件名提取 pixivId
	ext := filepath.Ext(files[0].Filename)
	base := strings.TrimSuffix(files[0].Filename, ext)
	// 例如 134997716_p0.jpg -> 134997716
	pixivId := strings.Split(base, "_")[0]
	//pixivId := dto.PixivId
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

	authorUUID, _ := uuid.Parse(dto.AuthorId)

	// 6. 构建 Illustration 实体
	illustration := models.Illustration{
		//Id:          uuid.New(), // 这里数据库还是可以用 UUID 生成，也可以替换成 pixivId 字符串
		Id:          pixivId,
		PixivId:     pixivId,
		Name:        dto.Name,
		AuthorId:    authorUUID,
		Link:        dto.Link,
		Description: dto.Description,
		Limited:     dto.Limited,
		Source:      dto.Source, // 假设 DTO 里有来源字段
	}

	// 7. 查询标签并绑定
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

	// 8. 保存 Illustration
	if err := this.Db.Create(&illustration).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save illustration"})
		return
	}

	// 9. 保存每张图片到磁盘，并生成 IllustrationImage
	var images []models.IllustrationImage
	for idx, file := range files {
		// 生成新的文件名 uuid_filename
		//fileUUID := uuid.New().String()
		//ext := filepath.Ext(file.Filename)
		//base := strings.TrimSuffix(file.Filename, ext)
		//newFileName := fmt.Sprintf("%s_%s%s", fileUUID, base, ext)
		fileUUID := uuid.New()
		ext := filepath.Ext(file.Filename)
		base := strings.TrimSuffix(file.Filename, ext)
		newFileName := fmt.Sprintf("%s_%s%s", fileUUID.String(), base, ext)

		// 保存并压缩文件
		_, width, height, err := this.saveAndCompressFile(file, newFileName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//images = append(images, models.IllustrationImage{
		//	Id:             uuid.New(),
		//	IllustrationId: illustration.Id, // ✅ 用插画集ID，不是文件名
		//	FilePath:       filePath,
		//	Order:          idx,
		//	Width:          width,
		//	Height:         height,
		//	CreatedAt:      timePtr(time.Now()),
		//	UpdatedAt:      timePtr(time.Now()),
		//})
		images = append(images, models.IllustrationImage{
			Id:             fileUUID, // ✅ 这里直接用 fileUUID
			IllustrationId: illustration.Id,
			FilePath:       newFileName, // 或 filePath
			Order:          idx,
			Width:          width,
			Height:         height,
			CreatedAt:      timePtr(time.Now()),
			UpdatedAt:      timePtr(time.Now()),
		})
	}

	if len(images) > 0 {
		if err := this.Db.Create(&images).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save illustration images"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "illustration uploaded successfully",
		"data":    illustration,
		"images":  images,
	})
}

// 获取指针辅助函数
func timePtr(t time.Time) *time.Time {
	return &t
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
//func (this *AdminService) RemoveIllustrationById(ctx *gin.Context) {
//	id, err := this.utils.GetAndParseParamUuid("id", ctx)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
//		return
//	}
//
//	paraData := struct {
//		RemoveFile bool `form:"remove_file" json:"remove_file"` // 默认为 false，true 时删除所有文件
//	}{}
//	if err := ctx.ShouldBindQuery(&paraData); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid param"})
//		return
//	}
//
//	// 查询插画
//	var illustration models.Illustration
//	if err := this.Db.Preload("Tags").First(&illustration, "id = ?", id).Error; err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"message": "illustration not found"})
//		return
//	}
//
//	// 启动事务
//	tx := this.Db.Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//
//	// 先清理标签关系
//	if err := tx.Model(&illustration).Association("Tags").Clear(); err != nil {
//		tx.Rollback()
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to clear tags relation"})
//		return
//	}
//
//	if paraData.RemoveFile {
//		// 删除三种清晰度的文件
//		baseDir := config.ExistingAppConfig.Illustration.SaveDir
//		paths := []string{
//			filepath.Join(baseDir, "original", illustration.FilePath),
//			filepath.Join(baseDir, "medium", illustration.FilePath),
//			filepath.Join(baseDir, "small", illustration.FilePath),
//		}
//		for _, p := range paths {
//			_ = os.Remove(p) // 忽略错误（比如文件不存在）
//		}
//	} else {
//		// 移动原图到 trash 并删除 medium/small
//		if err := this.moveOriginalToTrash(illustration.FilePath); err != nil {
//			tx.Rollback()
//			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to move original file to trash"})
//			return
//		}
//		baseDir := config.ExistingAppConfig.Illustration.SaveDir
//		_ = os.Remove(filepath.Join(baseDir, "medium", illustration.FilePath))
//		_ = os.Remove(filepath.Join(baseDir, "small", illustration.FilePath))
//	}
//
//	// 删除插画记录
//	if err := tx.Delete(&illustration).Error; err != nil {
//		tx.Rollback()
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete illustration"})
//		return
//	}
//
//	if err := tx.Commit().Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to commit transaction"})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
//}

func (this *AdminService) RemoveIllustrationById(ctx *gin.Context) {
	id := ctx.Param("id") // Illustration.Id 是 string，不再是 uuid
	if id == "" {
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

	// 查询插画并预加载 tags 和 images
	var illustration models.Illustration
	if err := this.Db.Preload("Tags").Preload("Images").First(&illustration, "id = ?", id).Error; err != nil {
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

	// 清理标签关系
	if err := tx.Model(&illustration).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to clear tags relation"})
		return
	}

	baseDir := config.ExistingAppConfig.Illustration.SaveDir

	// 删除图片文件
	for _, img := range illustration.Images {
		if paraData.RemoveFile {
			// 删除三种清晰度文件
			paths := []string{
				filepath.Join(baseDir, "original", img.FilePath),
				filepath.Join(baseDir, "medium", img.FilePath),
				filepath.Join(baseDir, "small", img.FilePath),
			}
			for _, p := range paths {
				_ = os.Remove(p)
			}
		} else {
			// 移动原图到 trash，并删除 medium/small
			if err := this.moveOriginalToTrash(img.FilePath); err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to move original file to trash"})
				return
			}
			_ = os.Remove(filepath.Join(baseDir, "medium", img.FilePath))
			_ = os.Remove(filepath.Join(baseDir, "small", img.FilePath))
		}
	}

	// 删除插画（会级联删除 images，因为有外键 references:Id）
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
