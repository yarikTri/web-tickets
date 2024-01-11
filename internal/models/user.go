package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username    string `gorm:"unique"`
	FullName    string
	Password    string
	IsModerator bool
}

func (u *User) ToTransfer() UserTransfer {
	return UserTransfer{
		ID:          u.ID,
		Username:    u.Username,
		FullName:    u.FullName,
		IsModerator: u.IsModerator,
	}
}

type UserTransfer struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	IsModerator bool   `json:"is_moderator"`
}
