package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string
	Age  int
	// Created time.Time
	// Company   Company
	// CompanyID uint
}
