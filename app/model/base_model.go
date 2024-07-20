package common

import (
    "gorm.io/plugin/soft_delete"
    "time"
)

// Model 通用字段
type Model struct {
    ID            uint64                `json:"id" gorm:"column:id;primary_key;comment:记录ID"`               // 记录 ID
    RowVersion    uint64                `json:"rowVersion" gorm:"column:row_version;comment:记录版本号"`      // 记录版本号
    RowCreateTime time.Time             `json:"rowCreateTime" gorm:"column:row_create_time;autoCreateTime"`   // 记录创建时间
    RowUpdateTime time.Time             `json:"rowUpdateTime" gorm:"column:row_update_time;autoUpdateTime"`   // 记录修改时间
    RowIsDeleted  soft_delete.DeletedAt `json:"rowIsDeleted" gorm:"column:row_is_deleted" soft_delete:"flag"` // 记录软删除标记 0:未删除 1:已删除
}
