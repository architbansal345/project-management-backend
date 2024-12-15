package models

type LeaveType struct {
	ID      uint   `gorm:"primaryKey"`
	Type    string `gorm:"size:50;not null"`
	MaxDays int    `gorm:"not null"`
}
