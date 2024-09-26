package model

import "gorm.io/gorm"

type Login struct {
	gorm.Model

	Username string `gorm:"uniqueIndex"`
	Password string
}
