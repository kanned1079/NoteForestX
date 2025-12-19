package dto

// 文章状态枚举
type ArticleStatus string

const (
	Draft     ArticleStatus = "draft"
	Published ArticleStatus = "published"
	Hidden    ArticleStatus = "hidden"
)

// 新增或更新文章请求 DTO
type AddOrUpdateArticleRequestDto struct {
	ID       *string  `json:"id,omitempty"`   // uuid，可选，更新时填
	Slug     *string  `json:"slug,omitempty"` // 可选 slug
	Title    string   `json:"title"`
	Top      bool     `json:"top"`                 // 是否置顶
	Status   string   `json:"status"`              // "draft" | "published" | "hidden"
	Content  string   `json:"content"`             // Markdown
	Tags     []string `json:"tags,omitempty"`      // 标签名称或 code，前端传即可
	ImageUrl *string  `json:"image_url,omitempty"` // 可选封面
}

// 获取文章列表请求 DTO
type GetArticleListRequestDto struct {
	Page   int    `form:"page" json:"page"`
	Size   int    `form:"size" json:"size"`
	Search string `form:"search" json:"search,omitempty"`
	Tag    string `form:"tag" json:"tag,omitempty"`       // 按标签过滤，可选
	Status string `form:"status" json:"status,omitempty"` // 可选过滤文章状态
}

// 根据 ID 获取文章请求 DTO
type GetArticleByIdRequestDto struct {
	ID string `form:"id" json:"id"` // uuid(string)
}

// 根据 ID 删除文章请求 DTO
type RemoveArticleByIdRequestDto struct {
	ID string `form:"id" json:"id"` // uuid(string)
}
