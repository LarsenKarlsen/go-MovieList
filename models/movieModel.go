package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title     string
	Year      int
	ImageLink string
}
