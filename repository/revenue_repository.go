package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type Revenue struct {
	models.Revenue
}

type RevenueRepository interface {
	GetAll() (*[]Revenue, error)
	GetById() (*Revenue, error)
	Create(Revenue) (*Revenue, error)
	Update(Revenue) (*Revenue, error)
	Delete(Revenue) (*Revenue, error)
}

func GetAll() (*[]Revenue, error) {

	var p *[]Revenue

	db := database.GetDatabase()

	err := db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetById(id int) (*Revenue, error) {

	var p *Revenue

	db := database.GetDatabase()

	err := db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func Create(new *Revenue) error {

	db := database.GetDatabase()
	err := db.Create(&new).Error
	if err != nil {
		return err
	}

	return nil
}

func Update(new *Revenue) error {

	db := database.GetDatabase()
	err := db.Save(&new).Error
	if err != nil {
		return err
	}

	return nil

}

func Delete(new Revenue, id int) error {

	db := database.GetDatabase()
	err := db.Delete(&new, id).Error
	if err != nil {
		return err
	}

	return nil

}
