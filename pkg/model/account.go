package model

import (
	"github.com/jinzhu/gorm"
)

// Account database model.
type Account struct {
	*gorm.Model
	Name      string
	BaseURL   string
	Publisher string
	APIToken  string
}
