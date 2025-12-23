package dto

import "noteforestx_server/internal/models"

// 文章状态枚举
type ArticleStatus string

const (
	Draft     ArticleStatus = "draft"
	Published ArticleStatus = "published"
	Hidden    ArticleStatus = "hidden"
)

// 新增或更新文章请求 DTO
type AddOrUpdateArticleRequestDto struct {
	//ID       *string  `json:"id,omitempty"`   // uuid，可选，更新时填
	Slug     *string             `json:"slug,omitempty"` // 可选 slug
	Title    string              `json:"title"`
	Top      bool                `json:"top"`                 // 是否置顶
	Status   string              `json:"status"`              // "draft" | "published" | "hidden"
	Content  string              `json:"content"`             // Markdown
	Tags     []models.ArticleTag `json:"tags,omitempty"`      // 标签名称或 code，前端传即可
	ImageUrl *string             `json:"image_url,omitempty"` // 可选封面
}

// 获取文章列表请求 DTO
type GetArticleListRequestDto struct {
	Page   int     `form:"page" json:"page"`                   // 总是需要
	Size   int     `form:"size" json:"size"`                   // 总是需要
	Search *string `form:"search" json:"search,omitempty"`     // 为空就全部 查询标题
	TagId  *string `form:"tag_id" json:"tag_id,omitempty"`     // 这是一个tag的uuid 找出使用了这个tag的文章 如果这个不为空 忽略Tag名搜索和标题搜索 但是状态部分还是要加入判断的
	Tag    *string `form:"tag" json:"tag,omitempty,omitempty"` // 按标签名过滤 可选 模糊查找用了相似标签名的文章 也要关心状态 忽略标题搜索
	Status *string `form:"status" json:"status,omitempty"`     // 可选过滤文章状态 留空即所有
}

// 根据 ID 获取文章请求 DTO
type GetArticleByIdRequestDto struct {
	ID string `form:"id" json:"id"` // uuid(string)
}

// 根据 ID 删除文章请求 DTO
type RemoveArticleByIdRequestDto struct {
	ID string `form:"id" json:"id"` // uuid(string)
}
