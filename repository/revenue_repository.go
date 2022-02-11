package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type Revenue struct {
	models.Revenue
}

type RevenueRepository interface {
	GetAll() (*Revenue, error)
	GetById() (*Revenue, error)
}

func GetAll() (*Revenue, error) {

	var p *Revenue

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