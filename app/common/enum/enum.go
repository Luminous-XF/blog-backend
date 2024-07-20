package enum

// Type 帖子类型枚举
type Type string

const (
	REGULAR Type = "REGULAR"
	HOT     Type = "HOT"
	TOP     Type = "HOT"
)

// Status 帖子状态枚举
type Status string

const (
	NORMAL  Status = "NORMAL"
	REVIEW  Status = "REVIEW"
	BLOCKED Status = "BLOCKED"
)
