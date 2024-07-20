package response

import (
	"blog-backend/app/common/enum"
	"github.com/google/uuid"
	"time"
)

type PostResponse struct {
	UUID           uuid.UUID   `json:"uuid"`
	Title          string      `json:"title"`
	Excerpt        string      `json:"excerpt"`
	Content        string      `json:"content"`
	Type           enum.Type   `json:"type"`
	CommentCount   uint64      `json:"comment_count"`
	Score          float64     `json:"score"`
	Status         enum.Status `json:"status"`
	Likes          uint64      `json:"likes"`
	PageViews      uint64      `json:"page_views"`
	HeaderImageUrl string      `json:"header_image_url"`
	CreateTime     time.Time   `json:"create_time"`
	UpdateTime     time.Time   `json:"update_time"`
}
