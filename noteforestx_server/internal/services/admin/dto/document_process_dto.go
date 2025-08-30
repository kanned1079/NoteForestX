package dto

type AddOrUpdateNewDocumentRequestDto struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}

type GetDocumentListRequestDto struct {
	Page   int    `form:"page" json:"page"`
	Size   int    `form:"page" json:"size"`
	Search string `form:"search" json:"search"`
}

type GetDocumentByIdRequestDto struct {
	Id string `form:"id" json:"id"` // uuid(string)
}

type RemoveDocumentByIdRequestDto struct {
	Id string `form:"id" json:"id"` // uuid(string)
}
