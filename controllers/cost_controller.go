package controllers

import (
	"strconv"

	"github.com/LOO2/business-remote-management-api/database"
	"github.com/LOO2/business-remote-management-api/models"
	"github.com/gin-gonic/gin"
)

func ShowAllCost(c *gin.Context) {

	db := database.GetDatabase()
	var p []models.Cost
	err := db.Preload("CostCategory").Preload("CostProvider").Find(&p).Error
	//err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find cost: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowCost(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Cost
	err = db.Preload("CostCategory").Preload("CostProvider").First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateCost(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Cost

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
			"error": "cannot create revenue: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func UpdateCost(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Cost

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
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteCost(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Cost{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete revenue: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
