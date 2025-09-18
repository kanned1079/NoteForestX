package admin

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"image"
	"io"
	"mime/multipart"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services/admin/dto"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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
