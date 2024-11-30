package user

import "gorm.io/gorm"
type users struct {
	gorm.Model
	Password string `gorm:"not null"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}
func (u *users) TableName() string {
	return "users"
}