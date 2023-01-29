package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title     string
	Year      string
	ImageLink string
}
