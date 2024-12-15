package models

import "gorm.io/gorm"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Completed   bool
}

func MigrateTask(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
