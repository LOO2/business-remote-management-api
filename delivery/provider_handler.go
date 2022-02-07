package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
	"github.com/gin-gonic/gin"
)

// represent the httphandler for revenue
type ProviderHandler struct {
	AUsecase models.ProviderUsecase
}

// will initialize the revenue/ resources endpoint
func NewProviderHandler(c *gin.Engine) {
	//handler := &RevenueHandler{}

	groupRoute := c.Group("/api")
	{
		revenueRoute := groupRoute.Group("revenue")
		{
			revenueRoute.GET("/", ShowAllProviders)
		}
	}

}

func ShowAllProviders(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Provider
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find provider by id: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func ShowProviders(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	db := database.GetDatabase()
	var p models.Provider
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find provider by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateProvider(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Provider

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
			"error": "cannot create provider: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func UpdateProvider(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Provider

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot update provider by ID: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteProvider(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Provider{}, newid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find provider by id: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
