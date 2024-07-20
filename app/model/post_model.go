package model

import (
	"blog-backend/app/common/enum"
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Model
	UUID           uuid.UUID   `gorm:"column:uuid"`
	UserId         uint64      `gorm:"author_id"`
	Title          string      `gorm:"title"`
	Excerpt        string      `gorm:"column:excerpt"`
	Content        string      `gorm:"content"`
	Type           enum.Type   `gorm:"type"`
	CommentCount   uint64      `gorm:"comment_count"`
	Score          float64     `gorm:"score"`
	Status         enum.Status `gorm:"status"`
	Likes          uint64      `gorm:"likes"`
	PageViews      uint64      `gorm:"page_views"`
	HeaderImageUrl string      `gorm:"header_image_url"`
	CreateTime     time.Time   `gorm:"create_time"`
	UpdateTime     time.Time   `gorm:"update_time"`
}

func (Post) TableName() string {
	return "post"
}
