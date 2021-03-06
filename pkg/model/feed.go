package model

import (
	"github.com/jinzhu/gorm"
)

// Feed database model.
type Feed struct {
	*gorm.Model
	Name     string
	User     string
	URL      string
	Provider string
	HashTags []HashTag
}
