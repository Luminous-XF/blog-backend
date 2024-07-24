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

type SendVerifyCodeWithEmailResponse struct {
    RequestID string `json:"requestID"`
}
