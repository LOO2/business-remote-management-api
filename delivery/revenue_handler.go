package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
	"github.com/gin-gonic/gin"
)

// RevenueHandler  represent the httphandler for revenue
type RevenueHandler struct {
	AUsecase models.RevenueUsecase
}

// NewRevenueleHandler will initialize the revenue/ resources endpoint
func NewRevenueHandler(c *Config) {
	//handler := &RevenueHandler{}

	groupRoute := c.R.Group("/api")
	{
		revenueRoute := groupRoute.Group("revenue")
		{
			revenueRoute.GET("/", ShowAllRevenues)
		}
	}

}

func ShowAllRevenues(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Revenue
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, p)
}

func ShowRevenue(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	db := database.GetDatabase()
	var p models.Revenue
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find revenue by ID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func CreateRevenue(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Revenue

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

	c.JSON(http.StatusOK, p)
}

func UpdateRevenue(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Revenue

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
			"error": "cannot save revenue: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func DeleteRevenue(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Revenue{}, newid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find revenue by id: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
