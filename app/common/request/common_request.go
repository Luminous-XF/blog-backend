package request

import "github.com/google/uuid"

// PageInfoRequest 分页信息
type PageInfoRequest struct {
	Page     int `json:"page" binding:"required,gte=1"`
	PageSize int `json:"pageSize" binding:"required,gte=1"`
}

type GetByUUIDRequest struct {
	UUID uuid.UUID `json:"uuid" binding:"required,uuid"`
}
