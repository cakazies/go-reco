package models

import "go-reco/app/configs"

type (
	StudentModels struct{}
)

func (*StudentModels) GetStudent() ([]Student, error) {
	result := make([]Student, 0)
	err := configs.GetDB.Model(&result).Scan(&result).Error
	return result, err
}

func (*StudentModels) Insert(data Student) error {
	err := configs.GetDB.Create(&data).Error
	return err
}
