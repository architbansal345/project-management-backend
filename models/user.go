package models

type User struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	FirstName       string `json:"firstName" gorm:"size:100;not null" binding:"required"`
	LastName        string `json:"lastName" gorm:"size:100;not null" binding:"required"`
	Email           string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password        string `json:"password" gorm:"not null" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" gorm:"not null" binding:"required,min=6"`
	Role            string `json:"role" gorm:"default:employee"`
	CreatedAt       int64  `gorm:"autoCreateTime:nano"`
	UpdatedAt       int64  `gorm:"autoUpdateTime:nano"`
}
