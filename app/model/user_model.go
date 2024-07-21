package model

import (
    "blog-backend/pkg/helper"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    Model
    UUID           uuid.UUID `gorm:"column:uuid"`
    Username       string    `gorm:"column:username;unique"`
    Nickname       string    `gorm:"column:nickname"`
    Password       string    `gorm:"column:password"`
    Salt           string    `gorm:"column:salt"`
    Email          string    `gorm:"column:email;unique"`
    AvatarImageURL string    `gorm:"column:avatar_image_url"`
}

func (user *User) TableName() string {
    return "user"
}

func (user *User) BeforeSave(*gorm.DB) (err error) {
    // 对 Password 拼接盐值然后进行 MD5加密
    if len(user.Password) != 32 {
        user.Salt = helper.MakeStr(16, helper.DigitAlphaPunct)
        user.Password = helper.MD5(user.Password + user.Salt)
    }
    return
}
