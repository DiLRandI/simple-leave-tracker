package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`

	LoginID uint

	Login Login
}
