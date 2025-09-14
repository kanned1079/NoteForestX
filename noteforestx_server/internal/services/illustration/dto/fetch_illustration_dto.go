package dto

import "noteforestx_server/internal/models"

type GetIllustrationListRequestDto struct {
	Page          int      `form:"page" json:"page"`
	Size          int      `form:"size" json:"size"`
	SearchAs      string   `form:"search_as" json:"search_as"`           // "author"作者 | "tag"标签名 | "name"插画名 默认使用tag 使用tag允许多个tag
	SearchContent []string `form:"search_content" json:"search_content"` // 如果是tag则是uuid数组 允许多个 如果是按照作者或插画名 则只允许这个数组中只有一个项目
	Sort          string   `form:"sort" json:"sort"`                     // "ASC" | "DESC" 默认按照created_at降序
	ShowLimited   bool     `form:"show_limited" json:"show_limited"`     // 如果是true 则显示带有limit属性的插画
}

type GetIllustrationListResponseDto struct {
	List        []models.Illustration `json:"list"`
	Total       int                   `json:"total"`
	ShowLimited bool                  `json:"show_limited"`
	Message     string                `json:"message"`
}
