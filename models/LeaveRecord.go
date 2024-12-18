package models

import "time"

type LeaveRecord struct {
	ID          uint      `gorm:"primaryKey "`
	UserId      uint      `gorm:"not null"`
	LeaveTypeId uint      `gorm:"not null" binding:"required"`
	StartDate   time.Time `gorm:"not null"  binding:"required"`
	EndDate     time.Time `gorm:"not null"  binding:"required"`
	Status      string    `gorm:"default:Pending"`
	Reason      string    `gorm:"type:text"  binding:"required"`
	FilePath    string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreatedTime"`
}
