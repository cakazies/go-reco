package models

import "go-reco/app/configs"

type (
	StudentModels struct{}
)

func (*StudentModels) GetStudents() ([]Student, error) {
	result := make([]Student, 0)
	err := configs.GetDB.Model(&result).Scan(&result).Error
	return result, err
}

func (stud *Student) GetStudent() (*Student, error) {
	err := configs.GetDB.Model(&stud).Where("id = ?", stud.ID).Scan(&stud).Error
	return stud, err
}

func (stud *Student) Insert() error {
	err := configs.GetDB.Create(&stud).Error
	return err
}

// UpdateApprov function
func (stud *Student) UpdateStudent(id string) error {
	err := configs.GetDB.Model(&stud).
		Where("id = ?", id).
		Updates(&stud).Error
	return err
}
