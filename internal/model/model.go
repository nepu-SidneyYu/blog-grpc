package model

type Model struct {
	ID        int   `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt int64 `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
}
