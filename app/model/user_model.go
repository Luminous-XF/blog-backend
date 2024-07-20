package user

import (
    "blog-backend/app/model/common"
    "github.com/google/uuid"
)

type User struct {
    common.Model
    UUID           uuid.UUID `gorm:"column:uuid"`
    Username       string    `gorm:"column:username;unique"`
    Nickname       string    `gorm:"column:nickname"`
    Password       string    `gorm:"column:password"`
    Salt           string    `gorm:"column:salt"`
    Email          string    `gorm:"column:email;unique"`
    AvatarImageURL string    `gorm:"column:avatar_image_url"`
}

func (User) TableName() string {
    return "user"
}
