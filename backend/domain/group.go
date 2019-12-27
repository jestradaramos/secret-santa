package domain

import (
//"github.com/jinzhu/gorm"
)

// Group struct to hold group data
type Group struct {
	ID           string `gorm:"PRIMARY_KEY"`
	MoneyLimit   float64
	Deadline     string
	ExchangeDate string
}
