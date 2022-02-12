package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type Provider struct {
	models.Provider
}
type ProviderRepository interface {
	GetAllProviders() (*[]Provider, error)
	GetProviderById() (*Provider, error)
	CreateProvider(Provider) (*Provider, error)
	UpdateProvider(Provider) (*Provider, error)
	DeleteProvider(Provider) error
}

func GetAllProviders() (*[]Provider, error) {

	var p *[]Provider

	db := database.GetDatabase()

	err := db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetProviderById(id int) (*Provider, error) {

	var p *Provider

	db := database.GetDatabase()

	err := db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func CreateProvider(new *Provider) error {

	db := database.GetDatabase()
	err := db.Create(&new).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateProvider(new *Provider) error {

	db := database.GetDatabase()
	err := db.Save(&new).Error
	if err != nil {
		return err
	}

	return nil

}

func DeleteProvider(new Provider, id int) error {

	db := database.GetDatabase()
	err := db.Delete(&new, id).Error
	if err != nil {
		return err
	}

	return nil

}
