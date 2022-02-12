package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type Cost struct {
	models.Cost
}
type CostRepository interface {
	GetAllCosts() (*[]Cost, error)
	GetCostById() (*Cost, error)
	CreateCost(Cost) (*Cost, error)
	UpdateCost(Cost) (*Cost, error)
	DeleteCost(Cost) error
}

func GetAllCosts() (*[]Cost, error) {

	var p *[]Cost

	db := database.GetDatabase()

	err := db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetCostByIdCost(id int) (*Cost, error) {

	var p *Cost

	db := database.GetDatabase()

	err := db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func CreateCost(new *Cost) error {

	db := database.GetDatabase()
	err := db.Create(&new).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateCost(new *Cost) error {

	db := database.GetDatabase()
	err := db.Save(&new).Error
	if err != nil {
		return err
	}

	return nil

}

func DeleteCost(new Cost, id int) error {

	db := database.GetDatabase()
	err := db.Delete(&new, id).Error
	if err != nil {
		return err
	}

	return nil

}
