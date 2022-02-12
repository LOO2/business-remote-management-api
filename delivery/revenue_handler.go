package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/repository"
	"github.com/gin-gonic/gin"
)

// RevenueHandler  represent the httphandler for revenue
type RevenueHandler struct {
	//AUsecase models.RevenueUsecase
}

// NewRevenueleHandler will initialize the revenue/ resources endpoint
func NewRevenueHandler(c *gin.Engine) {
	//handler := &RevenueHandler{}

	//groupRoute := c.Group("/api")
	//{
	//	revenueRoute := groupRoute.Group("revenue")
	//	{
	//		revenueRoute.GET("/", GetAll)
	//	}
	//}

}

func ShowAllRevenues(c *gin.Context) {

	result, err := repository.GetAll()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find revenue " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
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

	result, err := repository.GetById(newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find revenue by ID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateRevenue(c *gin.Context) {

	var p *repository.Revenue

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	total := p.Debit + p.Delivery + p.Credit + p.Voucher
	p.Total = total

	err = repository.Create(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create revenue: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func UpdateRevenue(c *gin.Context) {

	var p *repository.Revenue

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.Update(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update revenue: " + err.Error(),
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

	err = repository.Delete(repository.Revenue{}, newid)
	//err = db.Delete(&models.Revenue{}, newid).Error

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find revenue by id: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
