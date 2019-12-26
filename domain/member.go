package domain 

import (
	"github.com/jinzhu/gorm"
)

// Member struct to hold member data
type Member struct {
	gorm.Model
	Name string
	Email string
	Spouse string
	GroupID string `gorm:foreignkey`
}