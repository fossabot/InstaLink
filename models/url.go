package models

type URLEntry struct {
	gorm.Model
	actual  string
	shorten string
	refer   string
	count   int
}
