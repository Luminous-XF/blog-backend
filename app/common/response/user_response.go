package response

import "github.com/google/uuid"

// UserResponse 返回的user, 去除敏感字段
type UserResponse struct {
	UUID           uuid.UUID `json:"uuid"`
	Username       string    `json:"username"`
	Nickname       string    `json:"nickname"`
	Email          string    `json:"email"`
	AvatarImageURL string    `json:"avatar_image_url"`
}

// LoginResponse 登录返回,user token 和过期时间
type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type SendVerifyCodeWithEmailResponse struct {
	RequestID string `json:"request_id"`
}
