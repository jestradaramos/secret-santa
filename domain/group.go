package domain

import (
	"github.com/jinzhu/gorm"
)

// Group struct to hold group data
type Group struct {
	gorm.Model
	Name string
	MoneyLimit float64
	Deadline string
	ExchangeDate string
}