package dto

type AddNewIllustrationTagRequestDto struct {
	Name string `json:"name"`
}

type UpdateIllustrationTagByIdRequestDto struct {
	Name string `json:"name"`
}

type GetIllustrationTagListRequestDto struct {
	Page   int    `form:"page" json:"page"`     // default 1 if not provided
	Size   int    `form:"size" json:"size"`     // default 10 if not provided
	Search string `form:"search" json:"search"` // if empty or not provided, select all
}
