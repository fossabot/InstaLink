package models

import "github.com/jinzhu/gorm"

type URLEntry struct {
	gorm.Model

	actual  string
	shorten string
	refer   string
	count   int
}
