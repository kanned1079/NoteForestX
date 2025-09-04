package illustration

import (
	"bytes"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"io"
	"mime/multipart"
	"noteforestx_server/internal/config"
	"path/filepath"
	"strings"
	"time"
)

func (this *IllustrationService) ReadFromRedis(size, fileName string) ([]byte, error) {
	if !config.ExistingAppConfig.RedisConfig.Enabled {
		return nil, nil
	}

	key := "illustration:" + size + ":" + strings.TrimSuffix(fileName, filepath.Ext(fileName))

	redisCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := this.Rdb.Get(redisCtx, key).Bytes()
	if err != nil {
		// redis.Nil 表示 key 不存在，返回 nil
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func (this *IllustrationService) SaveToRedis(size string, file *multipart.FileHeader) error {
	if !config.ExistingAppConfig.RedisConfig.Enabled {
		return nil
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		return err
	}

	key := "illustration:" + size + ":" + strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	expire := 24 * time.Hour

	redisCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := this.Rdb.Set(redisCtx, key, buf.Bytes(), expire).Err(); err != nil {
		return err
	}

	return nil
}
