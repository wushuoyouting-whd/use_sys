package domain

import (
	"time"
)

// User 用户实体
type User struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(400)" json:"name"`
	Age       int        `gorm:"default:0" json:"age"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	BirthDate *time.Time `gorm:"type:date" json:"birthDate"`
	BeType    string     `gorm:"type:varchar(100)" json:"beType"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
