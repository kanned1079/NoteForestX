package dto

import "noteforestx_server/internal/models"

type UserLoginRequestDto struct {
	Email    string `json:"email" `
	Password string `json:"password"`
}

type UserLoginResponseDto struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

//type UserRegisterRequestDto struct {
//}
//
//type UserRegisterResponse struct {
//}
