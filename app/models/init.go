package models

import "github.com/jinzhu/gorm"

type (
	Student struct {
		Name    string `json:"name"`
		Class   string `json:"class"`
		Address string `json:"address"`
		Age     string `json:"age"`
		gorm.Model
	}
)
