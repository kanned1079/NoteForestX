package dto

// PatchArticleRequestDto
type PatchArticleRequestDto struct {
	Top    *bool   `json:"top,omitempty"`
	Status *string `json:"status,omitempty"`
}
