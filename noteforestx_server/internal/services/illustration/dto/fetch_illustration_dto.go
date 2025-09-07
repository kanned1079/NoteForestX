package dto

import "noteforestx_server/internal/models"

type GetIllustrationListRequestDto struct {
	Page          int    `form:"page" json:"page"`
	Size          int    `form:"size" json:"size"`
	SearchAs      string `form:"search_as" json:"search_as"`           // "author"作者 | "tag"标签名 | "name"插画名 默认使用tag
	SearchContent string `form:"search_content" json:"search_content"` // 搜寻的内容 留空全部
	Sort          string `form:"sort" json:"sort"`                     // "ASC" | "DESC" 默认按照created_at降序
}

type GetIllustrationListResponseDto struct {
	List    []models.Illustration `json:"list"`
	Total   int                   `json:"total"`
	Message string                `json:"message"`
}
