package model

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Fname     string
	Lname     string
	Tel       string
	Diagnosis string
}
