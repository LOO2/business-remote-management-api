package repository

import (
	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

type User struct {
	models.User
}

type UserRepository interface {
	GetAllUsers() (*[]User, error)
	GetUserById() (*User, error)
	CreateUser(User) (*User, error)
	UpdateUser(User) (*User, error)
	DeleteUser(User) error
}

func GetAllUsers() (*[]User, error) {

	var p *[]User

	db := database.GetDatabase()

	err := db.Find(&p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetUserById(id int) (*User, error) {

	var p *User

	db := database.GetDatabase()

	err := db.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func CreateUser(new *User) error {

	db := database.GetDatabase()
	err := db.Create(&new).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(new *User) error {

	db := database.GetDatabase()
	err := db.Save(&new).Error
	if err != nil {
		return err
	}

	return nil

}

func DeleteUser(new User, id int) error {

	db := database.GetDatabase()
	err := db.Delete(&new, id).Error
	if err != nil {
		return err
	}

	return nil

}
