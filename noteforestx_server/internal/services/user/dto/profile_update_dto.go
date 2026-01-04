package dto

type UpdateUsernameRequestDto struct {
	Username string `json:"username"`
}

type UpdateUserPasswordRequestDto struct {
	PreviousPassword string `json:"previous_password"`
	NewPassword      string `json:"new_password"`
}
