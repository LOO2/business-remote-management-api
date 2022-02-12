package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/repository"
	"github.com/gin-gonic/gin"
)

func ShowAllCosts(c *gin.Context) {

	result, err := repository.GetAllCosts()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Cost " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ShowCost(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	result, err := repository.GetCostByIdCost(newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find Cost by ID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateCost(c *gin.Context) {

	var p *repository.Cost

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.CreateCost(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Cost: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdateCost(c *gin.Context) {

	var p *repository.Cost

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.UpdateCost(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update Cost: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func DeleteCost(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.DeleteCost(repository.Cost{}, newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find Cost by id: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
