package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type FunderRepository interface {
	FindFunder() ([]models.Funder, error)
	GetFunder(ID int) (models.Funder, error)
	CreateFunder(transaction models.Funder) (models.Funder, error)
}

func RepositoryFunder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFunder() ([]models.Funder, error) {
	var funders []models.Funder
	err := r.db.Preload("Donation.User").Preload("User").Find(&funders).Error

	return funders, err
}

func (r *repository) GetFunder(ID int) (models.Funder, error) {
	var funder models.Funder
	err := r.db.Preload("Donation").Preload("User").First(&funder).Error

	return funder, err
}

func (r *repository) CreateFunder(funder models.Funder) (models.Funder, error) {

	err := r.db.Preload("Donation").Preload("User").Create(&funder).Error

	return funder, err
}