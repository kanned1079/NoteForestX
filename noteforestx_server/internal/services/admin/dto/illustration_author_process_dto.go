package dto

type AddNewIllustrationAuthorRequestDto struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}

type UpdateIllustrationAuthorByIdRequestDto struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}

type GetIllustrationAuthorListRequestDto struct {
	Page   int    `form:"page" json:"page"`     // default 1 if not provided
	Size   int    `form:"size" json:"size"`     // default 10 if not provided
	Search string `form:"search" json:"search"` // if empty or not provided, select all
}
