package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type CostCategory struct {
	models.CostCategory
}
type CostCategoryRepository interface {
	GetAllCostCategories() (*[]CostCategory, error)
	GetCostCategoryByID() (*CostCategory, error)
	CreateCostCategory(CostCategory) (*CostCategory, error)
	UpdateCostCategory(CostCategory) (*CostCategory, error)
	DeleteCostCategory(CostCategory) error
}

func GetAllCostCategories() (*[]CostCategory, error) {

	var p *[]CostCategory

	db := database.GetDatabase()

	err := db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetCostCategoryByID(id int) (*CostCategory, error) {

	var p *CostCategory

	db := database.GetDatabase()

	err := db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func CreateCostCategory(new *CostCategory) error {

	db := database.GetDatabase()
	err := db.Create(&new).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateCostCategory(new *CostCategory) error {

	db := database.GetDatabase()
	err := db.Save(&new).Error
	if err != nil {
		return err
	}

	return nil

}

func DeleteCostCategory(new CostCategory, id int) error {

	db := database.GetDatabase()
	err := db.Delete(&new, id).Error
	if err != nil {
		return err
	}

	return nil

}
