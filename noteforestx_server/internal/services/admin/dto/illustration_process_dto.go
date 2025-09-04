package dto

type AddNewIllustrationRequestDto struct {
	Name          string   `form:"name" binding:"required"`
	Author        string   `form:"author"`
	PixivAuthorId string   `form:"pixiv_author_id"`
	PixivLink     string   `form:"pixiv_link"`
	TagsId        []string `form:"tags_id[]"` // form-data 多值 key=tags[]
}

type AddNewIllustrationResponseDto struct {
}

type UpdateIllustrationByIdRequestDto struct {
}
