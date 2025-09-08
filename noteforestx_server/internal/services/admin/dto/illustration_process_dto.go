package dto

type AddNewIllustrationRequestDto struct {
	Name        string   `form:"name" binding:"required" json:"name"`           // 插画名称
	AuthorId    string   `form:"author_id" binding:"required" json:"author_id"` // 作者对应的uuid
	TagsId      []string `form:"tags_id[]" json:"tags_id"`                      // 对应的Tag的id 可以选择多个
	Link        string   `form:"link" binding:"required" json:"link"`           // 原pixiv地址
	Description string   `form:"description" json:"description"`
	Limited     bool     `form:"limited" json:"limited"` // 是否限制型
}

type AddNewIllustrationResponseDto struct {
}

type UpdateIllustrationByIdRequestDto struct {
	Name        string   `form:"name" binding:"required" json:"name"`           // 插画名称
	AuthorId    string   `form:"author_id" binding:"required" json:"author_id"` // 作者对应的uuid
	TagsId      []string `form:"tags_id[]" json:"tags_id"`                      // 对应的Tag的id 可以选择多个
	Link        string   `form:"link" binding:"required" json:"link"`           // 原pixiv地址
	Description string   `form:"description" json:"description"`
	Limited     bool     `form:"limited" binding:"required" json:"limited"`
}

//type GetIllustrationListRequestDto struct {
//	Page          int    `form:"page" json:"page"`
//	Size          int    `form:"size" json:"size"`
//	SearchAs      string `form:"search_as" json:"search_as"`           // "author"作者 | "tag"标签名 | "name"插画名 默认使用tag
//	SearchContent string `form:"search_content" json:"search_content"` // 搜寻的内容 留空全部
//	Sort          string `form:"sort" json:"sort"`                     // "ASC" | "DESC" 默认按照created_at降序
//}
