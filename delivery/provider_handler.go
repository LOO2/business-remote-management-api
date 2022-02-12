package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/repository"
	"github.com/gin-gonic/gin"
)

func ShowAllProviders(c *gin.Context) {

	result, err := repository.GetAllProviders()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find provider " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ShowProvider(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	result, err := repository.GetByIdProvider(newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find provider by ID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateProvider(c *gin.Context) {

	var p *repository.Provider

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.CreateProvider(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create provider: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdateProvider(c *gin.Context) {

	var p *repository.Provider

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.UpdateProvider(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update provider: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func DeleteProvider(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.DeleteProvider(repository.Provider{}, newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find provider by id: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
