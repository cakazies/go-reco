package models

import "github.com/jinzhu/gorm"

type (
	Student struct {
		Name    string `json:"name"`
		Class   int    `json:"class"`
		Address string `json:"address"`
		Age     int    `json:"age"`
		gorm.Model
	}
)
