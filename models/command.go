package models

import "gorm.io/gorm"

type Command struct {
	gorm.Model
	Command string `json:"command"`
	UserID  uint   `json:"user_id"`
	User    User
}
