package controllers

import (
	"strconv"

	"github.com/LOO2/business-remote-management-api/database"
	"github.com/LOO2/business-remote-management-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ShowAllUsers(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.User
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.User
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	var encriptPassword string

	encriptPassword, error := HashPassword(p.Password)
	if error != nil {
		c.JSON(400, gin.H{
			"error": "cannot hash password: " + error.Error(),
		})
		return
	}
	p.Password = encriptPassword

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.User{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
