package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type Revenue struct {
	models.Revenue
}

type RevenueRepository interface {
	GetAllRevenues() ([]Revenue, error)
	GetRevenueById() (*Revenue, error)
	CreateRevenue(Revenue) (*Revenue, error)
	UpdateRevenue(Revenue) (*Revenue, error)
	DeleteRevenue(Revenue) error
}

func GetAllRevenues() ([]Revenue, error) {

	var p []Revenue

	db := database.GetDatabase()

	err := db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetRevenueById(id int) (*Revenue, error) {

	var p *Revenue

	db := database.GetDatabase()

	err := db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func CreateRevenue(new *Revenue) error {

	db := database.GetDatabase()
	err := db.Create(&new).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateRevenue(new *Revenue) error {

	db := database.GetDatabase()
	err := db.Save(&new).Error
	if err != nil {
		return err
	}

	return nil

}

func DeleteRevenue(new Revenue, id int) error {

	db := database.GetDatabase()
	err := db.Delete(&new, id).Error
	if err != nil {
		return err
	}

	return nil

}
