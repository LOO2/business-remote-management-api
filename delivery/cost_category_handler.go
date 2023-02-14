package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/repository"
	"github.com/gin-gonic/gin"
)

func ShowAllCostCategories(c *gin.Context) {

	result, err := repository.GetAllCostCategories()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find CostCategory " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ShowCostCategory(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	result, err := repository.GetCostCategoryByID(newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find CostCategory by ID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateCostCategory(c *gin.Context) {

	var p *repository.CostCategory

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.CreateCostCategory(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create CostCategory: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdateCostCategory(c *gin.Context) {

	var p *repository.CostCategory

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.UpdateCostCategory(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update CostCategory: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func DeleteCostCategory(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.DeleteCostCategory(repository.CostCategory{}, newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find CostCategory by id: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
