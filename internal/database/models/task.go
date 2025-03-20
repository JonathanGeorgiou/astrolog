package models

// import "gorm.io/gorm"

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}
