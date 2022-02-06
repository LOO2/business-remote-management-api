package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// RevenueHandler  represent the httphandler for revenue
type RevenueHandler struct {
	AUsecase models.RevenueUsecase
}

type Config struct {
	R *gin.Engine
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
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(200, p)
}

func ShowRevenue(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	db := database.GetDatabase()
	var p models.Revenue
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func CreateRevenue(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Revenue

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func UpdateRevenue(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Revenue

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func DeleteRevenue(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Revenue{}, newid).Error

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
