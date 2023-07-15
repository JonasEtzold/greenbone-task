package repository

import (
	"fmt"

	"gorm.io/gorm"

	"greenbone-task/internal/persistence/db"
	models "greenbone-task/internal/persistence/models/computer"
)

type ComputerRepository struct {
	db *gorm.DB
}

var computerRepository *ComputerRepository

func GetComputer() *ComputerRepository {
	if computerRepository == nil {
		computerRepository = &ComputerRepository{
			db: db.Get(),
		}
	}

	return computerRepository
}

func (r *ComputerRepository) Get(name string) (*models.Computer, error) {
	var computer models.Computer
	r.db.First(&computer, fmt.Sprintf(`name='%s'`, name))

	if len(computer.Name) < 1 {
		return nil, fmt.Errorf("computer not found")
	}
	return &computer, nil
}

func (r *ComputerRepository) All() (*[]models.Computer, error) {
	var computers []models.Computer
	result := db.Get().Order("name asc").Find(&computers)
	return &computers, result.Error
}

func (r *ComputerRepository) Query(q *models.Computer) (*[]models.Computer, error) {
	var computers []models.Computer
	result := db.Get().Session(&gorm.Session{QueryFields: true}).Order("name asc").Find(&computers, q)
	return &computers, result.Error
}

func (r *ComputerRepository) Add(computer *models.Computer) error {
	return r.db.Create(&computer).Error
}

func (r *ComputerRepository) Update(computer *models.Computer) error {
	return db.Get().Save(&computer).Error
}

func (r *ComputerRepository) Delete(computer *models.Computer) error {
	return db.Get().Unscoped().Delete(&computer).Error
}
