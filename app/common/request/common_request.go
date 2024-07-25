package request

import "github.com/google/uuid"

// PageInfoRequest 分页信息
type PageInfoRequest struct {
    Page     int `form:"page" json:"page" binding:"required,gte=1"`
    PageSize int `form:"pageSize" json:"pageSize" binding:"required,gte=1"`
}

type GetByUUIDRequest struct {
    UUID uuid.UUID `uri:"uuid" json:"uuid" binding:"required,uuid"`
}
